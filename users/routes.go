package users

import (
	"github.com/eucatur/go-toolbox/handler"
	"github.com/labstack/echo"
)

// AddRoutes adiciona as rotas
func AddRoutes(e *echo.Echo) {
	user := e.Group("v2/user")
	user.GET("", UserHandler)
	user.POST("/login", LoginHandler, handler.MiddlewareBindAndValidate(&LoginParams{}, "POST"))
	user.GET("/validate-code-email", ValidateAccountCreationByCodeHandler, handler.MiddlewareBindAndValidate(&User{}))
	user.POST("/confirm-registration", ConfirmAccountHandler, handler.MiddlewareBindAndValidate(&User{}))
	user.POST("/recover-password", RecoverPasswordHandler, handler.MiddlewareBindAndValidate(&User{}))
	user.POST("/new-password", CreateNewPasswordHandler, handler.MiddlewareBindAndValidate(&User{}))

	user.POST("/leads-for-offers", CreateLeadsForOffersHandler, handler.MiddlewareBindAndValidate(&User{}))
	user.POST("/contact-by-client", SendMessageByClientHandler, handler.MiddlewareBindAndValidate(&Contact{}))

}
