package persons

import (
	"fmt"

	"github.com/api-qop-v2/address"
	"github.com/api-qop-v2/images"
	"github.com/eucatur/go-toolbox/handler"
	"github.com/labstack/echo"
)

func CreateImagemPersonPartnersHandler(c echo.Context) (err error) {

	var i images.ImageGeneric
	dir, err := CreateImageToPerson(i, c)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"diretorio": dir})

}

func CreatePersonPartnersHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*address.PersonWeb)
	err = SearchPersonExists(p)
	if err != nil {
		return
	}

	id, err := CreatePersonPartners(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"id": id})

}

func AlterPersonPartnerHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*address.PersonWeb)
	id, err := AlterPersonPartner(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"id": id})

}

func AlterPersonToAddTokenPushHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*address.PersonWeb)

	idPerson := fmt.Sprint(c.Get("PessoasId"))
	if err != nil {
		return
	}
	p.Id = idPerson
	err = AlterPersonToAddTokenPush(p)
	if err != nil {
		return
	}

	return c.JSON(201, "Sucesso")

}

func CreatePersonHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*address.PersonWeb)

	err = SearchPersonExists(p)
	if err != nil {
		return
	}

	id, err := CreatePerson(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"id": id})

}

func CreatTokenpushNotificationForWebHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func CreatTokenpushNotificationForAppHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchPersonHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func CreateLeadsForOffersHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func RecoverPasswordHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func CreateNewPasswordHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func ValidateAccountCreationByCodeHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchPersonExistsHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func AlterUserPersonHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*address.PersonWeb)
	idPerson := fmt.Sprint(c.Get("PessoasId"))
	if err != nil {
		return
	}

	p.Id = idPerson
	err = AlterUserPerson(p)
	if err != nil {
		return
	}
	return c.JSON(200, "Sucesso")
}

func ConfirmAccountHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}
