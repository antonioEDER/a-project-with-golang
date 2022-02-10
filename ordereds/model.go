package ordereds

import (
	"time"

	gotoolboxtime "github.com/eucatur/go-toolbox/time"
)

type OrderItens struct {
	ID                            int64  `query:"id" json:"id"`
	Pedidos_Id                    int64  `query:"pedidos_id" json:"pedidos_id"`
	Produtos_Id                   int64  `query:"produtos_id" json:"produtos_id"`
	Valor                         string `query:"valor" json:"valor"`
	Vinculo_Para_Produto_Composto string `query:"vinculo_para_produto_composto" json:"vinculo_para_produto_composto"`
	He_Ativo                      bool   `query:"he_ativo" json:"he_ativo"`
}

type Order struct {
	ID                 int64  `query:"id" json:"id"`
	Pedidos_Id         int64  `query:"pedidos_id" json:"pedidos_id"`
	Pedidos_Status_Ids string `query:"pedidos_status_ids" json:"pedidos_status_ids"`

	Parceiros_Id                  int64      `query:"parceiros_id" json:"parceiros_id"`
	Pessoas_Id                    string     `query:"pessoas_id" json:"pessoas_id"`
	Pessoas_Usuarios_Id           int64      `query:"pessoas_usuarios_id" json:"pessoas_usuarios_id"`
	Vinculo_Para_Produto_Composto string     `query:"vinculo_para_produto_composto" json:"vinculo_para_produto_composto"`
	Pedidos_Itens_Id              int64      `query:"pedidos_itens_id" json:"pedidos_itens_id"`
	Pedidos_Categorias_Id         int64      `query:"pedidos_categorias_id" json:"pedidos_categorias_id"`
	Pedidos_Status_Id             int64      `query:"pedidos_status_id" json:"pedidos_status_id"`
	Enderecos_Id                  int64      `query:"enderecos_id" json:"enderecos_id"`
	Status                        string     `query:"status" json:"status"`
	Valor_Total                   float64    `query:"valor_total" json:"valor_total"`
	Forma_Pagamento               string     `query:"forma_pagamento" json:"forma_pagamento"`
	He_Servico_Pagamento          int64      `query:"he_servico_pagamento" json:"he_servico_pagamento"`
	Troco                         string     `query:"troco" json:"troco"`
	Servico_Pagamento             string     `query:"servico_pagamento" json:"servico_pagamento"`
	Taxa_Entrega                  float64    `query:"taxa_entrega" json:"taxa_entrega"`
	Prazo_Entrega                 string     `query:"prazo_entrega" json:"prazo_entrega"`
	Tipo_Veiculo                  string     `query:"tipo_veiculo" json:"tipo_veiculo"`
	Placa_Veiculo                 string     `query:"placa_veiculo" json:"placa_veiculo"`
	Cor_Veiculo                   string     `query:"cor_veiculo" json:"cor_veiculo"`
	Entregador                    string     `query:"entregador" json:"entregador"`
	Visualizado                   bool       `query:"visualizado" json:"visualizado"`
	He_Ativo                      bool       `query:"he_ativo" json:"he_ativo"`
	Produtos                      []Produtos `query:"produtos" json:"produtos"`
	Produtos_Departamentos_Id     int64      `query:"produtos_departamentos_id" json:"produtos_departamentos_id"`
	Produtos_Marcas_Id            int64      `query:"produtos_marcas_id" json:"produtos_marcas_id"`
	Produtos_Categorias_Id        int64      `query:"produtos_categorias_id" json:"produtos_categorias_id"`
	Pedidos_Status_Descricao      string     `query:"pedidos_status_descricao" json:"pedidos_status_descricao"`
	Imagens_Id                    int64      `query:"imagens_id" json:"imagens_id"`
	Imagens_Descricao             string     `query:"imagens_descricao" json:"imagens_descricao"`
	Imagens_Diretorio             string     `query:"imagens_diretorio" json:"imagens_diretorio"`
	Codigo_De_Barras              string     `query:"codigo_de_barras" json:"codigo_de_barras"`
	Descricao                     string     `query:"descricao" json:"descricao"`
	Valor                         string     `query:"valor" json:"valor"`
	Quantidade                    string     `query:"quantidade" json:"quantidade"`
	Unidade_Medida                string     `query:"unidade_medida" json:"unidade_medida"`
	Video_Incorporado             string     `query:"video_incorporado" json:"video_incorporado"`
	He_Promocao                   int64      `query:"he_promocao" json:"he_promocao"`
	Valor_Promocao                string     `query:"valor_promocao" json:"valor_promocao"`
	He_Combo                      int64      `query:"he_combo" json:"he_combo"`
	He_Acompanhamento             int64      `query:"he_acompanhamento" json:"he_acompanhamento"`
	Informacao_Adicional          string     `query:"informacao_adicional" json:"informacao_adicional"`
	Peso                          string     `query:"peso" json:"peso"`
	Largura                       string     `query:"largura" json:"largura"`
	Altura                        string     `query:"altura" json:"altura"`
	Diametro                      string     `query:"diametro" json:"diametro"`
	Comprimento                   string     `query:"comprimento" json:"comprimento"`
	Produtos_Id_Adicional         int64      `query:"produtos_id_adicional" json:"produtos_id_adicional"`
	Produtos_Combos_Categorias_Id int64      `query:"produtos_combos_categorias_id" json:"produtos_combos_categorias_id"`
	Produtos_Combos_id            int64      `query:"produtos_combos_id" json:"produtos_combos_id"`
	Produtos_Id_Principal         int64      `query:"produtos_id_principal" json:"produtos_id_principal"`
	Consulta                      string     `query:"consulta" json:"consulta"`
	Produto_principal             string     `query:"produto_principal" json:"produto_principal"`
	Departamentos_Descricao       string     `db:"produtos_departamentos_descricao" query:"produtos_departamentos_descricao" json:"produtos_departamentos_descricao"`
	Categorias_Descricao          string     `db:"produtos_categorias_descricao" query:"produtos_categorias_descricao" json:"produtos_categorias_descricao"`
	Marcas_Descricao              string     `db:"produtos_marcas_descricao" query:"produtos_marcas_descricao" json:"produtos_marcas_descricao"`
	Produtos_Favoritos_id         int64      `query:"produtos_favoritos_id" json:"produtos_favoritos_id"`
	Produtos_Id                   int64      `query:"produtos_id" json:"produtos_id"`
	Pedidos_Categorias_Descricao  string     `query:"pedidos_categorias_descricao" json:"pedidos_categorias_descricao"`
	Pedidos_Status_Cor            string     `query:"pedidos_status_cor" json:"pedidos_status_cor"`
	Time_Zone                     string     `query:"time_zone" json:"time_zone"`
	Code_Referencia               string     `query:"code_referencia" json:"code_referencia"`
	CREATED_AT                    time.Time  `query:"created_at" json:"created_at"`

	Nome            string `query:"nome" json:"nome"`
	Cpf             string `query:"cpf" json:"cpf"`
	Data_Nascimento string `query:"data_nascimento" json:"data_nascimento"`
	Celular         string `query:"celular" json:"celular"`
	Telefone        string `query:"telefone" json:"telefone"`
	Email           string `query:"email" json:"email"`
	Code_Checkout   string `query:"code_checkout" json:"code_checkout"`

	Uf          string `db:"uf" json:"uf" `
	Cidade      string `db:"cidade" json:"cidade" `
	Numero      string `db:"numero" json:"numero" `
	Bairro      string `db:"bairro" json:"bairro" `
	Cep         string `db:"cep" json:"cep" `
	Logradouro  string `db:"logradouro" json:"logradouro" `
	Complemento string `db:"complemento" json:"complemento" `

	AnoValidade       string `json:"anoValidade"`
	BrandName         string `json:"brandName"`
	CardNumber        string `json:"cardNumber"`
	CpfCartao         string `json:"cpfCartao"`
	CpfComprador      string `json:"cpfComprador"`
	Cvv               string `json:"cvv"`
	HashCard          string `json:"hashCard"`
	MesValidade       string `json:"mesValidade"`
	NascimentoCartao  string `json:"nascimentoCartao"`
	NomeCartao        string `json:"nomeCartao"`
	TokenCard         string `json:"tokenCard"`
	QtdParcelas       int64  `json:"qtdParcelas"`
	ValorParcelas     string `json:"valorParcelas"`
	TelefoneComprador string `json:"telefoneComprador"`
}

