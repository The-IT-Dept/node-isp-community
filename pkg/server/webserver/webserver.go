package webserver

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"path/filepath"
	"sync"
	"time"

	"github.com/apex/log"
	"github.com/caddyserver/certmagic"
)

const (
	// HTTPChallengePort is the officially-designated port for
	// the HTTP challenge according to the ACME spec.
	HTTPChallengePort = 80

	// TLSALPNChallengePort is the officially-designated port for
	// the TLS-ALPN challenge according to the ACME spec.
	TLSALPNChallengePort = 443
)

// Port variables must remain their defaults unless you
// forward packets from the defaults to whatever these
// are set to; otherwise ACME challenges will fail.
var (
	// HTTPPort is the port on which to serve HTTP
	// and, as such, the HTTP challenge (unless
	// Default.AltHTTPPort is set).
	HTTPPort = 80

	// HTTPSPort is the port on which to serve HTTPS
	// and, as such, the TLS-ALPN challenge
	// (unless Default.AltTLSALPNPort is set).
	HTTPSPort = 443
)

// Variables for conveniently serving HTTPS.
var (
	lnMu   sync.Mutex
	httpWg sync.WaitGroup
)

func New(
	mux *http.ServeMux,
	dataDir string,
	domains []string,
	email string,
	log *log.Entry,
) *WebServer {
	return &WebServer{
		mux:     mux,
		dataDir: dataDir,
		domains: domains,
		email:   email,
		log:     log,
	}
}

type WebServer struct {
	dataDir string
	domains []string
	email   string

	mux *http.ServeMux

	log *log.Entry

	cache *certmagic.Cache
	magic *certmagic.Config
}

func (w *WebServer) Run(ctx context.Context) error {
	w.cache = certmagic.NewCache(certmagic.CacheOptions{
		GetConfigForCert: func(cert certmagic.Certificate) (*certmagic.Config, error) {
			return certmagic.New(w.cache, certmagic.Config{}), nil
		},
	})

	w.magic = certmagic.New(w.cache, certmagic.Config{
		Storage: &certmagic.FileStorage{Path: filepath.Join(w.dataDir, "/certs")},
	})

	myACME := certmagic.NewACMEIssuer(w.magic, certmagic.ACMEIssuer{
		CA:     certmagic.LetsEncryptProductionCA,
		Email:  w.email,
		Agreed: true,
	})

	w.magic.Issuers = []certmagic.Issuer{
		myACME,
	}

	if err := w.magic.ManageSync(context.TODO(), w.domains); err != nil {
		w.log.WithError(err).Fatal("Failed to manage certificates")
	}

	httpWg.Add(1)
	defer httpWg.Done()

	lnMu.Lock()

	httpLn, err := net.Listen("tcp", fmt.Sprintf(":%d", HTTPPort))
	if err != nil {
		log.WithError(err).Fatal("Failed to listen on HTTP port")
		lnMu.Unlock()
		return err
	}

	tlsConfig := w.magic.TLSConfig()
	tlsConfig.NextProtos = append([]string{"h2", "http/1.1"}, tlsConfig.NextProtos...)

	httpsLn, err := tls.Listen("tcp", fmt.Sprintf(":%d", HTTPSPort), tlsConfig)
	if err != nil {
		log.WithError(err).Fatal("Failed to listen on HTTPS port")
		lnMu.Unlock()
		return err
	}

	go func() {
		httpWg.Wait()
		lnMu.Lock()
		lnMu.Unlock()
	}()

	lnMu.Unlock()

	httpServer := &http.Server{
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
		BaseContext:       func(listener net.Listener) context.Context { return ctx },
	}

	httpServer.Handler = myACME.HTTPChallengeHandler(http.HandlerFunc(httpRedirectHandler))

	httpsServer := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      2 * time.Minute,
		IdleTimeout:       5 * time.Minute,
		Handler:           w.mux,
		BaseContext:       func(listener net.Listener) context.Context { return ctx },
	}

	// w.log.Infof("Serving HTTP->HTTPS on %s and %s", hln.Addr(), hsln.Addr())

	go httpServer.Serve(httpLn)
	return httpsServer.Serve(httpsLn)
}

func httpRedirectHandler(w http.ResponseWriter, r *http.Request) {
	toURL := "https://"

	// since we redirect to the standard HTTPS port, we
	// do not need to include it in the redirect URL
	requestHost := hostOnly(r.Host)

	toURL += requestHost
	toURL += r.URL.RequestURI()

	// get rid of this disgusting unencrypted HTTP connection 🤢
	w.Header().Set("Connection", "close")

	http.Redirect(w, r, toURL, http.StatusMovedPermanently)
}

func hostOnly(hostport string) string {
	host, _, err := net.SplitHostPort(hostport)
	if err != nil {
		return hostport // not a host:port
	}
	return host
}
