package send_emails

import (
	gotoolboxtime "github.com/eucatur/go-toolbox/time"
)

type User struct {
	Id int64 `json:"id" db:"sale_id"`

	CreatedAt gotoolboxtime.Timestamp `json:"created_at" db:"sale_created_at"`
}

//Request struct
type RequestEmail struct {
	From        string
	To          []string
	Subject     string
	Body        string
	Attachments map[string][]byte
	CC          []string
	BCC         []string
	Title       string
}

type EmailConfirmAccount struct {
	Name            string
	Email           string
	Cod_Confirmacao string
	Token           string `json:"token" db:"token" query:"token"`
	Url             string
}

type OrderItens struct {
	Itens []OrderCreate `json:"itens" db:"itens" query:"itens"`
}
type OrderCreate struct {
	Status_Id         int64  `json:"status_id" db:"Status_Id" query:"status_id"`
	Pedidos_Id        int64  `json:"pedidos_id" db:"pedidos_id" query:"pedidos_id"`
	Status_Descricao  string `json:"status_descricao" db:"status_descricao" query:"status_descricao"`
	Nome              string `json:"nome" db:"nome" query:"nome"`
	Pedido_Data       string `json:"pedido_data" db:"pedido_data" query:"pedido_data"`
	Descricao         string `json:"descricao" db:"descricao" query:"descricao"`
	Email             string
	EmailParceiro     string
	Imagens_Diretorio string
	QrCode            string
	ID                int64 `query:"id" json:"id"`
}

type Contact struct {
	Title         string
	Nome          string
	NomeParceiro  string
	Telefone      string
	Email         string
	EmailParceiro string
	Assunto       string
	Texto         string
}

type SendBudget struct {
	EmailCliente  string `json:"emailCliente"`
	EmailLoja     string `json:"emailLoja"`
	TextoDaOferta string `json:"textoDaOferta"`
	NomeParceiro  string `json:"nomeParceiro"`
}
