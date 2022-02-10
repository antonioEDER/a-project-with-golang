package partners

import (
	gotoolboxtime "github.com/eucatur/go-toolbox/time"
)

type Partners struct {
	Id                            string                  `db:"id" json:"id" `
	Tipo                          string                  `db:"tipo" json:"tipo" `
	Nome                          string                  `db:"nome" json:"nome" `
	Data_Nascimento               string                  `db:"data_nascimento" json:"data_nascimento" `
	Cpf                           string                  `db:"cpf" json:"cpf" `
	Cnpj                          string                  `db:"cnpj" json:"cnpj" `
	Fantasia                      string                  `db:"fantasia" json:"fantasia" `
	Razao_Social                  string                  `db:"razao_social" json:"razao_social" `
	Email                         string                  `db:"email" json:"email" `
	Usuarios_Id                   string                  `db:"usuarios_id" json:"usuarios_id" `
	Senha                         string                  `db:"senha" json:"senha" `
	Nome_Rede_Social              string                  `db:"rede_social" json:"nome_rede_social" `
	Uid                           string                  `db:"uid" json:"uid" `
	Celular                       string                  `db:"celular" json:"celular" `
	Foto                          string                  `db:"foto" json:"foto" `
	Uf                            string                  `db:"uf" json:"uf" `
	Estado                        string                  `db:"estado" json:"estado" `
	Cidade                        string                  `db:"cidade" json:"cidade" `
	Cidades_Id                    string                  `db:"cidades_id" json:"cidades_id" `
	Numero                        string                  `db:"numero" json:"numero" `
	Bairro                        string                  `db:"bairro" json:"bairro" `
	Cep                           string                  `db:"cep" json:"cep" `
	Logradouro                    string                  `db:"logradouro" json:"logradouro" `
	Latitude                      string                  `db:"latitude" json:"latitude" `
	Longitude                     string                  `db:"longitude" json:"longitude" `
	He_Venda_Produto_Digital      int64                   `db:"he_venda_produto_digital" json:"he_venda_produto_digital" `
	Complemento                   string                  `db:"complemento" json:"complemento" `
	He_Principal                  string                  `db:"he_principal" json:"he_principal" `
	He_Principal_Parceiro         string                  `db:"he_principal_parceiro" json:"he_principal_parceiro" `
	He_Entrega_Propria            string                  `db:"he_entrega_propria" json:"he_entrega_propria" `
	Area_Abrangencia              string                  `db:"area_abrangencia" json:"area_abrangencia" `
	Pessoas_Id                    int64                   `query:"pessoas_id" json:"pessoas_id"`
	Pessoas_Contatos_Id           int64                   `query:"pessoas_contatos_id" json:"pessoas_contatos_id"`
	Pessoas_Fisicas_Id            int64                   `query:"pessoas_fisicas_id" json:"pessoas_fisicas_id"`
	Contato                       string                  `query:"contato" json:"contato"`
	He_Ativo                      bool                    `query:"he_ativo" json:"he_ativo"`
	Enderecos_Id                  string                  `query:"enderecos_id" json:"enderecos_id"`
	Parceiros_Id                  int64                   `query:"parceiros_id" json:"parceiros_id"`
	Ramo_Atividades_Id            int64                   `query:"ramo_atividades_id" json:"ramo_atividades_id"`
	Ramo_Atividades_Tipo          string                  `query:"ramo_atividades_tipo" json:"ramo_atividades_tipo"`
	Img                           string                  `query:"img" json:"img"`
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
	Monitorar_Por_Sse             int64                   `query:"monitorar_por_sse" json:"monitorar_por_sse"`
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
	Mostrar_PagSeguro             bool                    `query:"mostrar_pagSeguro" json:"mostrar_pagSeguro"`
	Mostrar_PicPay                bool                    `query:"mostrar_picPay" json:"mostrar_picPay"`
	Mostrar_Pix                   bool                    `query:"mostrar_pix" json:"mostrar_pix"`
	Comercio_Mais_Proximo         string                  `query:"comercio_mais_proximo" json:"comercio_mais_proximo"`
}

type Plano_Produto struct {
	Id          string `db:"id" json:"id" `
	Qtd_Inicial string `query:"qtd_inicial" json:"qtd_inicial"`
	Qtd_Final   string `query:"qtd_final" json:"qtd_final"`
	Valor       string `query:"valor" json:"valor"`
	Parceiro_Id int64  `query:"parceiro_id" json:"parceiro_id"`
	He_Ativo    bool   `query:"he_ativo" json:"he_ativo"`
}

type Plano_Vendas struct {
	Id            string `db:"id" json:"id" `
	Valor_Inicial string `query:"valor_incial" db:"valor_incial" json:"valor_inicial"`
	Valor_Final   string `query:"valor_final" json:"valor_final"`
	Percentagem   string `query:"percentagem" json:"percentagem"`
	Parceiro_Id   int64  `query:"parceiro_id" json:"parceiro_id"`
	He_Ativo      bool   `query:"he_ativo" json:"he_ativo"`
}

type RangeActivity struct {
	ID         int64                   `query:"id" json:"id"`
	Descricao  string                  `query:"descricao" json:"descricao"`
	Icone      string                  `query:"icone" json:"icone"`
	Cor_Icone  string                  `query:"cor_icone" json:"cor_icone"`
	Background string                  `query:"background" json:"background"`
	He_Ativo   string                  `query:"he_ativo" json:"he_ativo"`
	CREATED_AT gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT  gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY  gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED    gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type ResumoVendas struct {
	ID       int64  `query:"id" json:"id"`
	Fantasia string `query:"fantasia" json:"fantasia"`
	Cnpj     string `query:"cnpj" json:"cnpj"`
	Email    string `query:"email" json:"email"`

	Planos_Produtos string `query:"planos_produtos" json:"planos_produtos"`
	Planos_Vendas   string `query:"planos_vendas" json:"planos_vendas"`
	Total_Orcamento string `query:"total_orcamento" json:"total_orcamento"`
	Total_Vendas    string `query:"total_vendas" json:"total_vendas"`

	Orcamentos_Total string `query:"orcamentos_total" json:"orcamentos_total"`
	Cartoes_Total    string `query:"cartoes_total" json:"cartoes_total"`
	Dinheiro_Total   string `query:"dinheiro_total" json:"dinheiro_total"`
	Valor_Total      string `query:"valor_total" json:"valor_total"`

	Dia                   string `query:"dia" json:"dia"`
	Receber_Em_Casa_Total string `query:"receber_em_casa_total" json:"receber_em_casa_total"`
	Retirar_Na_Loja_Total string `query:"retirar_na_loja_total" json:"retirar_na_loja_total"`
	Descricao             string `query:"descricao" json:"descricao"`
	Qtd                   string `query:"qtd" json:"qtd"`
	Qtd_Produtos          string `query:"qtd_produtos" json:"qtd_produtos"`
}
