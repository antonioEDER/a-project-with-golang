package users

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/api-qop-v2/address"
	"github.com/api-qop-v2/common"
	"github.com/api-qop-v2/config"
	"github.com/api-qop-v2/partners"
	"github.com/api-qop-v2/send_emails"
	"github.com/api-qop-v2/tools"
	"github.com/eucatur/go-toolbox/database"
	"github.com/eucatur/go-toolbox/env"
	"github.com/eucatur/go-toolbox/jwt"
	"github.com/labstack/echo"
)

func Login(param LoginParams, validateByPassword bool) (token string, users []PersonWeb, addressed []address.AddressWeb, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
	defer tx.Rollback()

	users, err = LoginDBTx(param, validateByPassword, tx)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	if len(users) <= 0 {
		err = echo.NewHTTPError(401, "Usuário e/ou senha inválido(s)")
		return
	}

	if param.Tipo == "ADM" && !users[0].He_Adm {
		err = echo.NewHTTPError(401, "Usuário sem permissão")
		return
	}

	userID := users[0].Id
	personID := users[0].Pessoas_Id
	email := users[0].Email
	tipo := ""

	switch users[0].Tipos_Pessoas_Id {
	case "1":
		tipo = "USUARIO"
	case "2":
		tipo = "PARCEIRO"
	case "3":
		tipo = "ADM"
	case "4":
		tipo = "FUNCIONARIO"
	case "5":
		tipo = "TRANSPORTADORA"
	case "6":
		tipo = "ENTREGADORES"
	}

	users[0].Tipo = tipo

	claims := map[string]interface{}{
		"ClaimIDKey": userID,
		"PessoasId":  personID,
		"Email":      email,
		"Tipo":       tipo,
	}

	token, err = jwt.CreateTokenWithClaims(claims, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return
	}

	if users[0].Tipo != "FUNCIONARIO" {
		addressed, err = address.SearchAddress(personID, tx)
		if err != nil {
			return
		}
	}
	return
}

func CreatTokenpushNotificationForWeb() {
	return
}

func CreatTokenpushNotificationForApp() {
	return
}

func CreateUser() {
	return
}

func CheckPassWord(user LoginParams) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	r, err := CheckPassWordDBTx(user, tx)
	if len(r) == 0 {
		tx.Rollback()
		return
	}

	ciphertext, _ := hex.DecodeString(r[0].Senha)
	result, err := tools.Decrypt(ciphertext)

	if err != nil {
		err = echo.NewHTTPError(422, "Erro ao Decrypt")
		return
	}

	passDecrypt := fmt.Sprintf("%s", result)

	if passDecrypt != user.Senha {
		err = echo.NewHTTPError(422, "Senha incorreta")
		return
	}

	err = tx.Commit()

	return
}

func RecoverPassword(user User) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	r, err := searchUserTx(user, tx)
	if len(r) == 0 {
		// err = echo.NewHTTPError(422, "Não encontramos uma conta com esse e-mail.")
		return
	}

	var lg LoginParams
	lg.Email = user.Email

	token, _, _, err := Login(lg, false)

	var params send_emails.EmailConfirmAccount
	params.Email = user.Email
	params.Name = r[0].Nome
	params.Token = token
	params.Url = env.MustString(common.UrlFront)

	go send_emails.SendRecoverPassword(params)

	return
}

