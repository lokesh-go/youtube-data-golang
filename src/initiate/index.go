package initiate

import (
	config "github.com/lokesh-go/youtube-data-golang/src/config"
	youtubePkg "github.com/lokesh-go/youtube-data-golang/src/pkg/youtube"
)

// Initializes the app
func Initialize() error {
	// Initialises the config
	config, err := config.Initialize()
	if err != nil {
		return err
	}

	// Initialises youtube services
	youtubeServices, err := youtubePkg.NewServices(config)
	if err != nil {
		return err
	}

	// Returns
	return nil
}
