package server

import (
	config "github.com/lokesh-go/youtube-data-golang/src/config"
	dal "github.com/lokesh-go/youtube-data-golang/src/dal"
	httpServer "github.com/lokesh-go/youtube-data-golang/src/server/http"
)

type server struct {
	config      *config.Config
	dalServices dal.Methods
}

// Methods ...
type Methods interface {
	Start()
}

// New ...
func New(config *config.Config, dalServices dal.Methods) Methods {
	return &server{config, dalServices}
}

// Start ...
func (s *server) Start() {
	// Starts http server
	httpServer := httpServer.New(s.config, s.dalServices)
	httpServer.Start()
}
