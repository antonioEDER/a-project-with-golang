package tools

import (
	"github.com/eucatur/go-toolbox/handler"

	"github.com/labstack/echo"
)

// AddRoutes adiciona as rotas
func AddRoutes(e *echo.Echo) {
	auth := e.Group("v2/encrypt")

	auth.POST("", FindByFilterHandler, handler.MiddlewareBindAndValidate(&Filter{}))
}
