package pix

import (
	gotoolboxtime "github.com/eucatur/go-toolbox/time"
)

type User struct {
	Id int64 `json:"id" db:"sale_id"`

	CreatedAt gotoolboxtime.Timestamp `json:"created_at" db:"sale_created_at"`
}

type Payment struct {
	Dados `json:"dados" db:"dados"`
}

type Dados struct {
	Chave_Pix   string  `json:"chave_pix" db:"chave_pix"`
	Valor_Total float64 `json:"valor_total" db:"valor_total"`
	Cidade      string  `json:"cidade" db:"cidade"`
	Pedidos_Id  int64   `json:"pedidos_id" db:"pedidos_id"`
	Fantasia    string  `json:"fantasia" db:"fantasia"`
}

type Response struct {
	Pix string `json:"pix"`
}
