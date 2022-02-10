package transports

import (
	"github.com/api-qop-v2/auth"
	"github.com/eucatur/go-toolbox/handler"

	"github.com/labstack/echo"
)

// AddRoutes adiciona as rotas
func AddRoutes(e *echo.Echo) {
	transport := e.Group("v2/transport", auth.Middleware)
	transport.POST("", CreateTransportsHandler, handler.MiddlewareBindAndValidate(&Transport{}))
	transport.GET("", SearchTransportsActiveHandler, handler.MiddlewareBindAndValidate(&Transport{}))
	transport.PUT("", AlterTransportHandler, handler.MiddlewareBindAndValidate(&Transport{}))
	transport.GET("/list-for-partner", SearchTransportsByActiveForPartnerHandler, handler.MiddlewareBindAndValidate(&Transport{}))

	transport.GET("/list", SearchTransportsActiveHandler, handler.MiddlewareBindAndValidate(&Transport{}))
	transport.POST("/service", CreateServiceHandler, handler.MiddlewareBindAndValidate(&TransportService{}))
	transport.PUT("/service", AlterServiceHandler, handler.MiddlewareBindAndValidate(&TransportService{}))

	transport.GET("/services", SearchServicesHandler, handler.MiddlewareBindAndValidate(&TransportService{}))
	transport.GET("/services-for-partner", SearchServicesByIdPartnerHandler, handler.MiddlewareBindAndValidate(&TransportService{}))

	transport.GET("/service-km-by-id", SearchServiceToKmByIdHandler, handler.MiddlewareBindAndValidate(&TransportServiceKM{}))
	transport.GET("/service-amount-by-id", SearchServiceToMoneyByIdHandler, handler.MiddlewareBindAndValidate(&TransportServiceAmount{}))

	transport.POST("/service-km", CreateServiceToKMHandler, handler.MiddlewareBindAndValidate(&TransportServiceKM{}))
	transport.POST("/service-amount", CreateServiceToMoneyHandler, handler.MiddlewareBindAndValidate(&TransportServiceAmount{}))

	transport.PUT("/service-km", AlterServiceToKMHandler, handler.MiddlewareBindAndValidate(&TransportServiceKM{}))
	transport.PUT("/service-amount", AlterServiceToMoneyHandler, handler.MiddlewareBindAndValidate(&TransportServiceAmount{}))

	transport.POST("/create-for-partner", CreateTransportsHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	transport.GET("/list", SearchTransportsActiveHandler, handler.MiddlewareBindAndValidate(&Filter{}))

	transport.GET("/list-for-partner-from-user", SearchTransportsByActiveForUserHandler, handler.MiddlewareBindAndValidate(&Transport{}))

}
