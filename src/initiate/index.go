package initiate

import (
	"log"

	config "github.com/lokesh-go/youtube-data-golang/src/config"
	dal "github.com/lokesh-go/youtube-data-golang/src/dal"
	jobModule "github.com/lokesh-go/youtube-data-golang/src/job"
	youtubePkg "github.com/lokesh-go/youtube-data-golang/src/pkg/youtube"
	server "github.com/lokesh-go/youtube-data-golang/src/server"
)

// Initializes the app
func Initialize() error {
	// Initialises the config
	config, err := config.Initialize()
	if err != nil {
		return err
	}
	log.Println("config initialised...")

	// Initialises youtube services
	ytServices, err := youtubePkg.NewServices(config)
	if err != nil {
		return err
	}
	log.Println("youtube services initialised...")

	// Intialises database services
	dalServices, err := dal.Initialize(config)
	if err != nil {
		return err
	}
	log.Println("database services initialised...")

	// Starts job to call the YouTube API continuously in background
	// And push data into database
	job := jobModule.New(config, ytServices, dalServices)
	job.Start()

	// Starts server
	server := server.New(config, dalServices)
	server.Start()

	// Returns
	return nil
}
