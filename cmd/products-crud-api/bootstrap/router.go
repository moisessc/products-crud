package bootstrap

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	// logsFormat the constant with the logger format
	logsFormat = "${time_custom} ip=${remote_ip} method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n"
	// timeFormat the constant with the time format for the logger
	timeFormat = "2006/01/02 15:04:05"
)

// newEchoRouter builds an instance of the echo router
func newEchoRouter() http.Handler {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           logsFormat,
		CustomTimeFormat: timeFormat,
	}))

	products := e.Group("/api/v1/products")
	products.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello world")
	})

	return e
}
