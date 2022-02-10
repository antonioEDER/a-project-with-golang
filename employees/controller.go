package employees

import (
	"encoding/hex"

	"github.com/api-qop-v2/address"
	"github.com/api-qop-v2/config"
	"github.com/api-qop-v2/persons"
	"github.com/api-qop-v2/tools"
	"github.com/eucatur/go-toolbox/database"
	"github.com/labstack/echo"
)

func CreateEmployees(person address.PersonWeb) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
	idPerson, err := persons.CreatePersonDBTx(person, 1, tx)
	if err != nil {
		err = echo.NewHTTPError(422, err)

		tx.Rollback()
		return
	}

	_, err = persons.CreateContactPersonTx(idPerson, person, 1, tx)
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

	idUser, err := persons.CreatePersonUserDBTx(idPerson, "", person, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	_, err = persons.CreatePersonMigrationDBTx(idUser, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = persons.CreatePersonStatusTx(idPerson, person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = persons.CreatePhysicalPersonTx(idPerson, person, 0, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	id, err = CreateEmployeesTx(idUser, person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func SearchEmployees() {
	return
}
func SearchEmployeesForDescription() {
	return
}

func AlterEmployees(person address.PersonWeb) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	person.He_Ativo = true
	_, err = persons.AlterPersonDBTx(person, tx)
	if err != nil {
		err = echo.NewHTTPError(422, err)

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

	err = persons.AlterPersonUserPassOrEmailDBTx(person.Pessoas_Id, person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = persons.AlterPersonStatusTx(person.Pessoas_Id, person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = persons.AlterPhysicalPersonTx(person.Pessoas_Id, person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = AlterEmployeesTx(person, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}
