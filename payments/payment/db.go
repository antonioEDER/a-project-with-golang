package payment

import (
	"errors"
	"fmt"

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

func CreatePaymentTx(order ordereds.Order, tx *sqlx.Tx) (id int64, err error) {
	query := `
	INSERT INTO public.pagamentos(
		pessoas_usuarios_id,
		pedidos_id,
		status,
		code_referencia,
		servico_pagamento,
		he_ativo,
		code_checkout)
	VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
		`

	codeCheckout := ""
	if order.Code_Checkout != "" {
		codeCheckout = order.Code_Checkout
	}

	args := []interface{}{
		order.Pessoas_Usuarios_Id,
		order.Pedidos_Id,
		"pendente",
		"PD-" + fmt.Sprint(order.Pedidos_Id),
		order.Servico_Pagamento,
		1,
		codeCheckout,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}
