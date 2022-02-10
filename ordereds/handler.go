package ordereds

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/api-qop-v2/apisexternals"
	"github.com/api-qop-v2/common"
	"github.com/api-qop-v2/users"
	"github.com/eucatur/go-toolbox/env"
	"github.com/eucatur/go-toolbox/handler"
	"github.com/eucatur/go-toolbox/jwt"
	"github.com/labstack/echo"
)

func SearchOrderToResendProductDigitalHandler(c echo.Context) (err error) {

	o := *c.Get(handler.PARAMETERS).(*Order)
	id, _ := strconv.ParseInt(fmt.Sprint(c.Param("id")), 10, 64)

	o.Pedidos_Id = id
	err = SearchOrderToResendProductDigital(o, c)
	if err != nil {
		return
	}

	order, err := SearchOrdersForId(o)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	u := users.User{}
	u.ID = order[0].Pessoas_Usuarios_Id
	us, err := users.SearchUserFromIdUser(u)
	if err != nil {
		return
	}
	push := apisexternals.Push{}
	push.To = us[0].Token_Push_App
	push.Notification.Body = "Foi enviado no e-mail o seu produto"
	push.Notification.Title = "qop - Enviado no e-mail"
	apisexternals.SenPushNotification(push)

	return c.JSON(201, id)
}

func SearchOrderSpecificHandler(c echo.Context) (err error) {

	timeZone := strings.TrimSpace(c.Request().Header.Get("QOP-User-Time-Zone"))
	o := *c.Get(handler.PARAMETERS).(*Order)

	idPerson := fmt.Sprint(c.Get("PessoasId"))
	if err != nil {
		return
	}

	o.Pessoas_Id = idPerson
	o.Time_Zone = timeZone

	orders, users, partners, err := SearchOrderSpecific(o)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{
		"pedido":   orders,
		"usuario":  users[0],
		"parceiro": partners[0],
	})
}

func SearchOrderHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchOrderForUserHandler(c echo.Context) (err error) {
	timeZone := strings.TrimSpace(c.Request().Header.Get("QOP-User-Time-Zone"))
	o := *c.Get(handler.PARAMETERS).(*Order)

	idPerson := fmt.Sprint(c.Get("PessoasId"))
	if err != nil {
		return
	}

	o.Pessoas_Id = idPerson
	o.Time_Zone = timeZone

	orders, err := SearchOrderForUser(o)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{
		"pedidos": orders,
	})
}

func SearchOrderDetailedUserHandler(c echo.Context) (err error) {
	timeZone := strings.TrimSpace(c.Request().Header.Get("QOP-User-Time-Zone"))
	o := *c.Get(handler.PARAMETERS).(*Order)

	idPerson := fmt.Sprint(c.Get("PessoasId"))
	if err != nil {
		return
	}

	o.Pessoas_Id = idPerson
	o.Time_Zone = timeZone
	orders, users, partners, err := SearchOrderSpecific(o)
	if err != nil {
		return
	}

	var p ParamsProductComposite
	p.Pedidos_Id = o.Pedidos_Id

	listComposite, err := SearchOrderDetailed(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{
		"pedido":           orders,
		"usuario":          users[0],
		"parceiro":         partners[0],
		"produto_composto": listComposite,
	})
}

func SearchOrderDetailedPartnerHandler(c echo.Context) (err error) {
	timeZone := strings.TrimSpace(c.Request().Header.Get("QOP-User-Time-Zone"))
	o := *c.Get(handler.PARAMETERS).(*Order)

	idPerson := fmt.Sprint(c.Get("PessoasId"))
	if err != nil {
		return
	}

	o.Pessoas_Id = idPerson
	o.Time_Zone = timeZone

	orders, users, partners, err := SearchOrderSpecificForParner(o)
	if err != nil {
		return
	}

	var p ParamsProductComposite
	p.Pedidos_Id = o.Pedidos_Id

	listComposite, err := SearchOrderDetailed(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{
		"pedido":           orders,
		"usuario":          users[0],
		"parceiro":         partners[0],
		"produto_composto": listComposite,
	})
}

func SearchOrderForPartnerHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func AlterOrdersAcceptedHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Order)
	timeZone := strings.TrimSpace(c.Request().Header.Get("QOP-User-Time-Zone"))
	p.Time_Zone = timeZone

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}

	userID := fmt.Sprint(claims["PessoasId"])
	p.Pessoas_Id = userID

	err = AlterOrdersAccepted(p)
	if err != nil {
		return
	}

	order, err := SearchOrdersForId(p)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	u := users.User{}
	u.ID = order[0].Pessoas_Usuarios_Id
	us, err := users.SearchUserFromIdUser(u)
	if err != nil {
		return
	}
	push := apisexternals.Push{}
	push.To = us[0].Token_Push_App
	push.Notification.Body = fmt.Sprintf(`Pedido - (%s)`, order[0].Pedidos_Status_Descricao)
	push.Notification.Title = fmt.Sprintf(`Status qop - (%s)`, order[0].Pedidos_Status_Descricao)
	apisexternals.SenPushNotification(push)

	return c.JSON(201, "Sucesso")
}
func SearchOrdersForFiltersHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Filter)
	timeZone := strings.TrimSpace(c.Request().Header.Get("QOP-User-Time-Zone"))
	p.Time_Zone = timeZone

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}

	userID, _ := strconv.ParseInt(fmt.Sprint(claims["PessoasId"]), 10, 64)

	p.Pessoas_Id = userID

	pedidos, err := SearchOrdersForFilters(p)

	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{
		"pedidos": pedidos,
	})
}

func SearchOrdersForCSVHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Filter)
	timeZone := strings.TrimSpace(c.Request().Header.Get("QOP-User-Time-Zone"))
	p.Time_Zone = timeZone

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}

	userID, _ := strconv.ParseInt(fmt.Sprint(claims["PessoasId"]), 10, 64)

	p.Pessoas_Id = userID

	csv, err := SearchOrdersForCSV(p)

	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{
		"csv": csv,
	})
}

func CreateOrderForPayCashHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func CreateOrderHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Order)

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}

	pessoasId := fmt.Sprint(claims["PessoasId"])
	p.Pessoas_Id = pessoasId
	id, err := CreateOrder(p, c)
	if err != nil {
		return
	}

	p.Pedidos_Id = id
	order, err := SearchOrdersForId(p)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	u := users.User{}
	u.ID = order[0].Pessoas_Usuarios_Id
	us, err := users.SearchUserFromIdUser(u)
	if err != nil {
		return
	}

	push := apisexternals.Push{}
	push.To = us[0].Token_Push_App
	push.Notification.Body = "Pedido realizado com sucesso!"
	push.Notification.Title = "qop - Pedido aberto"
	apisexternals.SenPushNotification(push)

	return c.JSON(201, echo.Map{"id": id})
}

func SearchCategoriesAndStatusOrderHandler(c echo.Context) (err error) {
	status, err := SearchCategoriesOrder()
	if err != nil {
		return
	}

	categorias, err := SearchStatusOrder()

	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{
		"status":     status,
		"categorias": categorias,
	})
}

func SearchOrderStatusHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchOrdersNewsHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func AlterOrderTrackingCodeHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func AlterOrderStatusHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Order)
	timeZone := strings.TrimSpace(c.Request().Header.Get("QOP-User-Time-Zone"))
	p.Time_Zone = timeZone

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}

	userID := fmt.Sprint(claims["PessoasId"])
	p.Pessoas_Id = userID

	err = AlterOrderStatus(p)

	order, err := SearchOrdersForId(p)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	u := users.User{}
	u.ID = order[0].Pessoas_Usuarios_Id
	us, err := users.SearchUserFromIdUser(u)
	if err != nil {
		return
	}
	push := apisexternals.Push{}
	push.To = us[0].Token_Push_App
	push.Notification.Body = fmt.Sprintf(`Pedido - (%s)`, order[0].Pedidos_Status_Descricao)
	push.Notification.Title = fmt.Sprintf(`Status qop - (%s)`, order[0].Pedidos_Status_Descricao)
	apisexternals.SenPushNotification(push)

	if err != nil {
		return
	}

	return c.JSON(201, "Sucesso")
}

func CancelOrderHandler(c echo.Context) (err error) {
	o := *c.Get(handler.PARAMETERS).(*Order)

	timeZone := strings.TrimSpace(c.Request().Header.Get("QOP-User-Time-Zone"))
	o.Time_Zone = timeZone

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}

	userID := fmt.Sprint(claims["PessoasId"])
	o.Pessoas_Id = userID

	o.Pedidos_Status_Id = 13
	err = AlterOrderStatus(o)
	if err != nil {
		return
	}

	err = CancelOrder(o)
	if err != nil {
		return
	}

	order, err := SearchOrdersForId(o)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	u := users.User{}
	u.ID = order[0].Pessoas_Usuarios_Id
	us, err := users.SearchUserFromIdUser(u)
	if err != nil {
		return
	}
	push := apisexternals.Push{}
	push.To = us[0].Token_Push_App
	push.Notification.Body = fmt.Sprintf(`Pedido - (%s)`, order[0].Pedidos_Status_Descricao)
	push.Notification.Title = fmt.Sprintf(`Status qop - (%s)`, order[0].Pedidos_Status_Descricao)
	apisexternals.SenPushNotification(push)

	return c.JSON(201, "Sucesso")
}

func GenerateCsvOrdersHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func CreateEmailBudgetHandler(c echo.Context) (err error) {

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}

	var o Order
	userID := fmt.Sprint(claims["PessoasId"])
	o.Pessoas_Id = userID

	err = CreateEmailBudget(o, c)
	if err != nil {
		return
	}

	return c.JSON(201, "Sucesso")
}
