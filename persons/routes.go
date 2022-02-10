package persons

import (
	"github.com/api-qop-v2/address"
	"github.com/api-qop-v2/auth"
	"github.com/eucatur/go-toolbox/handler"

	"github.com/labstack/echo"
)

func AddRoutes(e *echo.Echo) {

	public := e.Group("v2/person")
	public.POST("/create", CreatePersonHandler, handler.MiddlewareBindAndValidate(&address.PersonWeb{}, "POST"))

	person := e.Group("v2/person", auth.Middleware)
	person.PUT("", AlterUserPersonHandler, handler.MiddlewareBindAndValidate(&address.PersonWeb{}))
	person.GET("", SearchPersonHandler, handler.MiddlewareBindAndValidate(&Filter{}))

	person.POST("/partners", CreatePersonPartnersHandler, handler.MiddlewareBindAndValidate(&address.PersonWeb{}, "POST"))
	person.POST("/partners/imagem", CreateImagemPersonPartnersHandler, handler.MiddlewareBindAndValidate(&address.PersonWeb{}, "POST"))
	person.PUT("/partners", AlterPersonPartnerHandler, handler.MiddlewareBindAndValidate(&address.PersonWeb{}, "POST"))
	person.PUT("/push-notification", AlterPersonToAddTokenPushHandler, handler.MiddlewareBindAndValidate(&address.PersonWeb{}))

}
