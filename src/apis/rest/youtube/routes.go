package youtube

import (
	"github.com/labstack/echo/v4"

	config "github.com/lokesh-go/youtube-data-golang/src/config"
	dal "github.com/lokesh-go/youtube-data-golang/src/dal"
	middlewares "github.com/lokesh-go/youtube-data-golang/src/middlewares"
)

type dependencies struct {
	config      *config.Config
	dalServices dal.Methods
}

type Methods interface {
	RegisterRoutes(e *echo.Echo)
}

// New ...
func New(config *config.Config, dalServices dal.Methods) Methods {
	return &dependencies{config, dalServices}
}

// RegisterRoutes ...
func (d *dependencies) RegisterRoutes(e *echo.Echo) {
	// Define routes
	youtubeRoute := e.Group("/api/youtube")

	// Group V1 routes
	youtubeV1Routes := youtubeRoute.Group("/v1")
	{
		youtubeV1Routes.GET("/data", d.GetData, middlewares.Timeout(d.config.Timeout.GetDataHandler))
		youtubeV1Routes.POST("/data/:searchtext/search", d.Search, middlewares.Timeout(d.config.Timeout.SearchDataHandler))
	}
}
