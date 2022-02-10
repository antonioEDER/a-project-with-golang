package pagseguro

import (
	"github.com/api-qop-v2/auth"
	"github.com/api-qop-v2/ordereds"
	"github.com/eucatur/go-toolbox/handler"

	"github.com/labstack/echo"
)

// AddRoutes adiciona as rotas
func AddRoutes(e *echo.Echo) {

	pgSeguro := e.Group("v2/payment/pagseguro", auth.Middleware)
	pgSeguro.POST("/init-session", GenerateSessionPagSeguroHandler, handler.MiddlewareBindAndValidate(&ordereds.Order{}))
	pgSeguro.POST("/checkout", CheckoutPagSeguroHandler, handler.MiddlewareBindAndValidate(&ordereds.Order{}))

	pgSeguro.PUT("/alter-status", AlterStatusPagSeguroHandler, handler.MiddlewareBindAndValidate(&ordereds.Order{}))

}
