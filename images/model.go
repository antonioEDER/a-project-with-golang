package images

import (
	gotoolboxtime "github.com/eucatur/go-toolbox/time"
)

type ImageGeneric struct {
	ID                   int64                   `query:"id" json:"id"`
	Produtos_Imagens_Id  int64                   `query:"produtos_imagens_id" json:"produtos_imagens_id"`
	Codigo_De_Barras     string                  `query:"codigo_de_barras" json:"codigo_de_barras"`
	Descricao            string                  `query:"descricao" json:"descricao"`
	Produto_Upper        string                  `query:"produto_upper" json:"produto_upper"`
	Produto_Acento       string                  `query:"produto_acento" json:"produto_acento"`
	Peso                 string                  `query:"peso" json:"peso"`
	Ncm                  string                  `query:"ncm" json:"ncm"`
	Cest_Codigo          string                  `query:"cest_codigo" json:"cest_codigo"`
	Embalagem            string                  `query:"embalagem" json:"embalagem"`
	Quantidade_Embalagem string                  `query:"quantidade_embalagem" json:"quantidade_embalagem"`
	Diretorio            string                  `query:"diretorio" json:"diretorio"`
	Foto_Web             string                  `query:"foto_web" json:"foto_web"`
	Marca                string                  `query:"marca" json:"marca"`
	Preco_Medio          string                  `query:"preco_medio" json:"preco_medio"`
	Img_Gtin             string                  `query:"img_gtin" json:"img_gtin"`
	Categoria            string                  `query:"categoria" json:"categoria"`
	Ativar               bool                    `query:"ativar" json:"ativar"`
	CREATED_AT           gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY           gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT            gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY            gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED              gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
	Filename             string                  `query:"filename" json:"filename"`
}
