package address

import (
	"fmt"
	"strconv"

	"github.com/eucatur/go-toolbox/handler"
	"github.com/labstack/echo"
)

// func FindByFilterHandler(c echo.Context) (err error) {
// 	p := *c.Get(handler.PARAMETERS).(*Filter)

// 	sectionals, err := FindByFilter(p)
// 	if err != nil {
// 		return
// 	}

// 	return c.JSON(200, sectionals)
// }

func DistanceBetweenPontersHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func CreateAddressHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func AlterAddressHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*PersonWeb)
	idPerson := fmt.Sprint(c.Get("PessoasId"))
	if err != nil {
		return
	}

	p.Pessoas_Id, err = strconv.ParseInt(idPerson, 10, 64)
	if err != nil {
		return
	}

	err = AlterAddress(p)
	if err != nil {
		return
	}
	return c.JSON(201, "Sucesso")
}

func AlterAddressMainHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func CreateAddressWithUserAlreadyCreatedHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*PersonWeb)
	idPerson := fmt.Sprint(c.Get("PessoasId"))
	if err != nil {
		return
	}

	p.Pessoas_Id, err = strconv.ParseInt(idPerson, 10, 64)
	if err != nil {
		return
	}

	id, err := CreateAddress(p)
	if err != nil {
		return
	}
	return c.JSON(201, echo.Map{"id": id})
}

func CreateUserAddressNotCreatedHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchAddressHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchAddressForUserHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}
