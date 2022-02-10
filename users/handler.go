package users

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/api-qop-v2/address"
	"github.com/api-qop-v2/partners"
	"github.com/api-qop-v2/products"

	"github.com/api-qop-v2/common"
	"github.com/eucatur/go-toolbox/env"
	"github.com/eucatur/go-toolbox/handler"
	"github.com/eucatur/go-toolbox/jwt"
	"github.com/labstack/echo"
)

func LoginHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*LoginParams)
	// Recurso para adaptar a nova API
	err = CheckPassWord(p)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	token, user, addressUser, err := Login(p, false)
	if err != nil {
		return
	}

	branch := []partners.RangeActivity{}
	if p.Tipo == "ADM" {
		branch, err = partners.SearchRangeActivity()
		if err != nil {
			return
		}
	}

	partner := []partners.Partners{}
	departments := []products.Department{}
	brands := []products.Brand{}
	categorys := []products.Category{}
	employees := []address.PersonWeb{}

	if p.Tipo == "PARCEIRO" || p.Tipo == "FUNCIONARIO" {

		if p.Tipo == "PARCEIRO" {
			partner, err = partners.SearchPartnerFromUserId(user[0].Id)
			if err != nil {
				return
			}
		}

		if p.Tipo == "FUNCIONARIO" {
			idUserFromPartner := user[0].Id
			idUserFromPartner, err = partners.SearchIdUserFromPartnerWithEmployeeData(idUserFromPartner)
			if err != nil {
				return
			}

			partner, err = partners.SearchPartnerFromUserId(idUserFromPartner)
			if err != nil {
				return
			}
		}

		departments, brands, categorys, err = products.SearchDepartmentsBrandsCategory(partner[0].Id)
		if err != nil {
			return
		}

		employees, err = partners.SearchEmployees(partner[0].Id)

		if err != nil {
			return
		}

	}

	return c.JSON(200, echo.Map{
		"token":           token,
		"usuario":         user,
		"enderecos":       addressUser,
		"ramo_atividades": branch,
		"parceiro":        partner,
		"departamentos":   departments,
		"marcas":          brands,
		"categorias":      categorys,
		"funcionarios":    employees,
	})
}

func UserHandler(c echo.Context) (err error) {

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}

	var p LoginParams
	p.Email = fmt.Sprint(claims["Email"])
	p.Tipo = fmt.Sprint(claims["Tipo"])
	p.Senha = ""

	token, user, addressUser, err := Login(p, false)
	if err != nil {
		return
	}

	branch, err := partners.SearchRangeActivity()
	if err != nil {
		return
	}

	partner := []partners.Partners{}
	departments := []products.Department{}
	brands := []products.Brand{}
	categorys := []products.Category{}
	employees := []address.PersonWeb{}

	if p.Tipo == "PARCEIRO" {
		partner, err = partners.SearchPartnerFromUserId(user[0].Id)
		if err != nil {
			return
		}

		departments, brands, categorys, err = products.SearchDepartmentsBrandsCategory(partner[0].Id)
		if err != nil {
			return
		}

		employees, err = partners.SearchEmployees(partner[0].Id)

		if err != nil {
			return
		}

	}

	return c.JSON(200, echo.Map{
		"token":           token,
		"usuario":         user,
		"enderecos":       addressUser,
		"ramo_atividades": branch,
		"parceiro":        partner,
		"departamentos":   departments,
		"marcas":          brands,
		"categorias":      categorys,
		"funcionarios":    employees,
	})
}

func RecoverPasswordHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*User)

	err = RecoverPassword(p)
	if err != nil {
		return
	}

	return c.JSON(200, 0)
}

func CreateNewPasswordHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*User)

	claims, err := jwt.VerifyTokenAndGetClaims(p.Token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}
	userID, err := strconv.ParseInt(fmt.Sprint(claims["ClaimIDKey"]), 10, 64)
	if err != nil {
		return c.JSON(401, echo.Map{"erro": err})
	}

	var u User
	u.Email = fmt.Sprint(claims["Email"])
	u.ID = userID
	u.Senha = p.Senha

	err = CreateNewPassword(u)
	if err != nil {
		return
	}

	return c.JSON(200, "Sucesso")
}

//
func ValidateAccountCreationByCodeHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*User)

	token, user, address, err := ValidateAccountCreationByCode(p)
	if err != nil {
		return err
	}

	return c.JSON(200, echo.Map{
		"token":     token,
		"usuario":   user,
		"enderecos": address,
	})
}

func ConfirmAccountHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*User)

	claims, err := jwt.VerifyTokenAndGetClaims(p.Token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}

	var u User
	u.Email = fmt.Sprint(claims["Email"])
	u.Tipo = "USUARIO"

	token, user, address, err := ConfirmAccount(u)
	if err != nil {
		return err
	}

	return c.JSON(200, echo.Map{
		"token":     token,
		"usuario":   user,
		"enderecos": address,
	})
}

func CreateLeadsForOffersHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*User)

	id, err := CreateLeadsForOffers(p)
	if err != nil {
		return
	}
	return c.JSON(201, echo.Map{"id": id})
}

func SendMessageByClientHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Contact)

	err = SendMessageByClient(p)
	if err != nil {
		return
	}
	return c.JSON(201, "Sucesso")
}
