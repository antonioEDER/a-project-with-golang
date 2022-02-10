package employees

import (
	"github.com/api-qop-v2/address"
	"github.com/api-qop-v2/persons"
	"github.com/eucatur/go-toolbox/handler"
	"github.com/labstack/echo"
)

func CreateEmployeesHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*address.PersonWeb)

	err = persons.SearchPersonExists(p)
	if err != nil {
		return
	}

	id, err := CreateEmployees(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"id": id})
}

func SearchEmployeesHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchEmployeesForDescriptionHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func AlterEmployeesHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*address.PersonWeb)

	err = AlterEmployees(p)
	if err != nil {
		return
	}

	return c.JSON(201, "sucesso")
}