type Produtos struct {
	ID                            int64                `query:"id" json:"id"`
	Vinculo_Para_Produto_Composto string               `query:"vinculo_para_produto_composto" json:"vinculo_para_produto_composto"`
	ProdutosCompostoAdd           []ProdutoCompostoAdd `query:"produtos_composto_add" json:"produtos_composto_add"`
}

type ProdutoCompostoAdd struct {
	Identificador_Adicional string `query:"identificador_adicional" json:"identificador_adicional"`
	Produtos_Categorias_Id  int64  `query:"produtos_categorias_id" json:"produtos_categorias_id"`
	Produtos_Id             int64  `query:"produtos_id" json:"produtos_id"`
}

type OrderCategory struct {
	ID                 int64                   `query:"id" json:"id"`
	Descricao          string                  `query:"descricao" json:"descricao"`
	Tipo               string                  `query:"tipo" json:"tipo"`
	He_Ativo           bool                    `query:"he_ativo" json:"he_ativo"`
	Pedidos_Status_Ids string                  `query:"pedidos_status_ids" json:"pedidos_status_ids"`
	CREATED_AT         gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY         gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT          gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY          gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED            gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type OrderItems struct {
	ID                            int64                   `query:"id" json:"id"`
	Pedidos_Id                    int64                   `query:"pedidos_id" json:"pedidos_id"`
	Produtos_Id                   int64                   `query:"produtos_id" json:"produtos_id"`
	Valor                         bool                    `query:"valor" json:"valor"`
	Vinculo_Para_Produto_Composto bool                    `query:"vinculo_para_produto_composto" json:"vinculo_para_produto_composto"`
	He_Ativo                      bool                    `query:"he_ativo" json:"he_ativo"`
	CREATED_AT                    gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY                    gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT                     gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY                     gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED                       gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type OrderProductCompound struct {
	ID                            int64                   `query:"id" json:"id"`
	Parceiros_Id                  int64                   `query:"parceiros_id" json:"parceiros_id"`
	Pessoas_Usuarios_Id           int64                   `query:"pessoas_usuarios_id" json:"pessoas_usuarios_id"`
	Produtos_Id                   int64                   `query:"produtos_id" json:"produtos_id"`
	Principal_Produtos_Id         int64                   `query:"principal_produtos_id" json:"principal_produtos_id"`
	Pedidos_Id                    int64                   `query:"pedidos_id" json:"pedidos_id"`
	Produtos_Categorias_Id        int64                   `query:"produtos_categorias_id" json:"produtos_categorias_id"`
	Vinculo_Para_Produto_Composto bool                    `query:"vinculo_para_produto_composto" json:"vinculo_para_produto_composto"`
	Identificador_Adicional       bool                    `query:"identificador_adicional" json:"identificador_adicional"`
	He_Ativo                      bool                    `query:"he_ativo" json:"he_ativo"`
	CREATED_AT                    gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY                    gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT                     gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY                     gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED                       gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type OrderStatus struct {
	ID         int64                   `query:"id" json:"id"`
	Descricao  string                  `query:"descricao" json:"descricao"`
	Icone      string                  `query:"icone" json:"icone"`
	Cor        string                  `query:"cor" json:"cor"`
	He_Ativo   bool                    `query:"he_ativo" json:"he_ativo"`
	CREATED_AT gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT  gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY  gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED    gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type ParamsProductComposite struct {
	Pedidos_Id int64 `query:"pedidos_id" json:"pedidos_id"`
	Product
	Pessoas_Id                    string                     `query:"pessoas_id" json:"pessoas_id"`
	Produtos_Id                   string                     `query:"produtos_id" json:"produtos_id"`
	Parceiros_Id                  int64                      `query:"parceiros_id" json:"parceiros_id"`
	Produtos_Id_Principal         int64                      `query:"produtos_id_principal" json:"produtos_id_principal"`
	Principal_Produtos_Id         int64                      `query:"principal_produtos_id" json:"principal_produtos_id"`
	Categorias                    []ProductCompositeCategory `query:"categorias" json:"categorias"`
	He_Ativo                      int64                      `query:"he_ativo" json:"he_ativo"`
	Consulta                      string                     `query:"consulta" json:"consulta"`
	Produto_Principal_Id          string                     `query:"produto_principal_id" json:"produto_principal_id"`
	Imagens_Id                    int64                      `query:"imagens_id" json:"imagens_id"`
	Imagens_Descricao             string                     `query:"imagens_descricao" json:"imagens_descricao"`
	Imagens_Diretorio             string                     `query:"imagens_diretorio" json:"imagens_diretorio"`
	Vinculo_Para_Produto_Composto string                     `query:"vinculo_para_produto_composto" json:"vinculo_para_produto_composto"`
}
type AdditionalIdentifier struct {
	Identificador_Adicional string `query:"identificador_adicional" json:"identificador_adicional"`
}

type ProductCompositeCategory struct {
	ID                                   int64     `query:"id" json:"id"`
	Imagens_Id                           int64     `query:"imagens_id" json:"imagens_id"`
	Imagens_Descricao                    string    `query:"imagens_descricao" json:"imagens_descricao"`
	Imagens_Diretorio                    string    `query:"imagens_diretorio" json:"imagens_diretorio"`
	Vinculo_Para_Produto_Composto        string    `query:"vinculo_para_produto_composto" json:"vinculo_para_produto_composto"`
	Identificador_Adicional              string    `query:"identificador_adicional" json:"identificador_adicional"`
	Parceiros_Id                         int64     `query:"parceiros_id" json:"parceiros_id"`
	Produtos_Categorias_Id               int64     `query:"produtos_categorias_id" json:"produtos_categorias_id"`
	Produtos_Id_Principal                int64     `query:"produtos_id_principal" json:"produtos_id_principal"`
	Principal_Produtos_Id                int64     `query:"principal_produtos_id" json:"principal_produtos_id"`
	Descricao                            string    `query:"descricao" json:"descricao"`
	Tipo_Da_Selecao                      string    `query:"tipo_da_selecao" json:"tipo_da_selecao"`
	Quantidade_Max                       string    `query:"quantidade_max" json:"quantidade_max"`
	He_Obrigatorio_Selecao               int64     `query:"he_obrigatorio_selecao" json:"he_obrigatorio_selecao"`
	He_Adicional                         int64     `query:"he_adicional" json:"he_adicional"`
	He_Ativo                             int64     `query:"he_ativo" json:"he_ativo"`
	Produtos_Combos_Categorias_Descricao string    `query:"produtos_combos_categorias_descricao" json:"produtos_combos_categorias_descricao"`
	Produtos_Categorias                  []Product `query:"produtos_categoria" json:"produtos_categoria"`
}

type Product struct {
	ID                                   int64   `query:"id" json:"id"`
	Produtos_Id                          string  `query:"produtos_id" json:"produtos_id"`
	Parceiros_Id                         int64   `query:"parceiros_id" json:"parceiros_id"`
	Pessoas_Id                           string  `query:"pessoas_id" json:"pessoas_id"`
	Produtos_Departamentos_Id            int64   `query:"produtos_departamentos_id" json:"produtos_departamentos_id"`
	Produtos_Marcas_Id                   int64   `query:"produtos_marcas_id" json:"produtos_marcas_id"`
	Produtos_Categorias_Id               int64   `query:"produtos_categorias_id" json:"produtos_categorias_id"`
	Produtos_Combos_Categorias_Descricao string  `query:"produtos_combos_categorias_descricao" json:"produtos_combos_categorias_descricao"`
	Imagens_Id                           int64   `query:"imagens_id" json:"imagens_id"`
	Codigo_De_Barras                     string  `query:"codigo_de_barras" json:"codigo_de_barras"`
	Descricao                            string  `query:"descricao" json:"descricao"`
	Valor                                string  `query:"valor" json:"valor"`
	Quantidade                           string  `query:"quantidade" json:"quantidade"`
	Unidade_Medida                       string  `query:"unidade_medida" json:"unidade_medida"`
	Video_Incorporado                    string  `query:"video_incorporado" json:"video_incorporado"`
	He_Promocao                          int64   `query:"he_promocao" json:"he_promocao"`
	Valor_Promocao                       string  `query:"valor_promocao" json:"valor_promocao"`
	He_Adicional                         int64   `query:"he_adicional" json:"he_adicional"`
	He_Combo                             int64   `query:"he_combo" json:"he_combo"`
	He_Acompanhamento                    int64   `query:"he_acompanhamento" json:"he_acompanhamento"`
	Informacao_Adicional                 string  `query:"informacao_adicional" json:"informacao_adicional"`
	Peso                                 string  `query:"peso" json:"peso"`
	Largura                              string  `query:"largura" json:"largura"`
	Altura                               string  `query:"altura" json:"altura"`
	Diametro                             string  `query:"diametro" json:"diametro"`
	Comprimento                          string  `query:"comprimento" json:"comprimento"`
	He_Ativo                             int64   `query:"he_ativo" json:"he_ativo"`
	Produtos_Id_Adicional                int64   `query:"produtos_id_adicional" json:"produtos_id_adicional"`
	Produtos_Combos_Categorias_Id        int64   `query:"produtos_combos_categorias_id" json:"produtos_combos_categorias_id"`
	Produtos_Combos_id                   int64   `query:"produtos_combos_id" json:"produtos_combos_id"`
	Produtos_Id_Principal                int64   `query:"produtos_id_principal" json:"produtos_id_principal"`
	Vinculo_Para_Produto_Composto        string  `query:"vinculo_para_produto_composto" json:"vinculo_para_produto_composto"`
	Identificador_Adicional              string  `query:"identificador_adicional" json:"identificador_adicional"`
	Principal_Produtos_Id                int64   `query:"principal_produtos_id" json:"principal_produtos_id"`
	Consulta                             string  `query:"consulta" json:"consulta"`
	Produto_principal                    string  `query:"produto_principal" json:"produto_principal"`
	Departamentos_Descricao              string  `db:"produtos_departamentos_descricao" query:"produtos_departamentos_descricao" json:"produtos_departamentos_descricao"`
	Categorias_Descricao                 string  `db:"produtos_categorias_descricao" query:"produtos_categorias_descricao" json:"produtos_categorias_descricao"`
	Marcas_Descricao                     string  `db:"produtos_marcas_descricao" query:"produtos_marcas_descricao" json:"produtos_marcas_descricao"`
	Imagens_Descricao                    string  `query:"imagens_descricao" json:"imagens_descricao"`
	Qtd_Produto                          string  `query:"qtd_produto" json:"qtd_produto"`
	Imagens_Diretorio                    string  `query:"imagens_diretorio" json:"imagens_diretorio"`
	Produtos_Favoritos_id                int64   `query:"produtos_favoritos_id" json:"produtos_favoritos_id"`
	Produtos_Ids                         []int64 `query:"produtos_ids" json:"produtos_ids"`
}

type SendBudget struct {
	EmailCliente  string `json:"emailCliente"`
	EmailLoja     string `json:"emailLoja"`
	TextoDaOferta string `json:"textoDaOferta"`
	NomeParceiro  string `json:"nomeParceiro"`
}
