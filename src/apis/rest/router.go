package rest

import (
	"github.com/labstack/echo/v4"

	youtubeAPIs "github.com/lokesh-go/youtube-data-golang/src/apis/rest/youtube"
	config "github.com/lokesh-go/youtube-data-golang/src/config"
	dal "github.com/lokesh-go/youtube-data-golang/src/dal"
)

type router struct {
	config      *config.Config
	dalServices dal.Methods
}

type Methods interface {
	BuildRouter(e *echo.Echo)
}

// New ...
func New(config *config.Config, dalServices dal.Methods) Methods {
	return &router{config, dalServices}
}

// BuildRouter ...
func (r *router) BuildRouter(e *echo.Echo) {
	// Registers routes for youtube module
	youtubeModule := youtubeAPIs.New(r.config, r.dalServices)
	youtubeModule.RegisterRoutes(e)
}
