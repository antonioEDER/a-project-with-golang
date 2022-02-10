package payment

import (
	"github.com/api-qop-v2/config"
	"github.com/api-qop-v2/ordereds"
	"github.com/api-qop-v2/users"
	"github.com/eucatur/go-toolbox/database"
	"github.com/jmoiron/sqlx"
)

// FindByFilterTx consulta por filtro usando uma transação do banco de dados
func FindByFilterTx(filter Filter, tx *sqlx.Tx) (user string, err error) {
	return findByFilterDBTx(filter, tx)
}

func CreatePayment(order ordereds.Order) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	var u users.User
	u.Pessoas_Id = order.Pessoas_Id
	us, err := users.SearchUserFromIdPersonTx(u, tx)
	if err != nil {
		return
	}

	order.Pessoas_Usuarios_Id = us[0].ID

	id, err = CreatePaymentTx(order, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}
