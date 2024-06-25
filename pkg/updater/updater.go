package updater

import (
	"context"
	"errors"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/apex/log"
	"github.com/containers/image/v5/docker"
	"github.com/containers/image/v5/docker/reference"
	"github.com/containers/image/v5/types"
)

var appRepo = "ghcr.io/node-isp/node-isp"

// CurrentAppVersion is set via the server bootstrap process, from either the state or the baked in version
var CurrentAppVersion string

type Updater struct{}

type Update struct {
	Component  string
	Repository string
	Version    *semver.Version
}

// Start runs a check for updates every hour, and send them down a channel to the server for processing
func (u *Updater) Start() (<-chan Update, error) {

	l := log.WithField("component", "updater")

	updates := make(chan Update)

	go func() {
		for {
			l.Info("Checking for updates...")
			latestVersion, err := u.LatestAppVersion()

			if err != nil {
				l.WithError(err).Error("Failed to get latest app version")
				time.Sleep(1 * time.Hour)
				continue
			}

			currentVersion, err := semver.NewVersion(CurrentAppVersion)
			if err != nil {
				l.WithError(err).Error("Failed to parse current app version")
				time.Sleep(1 * time.Hour)
				continue
			}

			if currentVersion.LessThan(latestVersion) {
				l.
					WithField("current", currentVersion.String()).
					WithField("latest", latestVersion.String()).
					Info("New version available")

				updates <- Update{
					Component:  "app",
					Repository: appRepo,
					Version:    latestVersion,
				}

			} else {
				l.WithField("service", "app").WithField("current", currentVersion.String()).WithField("latest", latestVersion.String()).Info("No updates available")
			}

			// Sleep for 12 hours
			time.Sleep(1 * time.Hour)
		}
	}()

	return updates, nil
}

func (u *Updater) LatestAppVersion() (*semver.Version, error) {
	return u.latestVersion(appRepo)
}

func (u *Updater) latestVersion(repo string) (*semver.Version, error) {
	_, tagList, err := u.getRepoTags(repo)

	if err != nil {
		return nil, err
	}

	var vs []*semver.Version
	for _, tag := range tagList {
		v, err := semver.NewVersion(tag)
		if err != nil {
			continue
		}
		vs = append(vs, v)
	}

	return vs[len(vs)-1], nil
}

func (u *Updater) getRepoTags(containerImage string) (repositoryName string, tagList []string, err error) {
	// Get the latest tag from the registry
	imgRef, err := u.parseDockerRepositoryReference(containerImage)

	if err != nil {
		return
	}

	repositoryName = imgRef.DockerReference().Name()
	tagList, err = docker.GetRepositoryTags(context.TODO(), &types.SystemContext{}, imgRef)
	return
}

func (u *Updater) parseDockerRepositoryReference(refString string) (types.ImageReference, error) {
	ref, err := reference.ParseNormalizedNamed(refString)
	if err != nil {
		return nil, err
	}

	if !reference.IsNameOnly(ref) {
		return nil, errors.New(`no tag or digest allowed in reference`)
	}

	// Checks ok, now return a reference. This is a hack because the tag listing code expects a full image reference even though the tag is ignored
	return docker.NewReference(reference.TagNameOnly(ref))
}
