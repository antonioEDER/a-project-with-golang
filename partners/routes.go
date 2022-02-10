package partners

import (
	"github.com/api-qop-v2/auth"
	"github.com/eucatur/go-toolbox/handler"

	"github.com/labstack/echo"
)

// AddRoutes adiciona as rotas
func AddRoutes(e *echo.Echo) {

	partner := e.Group("v2/partner", auth.Middleware)

	partner.POST("/range", CreateRangeActivityHandler, handler.MiddlewareBindAndValidate(&RangeActivity{}))
	partner.PUT("/range", AlterRangeActivityHandler, handler.MiddlewareBindAndValidate(&RangeActivity{}))
	partner.GET("/list/private", SearchPartnersAllHandler, handler.MiddlewareBindAndValidate(&Partners{}))
	partner.GET("/plans", SearchPlansHandler, handler.MiddlewareBindAndValidate(&Partners{}))
	partner.PUT("/working-hours", AlterHoursOfOperationHandler, handler.MiddlewareBindAndValidate(&Partners{}))

	partner.GET("/financial-summary-to-adm", SearchADMFinancialSummaryHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	partner.GET("/financial-summary-to-partner", SearchFinancialPartnerSummaryHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	partner.GET("/commercial-invoice/private", SearchInvoicePartnerADMHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	partner.POST("/commercial-invoice/private", CreateEmailInvoiceHandler, handler.MiddlewareBindAndValidate(&Filter{}))

	partner.GET("/plans/private", SearchPlansPartnersHandler, handler.MiddlewareBindAndValidate(&Filter{}))

	partner.GET("/proximity", SearchPartnersFromProximityHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	partner.GET("/range/private", SearchPartnerTypesAllHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	partner.GET("/by-name-cnpj/private", SearchPartnersForDescriptionHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	partner.GET("/exist", ExistPartnerHandler, handler.MiddlewareBindAndValidate(&Filter{}))

	public := e.Group("v2/partner/public")
	public.GET("/specific", SearchPartnerForUserHandler, handler.MiddlewareBindAndValidate(&Partners{}))
	public.GET("/range", SearchPartnerTypesActivesPublicHandler, handler.MiddlewareBindAndValidate(&RangeActivity{}))
	public.GET("/list", SearchPartnersAllForUserHandler, handler.MiddlewareBindAndValidate(&Partners{}))
	public.POST("/leads-partner", CreatePotentialPartnersHandler, handler.MiddlewareBindAndValidate(&Filter{}))
}
