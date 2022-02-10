package pagseguro

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/api-qop-v2/common"
	"github.com/api-qop-v2/ordereds"
	"github.com/api-qop-v2/partners"
	"github.com/api-qop-v2/payments/payment"
	"github.com/eucatur/go-toolbox/env"
	"github.com/eucatur/go-toolbox/handler"
	"github.com/eucatur/go-toolbox/jwt"
	"github.com/labstack/echo"
)

func PaymentPagSeguroHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func GenerateSessionPagSeguroHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*ordereds.Order)

	var pt partners.Partners
	pt.Parceiros_Id = p.Parceiros_Id
	partner, err := partners.SearchPartner(pt)
	if err != nil {
		return
	}

	idSession, err := GenerateSessionPagSeguro(p, partner[0])
	if err != nil {
		return c.JSON(401, err.Error())
	}

	return c.JSON(201, echo.Map{
		"id": idSession,
	})
}

func CheckoutPagSeguroHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*ordereds.Order)

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}

	pessoasId := fmt.Sprint(claims["PessoasId"])
	p.Pessoas_Id = pessoasId

	var pt partners.Partners
	pt.Parceiros_Id = p.Parceiros_Id
	partner, err := partners.SearchPartner(pt)
	if err != nil || len(partner) == 0 {
		return
	}

	idOrder, err := ordereds.CreateOrder(p, c)
	if err != nil {
		return
	}
	p.Pedidos_Id = idOrder

	code, err := CheckoutPagSeguro(p, partner[0])
	if err != nil {
		userID, _ := strconv.ParseInt(fmt.Sprint(claims["ClaimIDKey"]), 10, 64)
		p.Pessoas_Usuarios_Id = userID
		_ = ordereds.CancelOrder(p)
		return c.JSON(401, err.Error())
	}

	p.Code_Checkout = code
	_, err = payment.CreatePayment(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"id": idOrder})
}

func AlterStatusPagSeguroHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchPaymentPagSeguroHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}
