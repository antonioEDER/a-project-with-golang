package address

import (
	gotoolboxtime "github.com/eucatur/go-toolbox/time"
)

type Address struct {
	Id                    int64                   `query:"id" json:"id"`
	Cidades_Id            int64                   `query:"cidades_id" json:"cidades_id"`
	Pessoas_Id            string                  `query:"pessoas_id" json:"pessoas_id"`
	Cep                   string                  `query:"cep" json:"cep"`
	Logradouro            string                  `query:"logradouro" json:"logradouro"`
	Bairro                string                  `query:"bairro" json:"bairro"`
	Complemento           string                  `query:"complemento" json:"complemento"`
	Numero                string                  `query:"numero" json:"numero"`
	Latitude              string                  `query:"latitude" json:"latitude"`
	Longitude             string                  `query:"longitude" json:"longitude"`
	Area_Abrangencia      string                  `query:"area_abrangencia" json:"area_abrangencia"`
	He_Principal          string                  `query:"he_principal" json:"he_principal"`
	He_Principal_Parceiro string                  `query:"he_principal_parceiro" json:"he_principal_parceiro"`
	CREATED_AT            gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY            gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT             gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY             gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED               gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type City struct {
	Id         int64                   `query:"id" json:"id"`
	Estados_Id int64                   `query:"estados_id" json:"estados_id"`
	Nome       string                  `query:"nome" json:"nome"`
	Cod_Ibge   string                  `query:"cod_ibge" json:"cod_ibge"`
	CREATED_AT gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT  gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY  gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED    gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type State struct {
	Id         int64                   `query:"id" json:"id"`
	Paises_Id  int64                   `query:"paises_id" json:"paises_id"`
	Nome       string                  `query:"nome" json:"nome"`
	Ddd        string                  `query:"ddd" json:"ddd"`
	Uf         string                  `query:"uf" json:"uf"`
	CREATED_AT gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT  gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY  gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED    gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type Country struct {
	ID         int64                   `query:"id" json:"id"`
	Nome       int64                   `query:"nome" json:"nome"`
	Ddi        int64                   `query:"ddi" json:"ddi"`
	CREATED_AT gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT  gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY  gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED    gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type AddressWeb struct {
	Id                    int64  `query:"id" json:"id"`
	Pessoas_Id            string `query:"pessoas_id" json:"pessoas_id"`
	Cep                   string `query:"cep" json:"cep"`
	Logradouro            string `query:"logradouro" json:"logradouro"`
	Bairro                string `query:"bairro" json:"bairro"`
	Complemento           string `query:"complemento" json:"complemento"`
	Numero                string `query:"numero" json:"numero"`
	Latitude              string `query:"latitude" json:"latitude"`
	Longitude             string `query:"longitude" json:"longitude"`
	Cidade                string `query:"cidade" json:"cidade"`
	Uf                    string `query:"uf" json:"uf"`
	Estado                string `query:"estado" json:"estado"`
	Rua                   string `query:"rua" json:"rua"`
	He_Principal          string `query:"he_principal" json:"he_principal"`
	Area_Abrangencia      string `db:"area_abrangencia" json:"area_abrangencia" `
	He_Ativo              bool   `query:"he_ativo" json:"he_ativo"`
	He_Principal_Parceiro string `db:"he_principal_parceiro" json:"he_principal_parceiro" `
}

type PersonWeb struct {
	Id                            string                  `db:"id" json:"id" `
	Tipo                          string                  `db:"tipo" json:"tipo" `
	Pessoas_Usuarios_Id           int64                   `query:"pessoas_usuarios_id" json:"pessoas_usuarios_id"`
	He_Adm                        bool                    `query:"he_adm" json:"he_adm"`
	Nome                          string                  `db:"nome" json:"nome" `
	Data_Nascimento               string                  `db:"data_nascimento" json:"data_nascimento" `
	Cpf                           string                  `db:"cpf" json:"cpf" `
	Cnpj                          string                  `db:"cpf" json:"cnpj" `
	Fantasia                      string                  `db:"fantasia" json:"fantasia" `
	Razao_Social                  string                  `db:"razao_social" json:"razao_social" `
	Email                         string                  `db:"email" json:"email" `
	Senha                         string                  `db:"senha" json:"senha" `
	Nome_Rede_Social              string                  `db:"rede_social" json:"nome_rede_social" `
	Uid                           string                  `db:"uid" json:"uid" `
	Celular                       string                  `db:"celular" json:"celular" `
	Foto                          string                  `db:"foto" json:"foto" `
	Uf                            string                  `db:"uf" json:"uf" `
	Cidade                        string                  `db:"cidade" json:"cidade" `
	Numero                        string                  `db:"numero" json:"numero" `
	Bairro                        string                  `db:"bairro" json:"bairro" `
	Cep                           string                  `db:"cep" json:"cep" `
	Logradouro                    string                  `db:"logradouro" json:"logradouro" `
	Latitude                      string                  `db:"latitude" json:"latitude" `
	Longitude                     string                  `db:"longitude" json:"longitude" `
	Complemento                   string                  `db:"complemento" json:"complemento" `
	He_Principal                  string                  `db:"he_principal" json:"he_principal" `
	He_Plano_Venda                bool                    `db:"he_plano_venda" json:"he_plano_venda" `
	He_Principal_Parceiro         string                  `db:"he_principal_parceiro" json:"he_principal_parceiro" `
	He_Entrega_Propria            bool                    `db:"he_entrega_propria" json:"he_entrega_propria" `
	He_Venda_Produto_Digital      int64                   `db:"he_venda_produto_digital" json:"he_venda_produto_digital" `
	Area_Abrangencia              string                  `db:"area_abrangencia" json:"area_abrangencia" `
	Pessoas_Id                    int64                   `query:"pessoas_id" json:"pessoas_id"`
	Pessoas_Contatos_Id           int64                   `query:"pessoas_contatos_id" json:"pessoas_contatos_id"`
	Pessoas_Fisicas_Id            int64                   `query:"pessoas_fisicas_id" json:"pessoas_fisicas_id"`
	Contato                       string                  `query:"contato" json:"contato"`
	He_Ativo                      bool                    `query:"he_ativo" json:"he_ativo"`
	Enderecos_Id                  string                  `query:"enderecos_id" json:"enderecos_id"`
	Parceiros_Id                  int64                   `query:"parceiros_id" json:"parceiros_id"`
	Ramo_Atividades_Id            int64                   `query:"ramo_atividades_id" json:"ramo_atividades_id"`
	Img                           string                  `query:"img" json:"img"`
	Tipos_Pessoas_Id              string                  `tipos_pessoas_id:"img" json:"tipos_pessoas_id"`
	Responsavel                   string                  `query:"responsavel" json:"responsavel"`
	Segunda                       string                  `query:"segunda" json:"segunda"`
	Terca                         string                  `query:"terca" json:"terca"`
	Quarta                        string                  `query:"quarta" json:"quarta"`
	Quinta                        string                  `query:"quinta" json:"quinta"`
	Sexta                         string                  `query:"sexta" json:"sexta"`
	Sabado                        string                  `query:"sabado" json:"sabado"`
	Domingo                       string                  `query:"domingo" json:"domingo"`
	Email_Pag_Seguro              string                  `query:"email_pag_seguro" json:"email_pag_seguro"`
	Token_Pag_Seguro              string                  `query:"token_pag_seguro" json:"token_pag_seguro"`
	Email_Pic_Pay                 string                  `query:"email_pic_pay" json:"email_pic_pay"`
	Token_Pic_Pay                 string                  `query:"token_pic_pay" json:"token_pic_pay"`
	Chave_Pix                     string                  `query:"chave_pix" json:"chave_pix"`
	Receber_Pedido_Por_Email      int64                   `query:"receber_pedido_por_email" json:"receber_pedido_por_email"`
	Finalizar_Com_Orcamento       int64                   `query:"finalizar_com_orcamento" json:"finalizar_com_orcamento"`
	Finalizar_Com_Receber_Em_Casa int64                   `query:"finalizar_com_receber_em_casa" json:"finalizar_com_receber_em_casa"`
	Finalizar_Com_Retirar_Na_Loja int64                   `query:"finalizar_com_retirar_na_loja" json:"finalizar_com_retirar_na_loja"`
	Usa_Correios                  int64                   `query:"usa_correios" json:"usa_correios"`
	Monitorar_Por_Sse             bool                    `query:"monitorar_por_sse" json:"monitorar_por_sse"`
	He_Pago_Na_Entrega            int64                   `query:"he_pago_na_entrega" json:"he_pago_na_entrega"`
	He_Pago_Pelo_Site             int64                   `query:"he_pago_pelo_site" json:"he_pago_pelo_site"`
	He_Servico_Pagamento          int64                   `query:"he_servico_pagamento" json:"he_servico_pagamento"`
	He_Fechado                    int64                   `query:"he_fechado" json:"he_fechado"`
	CREATED_AT                    gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY                    gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT                     gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY                     gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED                       gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
	Descricao                     string                  `query:"descricao" json:"descricao"`
	Obs                           string                  `query:"obs" json:"obs"`
	Transportadoras_Id            int64                   `query:"transportadoras_id" json:"transportadoras_id"`
	Plano_Produto                 []Plano_Produto         `json:"plano_produtos"`
	Plano_Vendas                  []Plano_Vendas          `json:"plano_vendas"`
	Lista_Transportadoras_Ids     []string                `json:"lista_transportadoras_ids"`
	Token_Push_Web                string                  `json:"token_push_web"`
	Token_Push_App                string                  `json:"token_push_app"`
}

type Plano_Produto struct {
	Id          string `db:"id" json:"id" `
	Qtd_Inicial string `query:"qtd_inicial" json:"qtd_inicial"`
	Qtd_Final   string `query:"qtd_final" json:"qtd_final"`
	Valor       string `query:"valor" json:"valor"`
}

type Plano_Vendas struct {
	Id            string `db:"id" json:"id" `
	Valor_Inicial string `query:"valor_incial" db:"valor_incial" json:"valor_inicial"`
	Valor_Final   string `query:"valor_final" json:"valor_final"`
	Percentagem   string `query:"percentagem" json:"percentagem"`
}
