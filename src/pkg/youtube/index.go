package youtube

import (
	"context"
	"net/http"

	option "google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"

	config "github.com/lokesh-go/youtube-data-golang/src/config"
	googlePkg "github.com/lokesh-go/youtube-data-golang/src/pkg/google"
)

type service struct {
	youtube *youtube.Service
	config  *config.Config
}

// NewServices ...
func NewServices(config *config.Config) (Methods, error) {
	// Gets google client
	client, err := getGoogleClient(config)
	if err != nil {
		return nil, err
	}

	// Gets youtube services
	youtubeServices, err := youtube.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	// Returns
	return &service{youtubeServices, config}, nil
}

func getGoogleClient(config *config.Config) (client *http.Client, err error) {
	// youtube scope
	youtubeScope := youtube.YoutubeReadonlyScope

	// Gets google client with youtube scope
	googleMethods := googlePkg.New(config.Youtube.Auth.ClientSecretFilePath, config.Youtube.Auth.TokenFilePath, youtubeScope)
	client, err = googleMethods.GetClient()
	if err != nil {
		return nil, err
	}

	// Returns
	return client, nil
}
