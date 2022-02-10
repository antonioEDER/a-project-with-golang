package partners

import (
	"fmt"
	"strings"

	"github.com/api-qop-v2/common"
	"github.com/eucatur/go-toolbox/env"
	"github.com/eucatur/go-toolbox/handler"
	"github.com/eucatur/go-toolbox/jwt"
	"github.com/labstack/echo"
)

func SearchPartnersFromProximityHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func ExistPartnerHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchPartnerHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchPartnerTypesActivesPublicHandler(c echo.Context) (err error) {

	ranges, err := SearchPartnerTypesActivesPublic()
	if err != nil {
		return
	}

	return c.JSON(200, echo.Map{
		"tipos": ranges,
	})
}

func SearchPartnerForUserHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Partners)

	listPartner, err := SearchPartnerForUser(p)
	if err != nil {
		return
	}

	return c.JSON(200, echo.Map{
		"parceiro": listPartner,
	})
}

func SearchPartnersAllForUserHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Partners)

	listPartners, err := SearchPartnersAllForUser(p)
	if err != nil {
		return
	}

	return c.JSON(200, echo.Map{
		"parceiros": listPartners,
	})
}

func SearchPartnersAllHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Partners)

	listPartners, err := SearchPartnersAll(p)
	if err != nil {
		return
	}

	return c.JSON(200, echo.Map{
		"parceiros": listPartners,
	})
}

func SearchPlansHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Partners)

	listProduct, listSales, err := SearchPlans(p)
	if err != nil {
		return
	}

	return c.JSON(200, echo.Map{
		"listPlansProducts": listProduct,
		"listplansSales":    listSales,
	})
}

func SearchPartnersForDescriptionHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchPartnerTypesAllHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func CreatePartnerHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func AlterHoursOfOperationHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Partners)

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}

	pessoasId := fmt.Sprint(claims["PessoasId"])

	err = AlterHoursOfOperation(pessoasId, p)
	if err != nil {
		return
	}

	return c.JSON(201, "sucesso")
}

func AlterPartnerHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func AlterPlanPartnerHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func CreatePlanPartnerHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchPlansPartnersHandler(c echo.Context) (err error) {
	f := *c.Get(handler.PARAMETERS).(*Filter)
	timeZone := strings.TrimSpace(c.Request().Header.Get("QOP-User-Time-Zone"))
	f.Time_Zone = timeZone

	planosValor, err := SearchPlansPartnerForValue(f)
	if err != nil {
		return
	}

	planosProduto, err := SearchPlansPartnerForProduct(f)
	if err != nil {
		return
	}

	return c.JSON(200, echo.Map{
		"planosValor":   planosValor,
		"planosProduto": planosProduto,
	})
}

func SearchFinancialSummaryHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchFinancialPartnerSummaryHandler(c echo.Context) (err error) {
	f := *c.Get(handler.PARAMETERS).(*Filter)
	timeZone := strings.TrimSpace(c.Request().Header.Get("QOP-User-Time-Zone"))
	f.Time_Zone = timeZone

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}

	personID := fmt.Sprint(claims["PessoasId"])
	if err != nil {
		return
	}

	partner, err := SearchPartnerFromPersonId(personID)
	f.Parceiros_Id = partner[0].Id

	resumo, grafico, err := SearchFinancialPartnerSummary(f)
	if err != nil {
		return
	}

	return c.JSON(200, echo.Map{
		"resumo":  resumo,
		"grafico": grafico,
	})
}

func SearchADMFinancialSummaryHandler(c echo.Context) (err error) {
	f := *c.Get(handler.PARAMETERS).(*Filter)
	timeZone := strings.TrimSpace(c.Request().Header.Get("QOP-User-Time-Zone"))
	f.Time_Zone = timeZone

	resumo, resumosParceiros, qtdPedidoPorStatus, qtdProdutosPorParceiros, grafico, graficosParceiros, err := SearchADMFinancialSummary(f)
	if err != nil {
		return
	}

	return c.JSON(200, echo.Map{
		"resumo":                  resumo,
		"resumosParceiros":        resumosParceiros,
		"qtdPedidoPorStatus":      qtdPedidoPorStatus,
		"qtdProdutosPorParceiros": qtdProdutosPorParceiros,
		"grafico":                 grafico,
		"graficosParceiros":       graficosParceiros,
	})
}

func CreateRangeActivityHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*RangeActivity)

	id, err := CreateRangeActivity(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"id": id})
}

func AlterRangeActivityHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*RangeActivity)

	err = AlterRangeActivity(p)
	if err != nil {
		return
	}

	return c.JSON(201, "sucesso")
}

func SearchInvoicePartnerADMHandler(c echo.Context) (err error) {
	f := *c.Get(handler.PARAMETERS).(*Filter)
	timeZone := strings.TrimSpace(c.Request().Header.Get("QOP-User-Time-Zone"))
	f.Time_Zone = timeZone

	fatura, err := SearchInvoicePartnerADM(f)
	if err != nil {
		return
	}

	return c.JSON(200, echo.Map{
		"fatura": fatura,
	})
}

func CreatePotentialPartnersHandler(c echo.Context) (err error) {
	f := *c.Get(handler.PARAMETERS).(*Filter)

	id, err := CreatePotentialPartners(f)
	if err != nil {
		return
	}

	return c.JSON(200, echo.Map{
		"id": id,
	})
}

func CreateEmailInvoiceHandler(c echo.Context) (err error) {
	err = CreateEmailInvoice(c)
	if err != nil {
		return
	}

	return c.JSON(201, "Sucesso")
}
