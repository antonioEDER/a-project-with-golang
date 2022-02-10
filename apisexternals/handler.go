package apisexternals

import (
	"github.com/eucatur/go-toolbox/handler"
	"github.com/labstack/echo"
)

func DistanceBetweenPontersHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SenPushNotificationHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Push)
	err = SenPushNotification(p)
	if err != nil {
		return
	}

	return c.JSON(201, "Sucesso")
}

func SearcheAddressViaCepHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearcheAddressGoogleApisPostalCodeHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchServiceCorreiosForDescriptionHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearcheServiceCorreiosForServiceHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchServiceCorreiosAmountHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchServiceCorreiosHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}
