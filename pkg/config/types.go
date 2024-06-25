package config

type Config struct {
	HTTPServer *HTTPServer `yaml:"http"`
	Licence    *Licence    `yaml:"licence"`
	Storage    *Storage    `yaml:"storage"`

	App      *App      `yaml:"app"`
	Database *Database `yaml:"database"`
	Redis    *Redis    `yaml:"redis"`

	Services *Services `yaml:"services"`
}

type HTTPServer struct {
	Domains []string `yaml:"domains"`
	TLS     *TLS     `yaml:"tls"`
}

type TLS struct {
	Email string `yaml:"email"`
}

type Licence struct {
	ID  string `yaml:"id"`
	Key string `yaml:"key"`
}

type Storage struct {
	Data string `yaml:"data" default:"/var/lib/node-isp/"`
	Logs string `yaml:"logs" default:"/var/log/node-isp/"`
}

type App struct {
	Name string `yaml:"name"`
	Key  string `yaml:"key"`
}

type Database struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
}

type Redis struct {
	Password string `yaml:"password"`
}

type Services struct {
	GoogleMapsApiKey string `yaml:"google_maps_api_key"`
}
