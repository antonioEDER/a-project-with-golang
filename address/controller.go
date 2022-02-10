package address

import (
	"github.com/api-qop-v2/config"
	"github.com/eucatur/go-toolbox/database"
	"github.com/jmoiron/sqlx"
)

func SearchCity(uf string, cidade string, tx *sqlx.Tx) (cities []City, err error) {
	return searchCityTx(uf, cidade, tx)
}

func CreateAddress(address PersonWeb) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	id, err = CreateAddressTx(address, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func AlterAddress(address PersonWeb) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	err = AlterAddressTx(address, tx)
	if err != nil {
		tx.Rollback() //echo.NewHTTPError(422, "Erro ao alterar cadastro")
		return
	}

	err = tx.Commit()

	return
}

func AlterAddressMain() {
	return
}

func DistanceBetweenPonters() {
	return
}

func CreateAddressWithUserAlreadyCreated() {
	return
}

func CreateUserAddressNotCreated() {
	return
}

func SearchAddress(personId string, tx *sqlx.Tx) (address []AddressWeb, err error) {
	return SearchAddresTx(personId, tx)

}
