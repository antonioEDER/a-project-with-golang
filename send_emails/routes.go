package send_emails

import (
	"github.com/api-qop-v2/auth"
	"github.com/eucatur/go-toolbox/handler"

	"github.com/labstack/echo"
)

// AddRoutes adiciona as rotas
func AddRoutes(e *echo.Echo) {
	emails := e.Group("v2/send-emails", auth.Middleware)

	emails.POST("/send-email-proposal", SendEmailWithProposalHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	emails.POST("/commercial-invoice", SendEmailWithInvoiceHandler, handler.MiddlewareBindAndValidate(&Filter{}))

}
