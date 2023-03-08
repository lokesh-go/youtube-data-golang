package youtube

import (
	"context"

	option "google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"

	config "github.com/lokesh-go/youtube-data-golang/src/config"
	utils "github.com/lokesh-go/youtube-data-golang/src/utils"
)

type service struct {
	youtube *youtube.Service
	config  *config.Config
}

// NewServices ...
func NewServices(config *config.Config) (Methods, error) {
	// Gets youtube services
	youtubeServices, err := youtube.NewService(context.Background(), option.WithAPIKey(utils.GetEnv(config.Youtube.APIKey)))
	if err != nil {
		return nil, err
	}

	// Returns
	return &service{youtubeServices, config}, nil
}
