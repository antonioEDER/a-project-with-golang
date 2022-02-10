package apisexternals

import (
	"github.com/eucatur/go-toolbox/handler"

	"github.com/labstack/echo"
)

// AddRoutes adiciona as rotas
func AddRoutes(e *echo.Echo) {
	auth := e.Group("v2/service-external")

	auth.GET("/viacep", SearcheAddressViaCepHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	auth.GET("/send-push", SenPushNotificationHandler, handler.MiddlewareBindAndValidate(&Push{}))
	auth.GET("/google-postal-code", SearcheAddressGoogleApisPostalCodeHandler, handler.MiddlewareBindAndValidate(&Filter{}))

	auth.GET("/service-correios", SearchServiceCorreiosForDescriptionHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	auth.GET("/service-amount-correios", SearchServiceCorreiosHandler, handler.MiddlewareBindAndValidate(&Filter{}))

}
