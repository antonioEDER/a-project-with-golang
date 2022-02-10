package ordereds

import (
	"github.com/api-qop-v2/auth"
	"github.com/eucatur/go-toolbox/handler"

	"github.com/labstack/echo"
)

// AddRoutes adiciona as rotas
func AddRoutes(e *echo.Echo) {

	auth := e.Group("v2/order", auth.Middleware)
	auth.POST("", CreateOrderHandler, handler.MiddlewareBindAndValidate(&Order{}))
	auth.GET("/specific", SearchOrderSpecificHandler, handler.MiddlewareBindAndValidate(&Order{}))
	auth.GET("/details-for-user", SearchOrderDetailedUserHandler, handler.MiddlewareBindAndValidate(&Order{}))
	auth.GET("/details-for-partner", SearchOrderDetailedPartnerHandler, handler.MiddlewareBindAndValidate(&Order{}))
	auth.GET("/list", SearchOrderForUserHandler, handler.MiddlewareBindAndValidate(&Order{}))

	auth.GET("/categories-status", SearchCategoriesAndStatusOrderHandler)
	auth.GET("/for-filter", SearchOrdersForFiltersHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	auth.PUT("/accepted", AlterOrdersAcceptedHandler, handler.MiddlewareBindAndValidate(&Order{}))
	auth.PUT("/status", AlterOrderStatusHandler, handler.MiddlewareBindAndValidate(&Order{}))

	auth.POST("/email/budget", CreateEmailBudgetHandler, handler.MiddlewareBindAndValidate(&SendBudget{}))

	auth.DELETE("", CancelOrderHandler, handler.MiddlewareBindAndValidate(&Order{}))

	auth.GET("/generate-csv", SearchOrdersForCSVHandler, handler.MiddlewareBindAndValidate(&Filter{}))

	public := e.Group("v2/order/public")
	public.GET("/resend-product-digital/:id", SearchOrderToResendProductDigitalHandler, handler.MiddlewareBindAndValidate(&Order{}))

}
