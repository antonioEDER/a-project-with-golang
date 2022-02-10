package payment

import (
	gotoolboxtime "github.com/eucatur/go-toolbox/time"
)

type Order struct {
	ID                  int64                   `query:"id" json:"id"`
	Pedidos_Id          int64                   `query:"pedidos_id" json:"pedidos_id"`
	Pessoas_Usuarios_Id int64                   `query:"pessoas_usuarios_id" json:"pessoas_usuarios_id"`
	Code_Transacao      int64                   `query:"code_transacao" json:"code_transacao"`
	Status              string                  `query:"status" json:"status"`
	Code_Referencia     string                  `query:"code_referencia" json:"code_referencia"`
	Code_Checkout       string                  `query:"code_checkout" json:"code_checkout"`
	Servico_Pagamento   string                  `query:"servico_pagamento" json:"servico_pagamento"`
	He_Ativo            bool                    `query:"he_ativo" json:"he_ativo"`
	CREATED_AT          gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY          gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT           gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY           gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED             gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}
