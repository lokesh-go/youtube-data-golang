package initiate

import (
	config "github.com/lokesh-go/youtube-data-golang/src/config"
	dal "github.com/lokesh-go/youtube-data-golang/src/dal"
	jobModule "github.com/lokesh-go/youtube-data-golang/src/job"
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
	ytServices, err := youtubePkg.NewServices(config)
	if err != nil {
		return err
	}

	// Intialises database services
	dalServices, err := dal.Initialize(config)
	if err != nil {
		return err
	}

	// Starts job to call the YouTube API continuously in background
	// And push data into database
	job := jobModule.New(config, ytServices, dalServices)
	job.Start()

	// Returns
	return nil
}
