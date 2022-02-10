package picpay

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

func CheckoutPicPayHandler(c echo.Context) (err error) {
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

	var pt partners.Partners
	pt.Parceiros_Id = p.Parceiros_Id
	partner, err := partners.SearchPartner(pt)
	if err != nil {
		return
	}

	p.Pedidos_Id = idOrder

	responsePicPay, err := CheckoutPicPay(p, partner[0])
	if err != nil {
		return c.JSON(401, err.Error())
	}

	_, err = payment.CreatePayment(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{
		"id":     idOrder,
		"picpay": responsePicPay,
	})
}

func CancelpaymentPicPayHandler(c echo.Context) (err error) {
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

	orders, userPerson, _, err := ordereds.SearchOrderSpecific(p)

	if len(orders) == 0 || len(userPerson) == 0 {
		err = fmt.Errorf("Pedido não encontrado")
		return c.JSON(401, err.Error())
	}

	var pt partners.Partners
	pt.Parceiros_Id = orders[0].Parceiros_Id
	partner, err := partners.SearchPartner(pt)
	if err != nil {
		return c.JSON(401, err.Error())
	}

	p.Pessoas_Usuarios_Id, err = strconv.ParseInt(fmt.Sprint(userPerson[0].Id), 10, 64)
	if err != nil {
		return c.JSON(401, err.Error())
	}

	err = CancelPaymentPicPay(p, partner[0])
	if err != nil {
		return c.JSON(401, err.Error())
	}

	return c.JSON(201, "Sucesso")
}

func SearchPaymentPicPayHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}
