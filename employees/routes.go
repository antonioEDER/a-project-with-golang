package employees

import (
	"github.com/api-qop-v2/address"
	"github.com/api-qop-v2/auth"
	"github.com/eucatur/go-toolbox/handler"

	"github.com/labstack/echo"
)

// AddRoutes adiciona as rotas
func AddRoutes(e *echo.Echo) {

	employee := e.Group("v2/employees", auth.Middleware)
	employee.POST("", CreateEmployeesHandler, handler.MiddlewareBindAndValidate(&address.PersonWeb{}))
	employee.PUT("", AlterEmployeesHandler, handler.MiddlewareBindAndValidate(&address.PersonWeb{}))

	employee.GET("", SearchEmployeesHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	employee.GET("/description", SearchEmployeesForDescriptionHandler, handler.MiddlewareBindAndValidate(&Filter{}))

}
