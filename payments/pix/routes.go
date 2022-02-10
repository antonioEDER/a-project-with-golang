package pix

import (
	"github.com/api-qop-v2/auth"
	"github.com/api-qop-v2/ordereds"
	"github.com/eucatur/go-toolbox/handler"

	"github.com/labstack/echo"
)

// AddRoutes adiciona as rotas
func AddRoutes(e *echo.Echo) {
	pgPix := e.Group("v2/payment/pix", auth.Middleware)

	pgPix.POST("/checkout", CheckoutPixHandler, handler.MiddlewareBindAndValidate(&ordereds.Order{}))
	pgPix.DELETE("/cancel", CancelPaymentPixHandler, handler.MiddlewareBindAndValidate(&ordereds.Order{}))

	pgPix.POST("/alter-status", CheckoutPixHandler, handler.MiddlewareBindAndValidate(&ordereds.Order{}))

}
