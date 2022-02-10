package persons

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/api-qop-v2/address"
	"github.com/api-qop-v2/common"
	"github.com/api-qop-v2/config"
	"github.com/api-qop-v2/images"
	"github.com/api-qop-v2/send_emails"
	"github.com/api-qop-v2/tools"
	"github.com/api-qop-v2/users"

	"github.com/eucatur/go-toolbox/database"
	"github.com/eucatur/go-toolbox/env"
	"github.com/labstack/echo"
)

func CreatePerson(person address.PersonWeb) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	id, err = CreatePersonDBTx(person, 0, tx)
	if err != nil {
		err = echo.NewHTTPError(422, err)
		tx.Rollback()
		return
	}

	_, err = CreateContactPersonTx(id, person, 0, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	person.Pessoas_Id = id

	_, err = address.CreateAddressTx(person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	codValidate, _ := rand.Prime(rand.Reader, 11)

	plaintext := []byte(person.Senha)
	passEncryp, err := tools.Encrypt(plaintext)
	if err != nil {
		tx.Rollback()
		return
	}
	person.Senha = hex.EncodeToString(passEncryp)

	idUser, err := CreatePersonUserDBTx(id, fmt.Sprint(codValidate), person, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	_, err = CreatePersonMigrationDBTx(idUser, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	types, err := searchTypesPersonTx(person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	person.Tipos_Pessoas_Id = types[0].Id

	_, err = CreatePhysicalPersonTx(id, person, 0, tx)
	if err != nil {
		if strings.Contains(err.Error(), "pessoas_fisicas") {
			err = echo.NewHTTPError(422, MessageErrorCNPJExists)
		} else {
			fmt.Println(err.Error())
		}
		tx.Rollback()
		return
	}

	err = tx.Commit()

	var lg users.LoginParams
	lg.Email = person.Email
	lg.Tipo = person.Tipo

	token, _, _, err := users.Login(lg, false)

	var params send_emails.EmailConfirmAccount
	params.Email = person.Email
	params.Name = person.Nome
	params.Cod_Confirmacao = fmt.Sprint(codValidate)
	params.Token = token
	params.Url = env.MustString(common.UrlFront)

	go send_emails.SendAccountConfirmation(params)

	return
}

func SearchPersonExists(person address.PersonWeb) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	p, err := searchPersonExistsTx(person, tx)
	if len(p) > 0 {
		err = echo.NewHTTPError(422, "Usuário já tem cadastro.")
		if !p[0].He_Ativo {
			err = echo.NewHTTPError(422, "Usuário tem cadastro mas está desativado.")
		}
		return
	}

	if err != nil {
		err = echo.NewHTTPError(422, MessageErrorEmailWebExists)
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func CreatTokenpushNotificationForWeb() {
	return
}

func CreatTokenpushNotificationForApp() {
	return
}

func SearchPerson() {
	return
}

func CreateLeadsForOffers() {
	return
}

func RecoverPassword() {
	return
}

func CreateNewPassword() {
	return
}

func ValidateAccountCreationByCode() {
	return
}

func AlterUserPerson(person address.PersonWeb) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	err = AlterUserPersonTx(person, tx)
	if err != nil {
		tx.Rollback() // echo.NewHTTPError(422, "Erro ao alterar cadastro")
		return
	}

	err = tx.Commit()

	return
}

func CreatePersonPartners(person address.PersonWeb) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
	person.He_Ativo = true
	id, err = CreatePersonDBTx(person, 1, tx)
	if err != nil {
		err = echo.NewHTTPError(422, err)

		tx.Rollback()
		return
	}

	_, err = CreateContactPersonTx(id, person, 1, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	person.Pessoas_Id = id
	idPartner, err := CreatePersonPartnersTx(person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	person.Parceiros_Id = idPartner
	idTransport, err := CreateTransportTx(person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	person.Transportadoras_Id = idTransport

	_, err = CreatePartnerTransportTx(person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	person.Pessoas_Id = id
	_, err = address.CreateAddressTx(person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	person.Uid = "0"

	plaintext := []byte(person.Senha)
	passEncryp, err := tools.Encrypt(plaintext)
	if err != nil {
		tx.Rollback()
		return
	}
	person.Senha = hex.EncodeToString(passEncryp)

	idUser, err := CreatePersonUserDBTx(id, "", person, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	_, err = CreatePersonMigrationDBTx(idUser, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = CreatePersonStatusTx(id, person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	types, err := searchTypesPersonTx(person, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	person.Tipos_Pessoas_Id = types[0].Id

	if person.Cpf != "" {
		_, err = CreatePhysicalPersonTx(id, person, 0, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	} else {
		_, err = CreateLegalPersonTx(id, person, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}

	activePlanProduct := true
	activePlanSales := false

	if person.He_Plano_Venda {
		activePlanSales = true
		activePlanProduct = false
	}

	_, err = CreatePlanForProductTx(activePlanProduct, person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = CreatePlanForSaleTx(activePlanSales, person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func AlterPersonPartner(person address.PersonWeb) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	person.He_Ativo = true
	id, err = AlterPersonDBTx(person, tx)
	if err != nil {
		err = echo.NewHTTPError(422, err)

		tx.Rollback()
		return
	}

	id = person.Pessoas_Id
	err = AlterContactPersonTx(id, person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = AlterPersonPartnersTx(person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = address.AlterAddressTx(person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	person.Uid = "0"

	if person.Senha != "" {
		plaintext := []byte(person.Senha)
		passEncryp, _ := tools.Encrypt(plaintext)
		person.Senha = hex.EncodeToString(passEncryp)
	}

	err = AlterPersonUserDBTx(id, person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	var u users.User
	u.Pessoas_Id = fmt.Sprint(person.Pessoas_Id)
	user, _ := users.SearchUserFromIdPersonTx(u, tx)
	_, err = CreatePersonMigrationDBTx(user[0].ID, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = AlterPersonStatusTx(id, person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	if person.Cpf != "" {
		err = AlterPhysicalPersonTx(id, person, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	} else {
		err = AlterLegalPersonTx(id, person, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}

	idPartner := person.Parceiros_Id
	err = AlterLinkPartnerTransportTx(idPartner, person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	activePlanProduct := true
	activePlanSales := false

	if person.He_Plano_Venda {
		activePlanSales = true
		activePlanProduct = false
	}
	err = AlterPlanForProductTx(activePlanProduct, person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = AlterPlanForSaleTx(activePlanSales, person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func ConfirmAccount() {
	return
}

func AlterPersonToAddTokenPush(person address.PersonWeb) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	err = AlterPersonToAddTokenPushTx(person, tx)
	if err != nil {
		tx.Rollback() // echo.NewHTTPError(422, "Erro ao alterar cadastro")
		return
	}

	err = tx.Commit()

	return
}

func CreateImageToPerson(i images.ImageGeneric, c echo.Context) (diretorio string, err error) {
	file, err := c.FormFile("file")
	if err != nil {
		return
	}

	i.Filename = c.FormValue("filename")
	fileExtension := path.Ext(file.Filename)
	fileName := fmt.Sprintf(`%s%s`, i.Filename, fileExtension)
	dirBucket := "/imagens_parceiros/"

	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()

	// Destination
	PathFileTemp := env.MustString(common.PathFileTemp)
	dst, err := os.Create(PathFileTemp + dirBucket + fileName)
	if err != nil {
		return
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return
	}

	diretorio = dirBucket + fileName

	return
}
