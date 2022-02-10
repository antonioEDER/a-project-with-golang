package example

import (
	"github.com/api-qop-v2/config"
	"github.com/eucatur/go-toolbox/database"
	"github.com/jmoiron/sqlx"
)

// FindByFilter consulta por filtro criando uma transação
func FindByFilter(filter Filter) (user string, err error) {

	// Pattern to create a tx connection with the db
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
	defer tx.Rollback()

	user, err = FindByFilterTx(filter, tx)
	if err != nil {
		return
	}

	return
}

// FindByFilterTx consulta por filtro usando uma transação do banco de dados
func FindByFilterTx(filter Filter, tx *sqlx.Tx) (user string, err error) {
	return findByFilterDBTx(filter, tx)
}
