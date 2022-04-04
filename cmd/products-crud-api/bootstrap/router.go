package bootstrap

import (
	"net/http"

	pv "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"products-crud/internal/controller"
	"products-crud/pkg/validator"
)

const (
	// logsFormat the constant with the logger format
	logsFormat = "${time_custom} ip=${remote_ip} method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n"
	// timeFormat the constant with the time format for the logger
	timeFormat = "2006/01/02 15:04:05"
)

// newEchoRouter builds an instance of the echo router
func newEchoRouter(ph *controller.ProductsHandler) http.Handler {
	e := echo.New()
	e.Validator = validator.New(pv.New())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           logsFormat,
		CustomTimeFormat: timeFormat,
	}))

	products := e.Group("/api/v1/products")
	products.POST("", ph.Create)

	return e
}
