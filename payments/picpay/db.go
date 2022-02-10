package picpay

import (
	"errors"

	"github.com/api-qop-v2/ordereds"
	"github.com/jmoiron/sqlx"
)

func findByFilterDBTx(filter Filter, tx *sqlx.Tx) (user string, err error) {

	if filter.ID > 0 && tx != nil {
		user = "Seens we have a user :)"
	} else {
		err = errors.New("There's no user")
	}

	return
}

func CancelPaymentTx(order ordereds.Order, tx *sqlx.Tx) (err error) {
	query := `
		UPDATE pagamentos 
			SET status = $1, updated_at = NOW(), he_ativo = 0
        WHERE pedidos_id = $2 AND pessoas_usuarios_id = $3
		`

	args := []interface{}{
		"Cancelado",
		order.Pedidos_Id,
		order.Pessoas_Usuarios_Id,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}
