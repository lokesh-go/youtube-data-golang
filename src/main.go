package main

import (
	"log"

	"github.com/lokesh-go/youtube-data-golang/src/initiate"
)

func main() {
	// Initializes the app
	err := initiate.Initialize()
	if err != nil {
		log.Fatal("failed to initialise app: ", err.Error())
	}
}