func CreateNewPassword(user User) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	r, err := searchUserTx(user, tx)
	if err != nil {
		err = echo.NewHTTPError(422, "Erro validar email")
		tx.Rollback()
		return
	}

	if len(r) == 0 {
		tx.Rollback()
		err = echo.NewHTTPError(422, "Não encontramos uma conta com esse email.")
		return
	}

	plaintext := []byte(user.Senha)
	passEncryp, err := tools.Encrypt(plaintext)
	if err != nil {
		tx.Rollback()
		return
	}
	user.Senha = hex.EncodeToString(passEncryp)

	err = NewPasswordTx(user, tx)
	if err != nil {
		err = echo.NewHTTPError(422, "Erro criar nova senha.")
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func ValidateAccountCreationByCode(user User) (token string, users []PersonWeb, addressed []address.AddressWeb, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	r, err := searchCodeValidateTx(user, tx)
	if len(r) == 0 {
		err = echo.NewHTTPError(422, MessageErrorValidateAccountByCode)
		return
	}

	if err != nil {
		err = echo.NewHTTPError(422, MessageErrorValidateAccountByCode)
		tx.Rollback()
		return
	}

	if r[0].He_Ativo {
		err = echo.NewHTTPError(422, "A conta já está ativa. Tente recuperar a senha.")
		tx.Rollback()
		return
	}

	err = activeUserTx(r[0].Pessoas_Id, tx)
	if err != nil {
		err = echo.NewHTTPError(422, MessageErrorValidateAccountByCode)
		tx.Rollback()
		return
	}

	err = tx.Commit()

	var lg LoginParams
	lg.Email = r[0].Email
	lg.Senha = r[0].Senha
	lg.Tipo = r[0].Tipo

	token, users, addressed, err = Login(lg, false)

	return
}

func SearchUserExists() {
	return
}

func AlterDataUser() {
	return
}

func SearchUserAllDataFromIdPerson(user User) (users []PersonWeb, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	users, err = SearchUserAllDataFromIdPersonTx(user, tx)
	if len(users) == 0 {
		err = echo.NewHTTPError(422, MessageErrorValidateAccountByCode)
		return
	}

	err = tx.Commit()
	return
}

func ConfirmAccount(user User) (token string, users []PersonWeb, addressed []address.AddressWeb, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	r, err := searchUserTx(user, tx)
	if len(r) == 0 {
		err = echo.NewHTTPError(422, MessageErrorValidateAccountByCode)
		return
	}

	if err != nil {
		err = echo.NewHTTPError(422, MessageErrorValidateAccountByCode)
		tx.Rollback()
		return
	}

	if r[0].He_Ativo {
		err = echo.NewHTTPError(422, "A conta já está ativa. Tente recuperar a senha.")
		tx.Rollback()
		return
	}

	err = activeUserTx(r[0].Pessoas_Id, tx)
	if err != nil {
		err = echo.NewHTTPError(422, MessageErrorValidateAccountByCode)
		tx.Rollback()
		return
	}

	err = tx.Commit()

	var lg LoginParams
	lg.Email = r[0].Email
	lg.Senha = r[0].Senha
	lg.Tipo = "USUARIO"

	token, users, addressed, err = Login(lg, false)

	return
}

func CreateLeadsForOffers(user User) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	id, err = CreateLeadsForOffersTx(user, tx)
	if err != nil {
		tx.Rollback() //echo.NewHTTPError(422, "Erro ao alterar cadastro")
		return
	}

	err = tx.Commit()

	return
}

func SendMessageByClient(c Contact) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	partnerID, err := strconv.ParseInt(c.Partner_ID, 10, 64)
	if err != nil {
		return
	}

	var p partners.Partners
	p.Parceiros_Id = partnerID

	partner, err := partners.SearchPartnerForUserTx(p, tx)
	if err != nil {
		tx.Rollback() //echo.NewHTTPError(422, "Erro ao alterar cadastro")
		return
	}

	var params send_emails.Contact
	params.Nome = c.Name
	params.Telefone = c.Phone
	params.Email = c.Email
	params.Assunto = c.Subject
	params.Texto = c.Text
	params.Title = "Você tem um nova mensagem"

	if len(partner) > 0 {
		params.NomeParceiro = partner[0].Fantasia
		params.EmailParceiro = partner[0].Email
		go send_emails.SendContactClient(params)
	}

	params.NomeParceiro = c.Name
	params.EmailParceiro = c.Email
	params.Title = "Você enviou uma mensagem no qop"

	go send_emails.SendContactClient(params)

	err = tx.Commit()

	return
}

func SearchUserFromIdUser(u User) (user []User, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	user, err = SearchUserFromIdUserTx(u, tx)
	if err != nil {
		tx.Rollback() //echo.NewHTTPError(422, "Erro ao alterar cadastro")
		return
	}

	err = tx.Commit()

	return
}
