package address

import (
	"github.com/api-qop-v2/auth"
	"github.com/eucatur/go-toolbox/handler"

	"github.com/labstack/echo"
)

func AddRoutes(e *echo.Echo) {
	address := e.Group("v2/address", auth.Middleware)

	address.POST("", CreateAddressWithUserAlreadyCreatedHandler, handler.MiddlewareBindAndValidate(&PersonWeb{}))
	address.PUT("", AlterAddressHandler, handler.MiddlewareBindAndValidate(&PersonWeb{}))

	address.GET("", SearchAddressForUserHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	address.PUT("/activate-main", AlterAddressMainHandler, handler.MiddlewareBindAndValidate(&Filter{}))

	address.GET("/distance-between-ponters", DistanceBetweenPontersHandler, handler.MiddlewareBindAndValidate(&Filter{}))

}
