package pix

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/api-qop-v2/common"
	"github.com/api-qop-v2/ordereds"
	"github.com/api-qop-v2/partners"
	"github.com/api-qop-v2/payments/payment"
	"github.com/api-qop-v2/users"
	"github.com/eucatur/go-toolbox/env"
	"github.com/eucatur/go-toolbox/handler"
	"github.com/eucatur/go-toolbox/jwt"
	"github.com/labstack/echo"
)

func PaymentPagSeguroHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func GenerateSessionPagSeguroHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func CheckoutPixHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*ordereds.Order)

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}

	pessoasId := fmt.Sprint(claims["PessoasId"])
	p.Pessoas_Id = pessoasId

	idOrder, err := ordereds.CreateOrder(p, c)
	if err != nil {
		return
	}

	p.Pedidos_Id = idOrder

	_, err = payment.CreatePayment(p)
	if err != nil {
		return
	}

	var pt partners.Partners
	pt.Parceiros_Id = p.Parceiros_Id
	partner, err := partners.SearchPartner(pt)
	if err != nil {
		return
	}

	var u users.User
	u.Pessoas_Id = pessoasId
	user, err := users.SearchUserAllDataFromIdPerson(u)

	result, err := GenerateQrCode(p, partner[0], user[0])
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{
		"id":                idOrder,
		"pix":               result.Pix,
		"chave_publica_pix": partner[0].Chave_Pix,
	})
}

func CancelPaymentPixHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*ordereds.Order)
	timeZone := strings.TrimSpace(c.Request().Header.Get("QOP-User-Time-Zone"))
	p.Time_Zone = timeZone

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}

	pessoasId := fmt.Sprint(claims["PessoasId"])
	p.Pessoas_Id = pessoasId

	orders, userPerson, partner, err := ordereds.SearchOrderSpecific(p)
	if len(orders) == 0 || len(userPerson) == 0 || len(partner) == 0 {
		err = fmt.Errorf("Pedido não encontrado")
		return c.JSON(401, err.Error())
	}

	var pt partners.Partners
	pt.Parceiros_Id = orders[0].Parceiros_Id

	p.Pessoas_Usuarios_Id, err = strconv.ParseInt(fmt.Sprint(userPerson[0].Id), 10, 64)
	if err != nil {
		return c.JSON(401, err.Error())
	}

	err = CancelPaymentPix(p)
	if err != nil {
		return c.JSON(401, err.Error())
	}

	return c.JSON(201, "Sucesso")
}

func AlterStatusPagSeguroHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchPaymentPagSeguroHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func CheckoutPicPayHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchPaymentPicPayHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}
