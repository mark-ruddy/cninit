package cninit

import (
	"fmt"
	"github.com/mittwald/go-helm-client"
)

// CreateHelmClient will return a helmclient.Client which is tied to a namespace
func CreateHelmClient(namespace string) (helmclient.Client, error) {
	var helmClient helmclient.Client
	config, err := getClientRestConfig()
	if err != nil {
		return helmClient, err
	}
	opt := &helmclient.RestConfClientOptions{
		Options: &helmclient.Options{
			RepositoryCache:  "/tmp/.helmcache",
			RepositoryConfig: "/tmp/.helmrepo",
			Debug:            true,
			Linting:          true,
			Namespace:        namespace,
		},
		RestConfig: config,
	}
	helmClient, err = helmclient.NewClientFromRestConf(opt)
	if err != nil {
		return helmClient, fmt.Errorf("while creating helmClient: %w", err)
	}
	return helmClient, nil
}
