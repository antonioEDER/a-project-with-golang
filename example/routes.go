package example

import (
	"github.com/api-qop-v2/auth"
	"github.com/eucatur/go-toolbox/handler"

	"github.com/labstack/echo"
)

// AddRoutes adiciona as rotas
func AddRoutes(e *echo.Echo) {
	auth := e.Group("v2/example", auth.Middleware)

	auth.GET("", FindByFilterHandler, handler.MiddlewareBindAndValidate(&Filter{}))
}
