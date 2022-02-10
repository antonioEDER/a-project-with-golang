package transports

import (
	gotoolboxtime "github.com/eucatur/go-toolbox/time"
)

type Transport struct {
	ID                          int64                   `query:"id" json:"id"`
	Codigo                      string                  `query:"codigo" json:"codigo"`
	Nome                        string                  `query:"nome" json:"nome"`
	Pessoas_Id                  int64                   `query:"pessoas_id" json:"pessoas_id"`
	Pessoas_Fisicas_Id          int64                   `query:"pessoas_fisicas_id" json:"pessoas_fisicas_id"`
	Pessoas_Contatos_Id         int64                   `query:"pessoas_contatos_id" json:"pessoas_contatos_id"`
	Tipo                        string                  `query:"tipo" json:"tipo"`
	Descricao                   string                  `query:"descricao" json:"descricao"`
	Razao_Social                string                  `query:"razao_social" json:"razao_social"`
	Fantasia                    string                  `query:"fantasia" json:"fantasia"`
	Cnpj                        string                  `query:"cnpj" json:"cnpj"`
	Cpf                         string                  `query:"cpf" json:"cpf"`
	Documento                   string                  `query:"documento" json:"documento"`
	Tipo_Documento              string                  `query:"tipo_documento" json:"tipo_documento"`
	Data_Nascimento             string                  `query:"data_nascimento" json:"data_nascimento"`
	Tipo_Contato                string                  `query:"tipo" json:"tipo_contato"`
	Contato                     string                  `query:"contato" json:"contato"`
	Email                       string                  `query:"email" json:"email"`
	Uf                          string                  `db:"uf" json:"uf" `
	Estado                      string                  `db:"estado" json:"estado" `
	Cidade                      string                  `db:"cidade" json:"cidade" `
	Numero                      string                  `db:"numero" json:"numero" `
	Bairro                      string                  `db:"bairro" json:"bairro" `
	Cep                         string                  `db:"cep" json:"cep" `
	Logradouro                  string                  `db:"logradouro" json:"logradouro" `
	Latitude                    string                  `db:"latitude" json:"latitude" `
	Longitude                   string                  `db:"longitude" json:"longitude" `
	Complemento                 string                  `db:"complemento" json:"complemento" `
	He_Principal                string                  `db:"he_principal" json:"he_principal" `
	He_Principal_Parceiro       string                  `db:"he_principal_parceiro" json:"he_principal_parceiro" `
	Area_Abrangencia            string                  `db:"area_abrangencia" json:"area_abrangencia" `
	Parceiros_Id                int64                   `query:"parceiros_id" json:"parceiros_id"`
	Enderecos_Id                int64                   `query:"enderecos_id" json:"enderecos_id"`
	Usuarios_Id                 int64                   `query:"usuarios_id" json:"usuarios_id"`
	Transportadoras_Id          int64                   `query:"transportadoras_id" json:"transportadoras_id"`
	Prazo                       string                  `query:"prazo" json:"prazo"`
	He_Entrega_Propria          bool                    `query:"he_entrega_propria" json:"he_entrega_propria"`
	Transportadoras_Servicos_Id int64                   `query:"transportadoras_servicos_id" json:"transportadoras_servicos_id"`
	He_Ativo                    bool                    `query:"he_ativo" json:"he_ativo"`
	CREATED_AT                  gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY                  gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT                   gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY                   gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED                     gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
	Valor_Pedido                string                  `query:"valor_pedido" json:"valor_pedido"`
	Frete                       float64                 `query:"frete" json:"frete"`
	Valor                       string                  `query:"valor" json:"valor"`
	Produtos_Ids                string                  `query:"produtos_ids" json:"produtos_ids"`
	Km                          string                  `query:"km" json:"km"`
}

type PartnerTransport struct {
	ID                 int64                   `query:"id" json:"id"`
	Parceiros_Id       int64                   `query:"parceiros_id" json:"parceiros_id"`
	Transportadoras_Id int64                   `query:"transportadoras_id" json:"transportadoras_id"`
	He_Ativo           bool                    `query:"he_ativo" json:"he_ativo"`
	CREATED_AT         gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY         gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT          gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY          gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED            gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type TransportService struct {
	ID                 int64                   `query:"id" json:"id"`
	Parceiros_Id       int64                   `query:"parceiros_id" json:"parceiros_id"`
	Transportadoras_Id int64                   `query:"transportadoras_id" json:"transportadoras_id"`
	Descricao          string                  `query:"descricao" json:"descricao"`
	Codigo             string                  `query:"codigo" json:"codigo"`
	Detalhes           string                  `query:"detalhes" json:"detalhes"`
	Valor              string                  `query:"valor" json:"valor"`
	Prazo              string                  `query:"prazo" json:"prazo"`
	He_Ativo           bool                    `query:"he_ativo" json:"he_ativo"`
	CREATED_AT         gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY         gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT          gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY          gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED            gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type TransportServiceKM struct {
	ID                          int64                   `query:"id" json:"id"`
	Parceiros_Id                int64                   `query:"parceiros_id" json:"parceiros_id"`
	Transportadoras_Servicos_Id int64                   `query:"transportadoras_servicos_id" json:"transportadoras_servicos_id"`
	Km                          string                  `query:"km" json:"km"`
	Frete                       string                  `query:"frete" json:"frete"`
	He_Ativo                    bool                    `query:"he_ativo" json:"he_ativo"`
	CREATED_AT                  gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY                  gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT                   gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY                   gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED                     gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type TransportServiceAmount struct {
	ID                          int64                   `query:"id" json:"id"`
	Parceiros_Id                int64                   `query:"parceiros_id" json:"parceiros_id"`
	Transportadoras_Servicos_Id int64                   `query:"transportadoras_servicos_id" json:"transportadoras_servicos_id"`
	Valor                       string                  `query:"valor" json:"valor"`
	Frete                       string                  `query:"frete" json:"frete"`
	He_Ativo                    bool                    `query:"he_ativo" json:"he_ativo"`
	CREATED_AT                  gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY                  gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT                   gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY                   gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED                     gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}
