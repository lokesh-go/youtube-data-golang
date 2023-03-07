package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	utils "github.com/lokesh-go/youtube-data-golang/src/utils"
)

// Timeout ...
func Timeout(t int64) echo.MiddlewareFunc {
	// Forms error msg
	errMsg := map[string]interface{}{"code": "01m01t01", "message": "handler timeout"}
	errBytes, _ := utils.JSONMarshal(errMsg)

	// Returns
	return middleware.TimeoutWithConfig(
		middleware.TimeoutConfig{
			Timeout:      time.Duration(t) * time.Millisecond,
			ErrorMessage: string(errBytes),
		})
}
