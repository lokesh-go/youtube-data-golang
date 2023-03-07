package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	config "github.com/lokesh-go/youtube-data-golang/src/config"
	dal "github.com/lokesh-go/youtube-data-golang/src/dal"
)

type httpConfig struct {
	config      *config.Config
	dalServices dal.Methods
}

// Methods ...
type Methods interface {
	Start()
}

// New ...
func New(config *config.Config, dalServices dal.Methods) Methods {
	return &httpConfig{config, dalServices}
}

// Start ...
func (h *httpConfig) Start() {
	// Initialises server
	echoServer := echo.New()

	// Health check
	echoServer.GET("/ping", h.checkPing)

	// Starts servers
	echoServer.Start(h.config.Server.HTTP.Port)
}

func (h *httpConfig) checkPing(c echo.Context) (err error) {
	// Ping response
	status := map[string]interface{}{
		"app": h.config.App,
	}

	// Returns
	return c.JSON(http.StatusOK, status)
}
