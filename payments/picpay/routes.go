package picpay

import (
	"github.com/api-qop-v2/auth"
	"github.com/api-qop-v2/ordereds"
	"github.com/eucatur/go-toolbox/handler"

	"github.com/labstack/echo"
)

// AddRoutes adiciona as rotas
func AddRoutes(e *echo.Echo) {

	pgPay := e.Group("v2/payment/picpay", auth.Middleware)
	pgPay.POST("/checkout", CheckoutPicPayHandler, handler.MiddlewareBindAndValidate(&ordereds.Order{}))

	pgPay.DELETE("/cancel", CancelpaymentPicPayHandler, handler.MiddlewareBindAndValidate(&ordereds.Order{}))

	pgPay.POST("/alter-status", SearchPaymentPicPayHandler, handler.MiddlewareBindAndValidate(&ordereds.Order{}))

}
