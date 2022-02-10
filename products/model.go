package products

type ImageProduct struct {
	ID          int64 `query:"id" json:"id"`
	Produtos_Id int64 `query:"produtos_id" json:"produtos_id"`
	Imagens_Id  int64 `query:"imagens_id" json:"imagens_id"`
	He_Ativo    bool  `query:"he_ativo" json:"he_ativo"`
}

type CubageProduct struct {
	Dimensao_Altura      string `query:"dimensao_altura" json:"dimensao_altura"`
	Dimensao_Comprimento string `query:"dimensao_comprimento" json:"dimensao_comprimento"`
	Dimensao_Largura     string `query:"dimensao_largura" json:"dimensao_largura"`
	Dimensao_Peso        string `query:"dimensao_peso" json:"dimensao_peso"`
}

type Product struct {
	ID                            int64        `query:"id" json:"id"`
	Parceiros_Id                  int64        `query:"parceiros_id" json:"parceiros_id"`
	Pessoas_Id                    string       `query:"pessoas_id" json:"pessoas_id"`
	Produtos_Departamentos_Id     int64        `query:"produtos_departamentos_id" json:"produtos_departamentos_id"`
	Produtos_Marcas_Id            int64        `query:"produtos_marcas_id" json:"produtos_marcas_id"`
	Produtos_Categorias_Id        int64        `query:"produtos_categorias_id" json:"produtos_categorias_id"`
	Imagens_Id                    int64        `query:"imagens_id" json:"imagens_id"`
	Imagens_Descricao             string       `query:"imagens_descricao" json:"imagens_descricao"`
	Imagens_Diretorio             string       `query:"imagens_diretorio" json:"imagens_diretorio"`
	Codigo_De_Barras              string       `query:"codigo_de_barras" json:"codigo_de_barras"`
	Descricao                     string       `query:"descricao" json:"descricao"`
	Valor                         string       `query:"valor" json:"valor"`
	Quantidade                    string       `query:"quantidade" json:"quantidade"`
	Unidade_Medida                string       `query:"unidade_medida" json:"unidade_medida"`
	Video_Incorporado             string       `query:"video_incorporado" json:"video_incorporado"`
	He_Promocao                   int64        `query:"he_promocao" json:"he_promocao"`
	Valor_Promocao                string       `query:"valor_promocao" json:"valor_promocao"`
	He_Combo                      int64        `query:"he_combo" json:"he_combo"`
	He_Acompanhamento             int64        `query:"he_acompanhamento" json:"he_acompanhamento"`
	Informacao_Adicional          string       `query:"informacao_adicional" json:"informacao_adicional"`
	Peso                          string       `query:"peso" json:"peso"`
	Largura                       string       `query:"largura" json:"largura"`
	Altura                        string       `query:"altura" json:"altura"`
	Diametro                      string       `query:"diametro" json:"diametro"`
	Comprimento                   string       `query:"comprimento" json:"comprimento"`
	He_Ativo                      int64        `query:"he_ativo" json:"he_ativo"`
	Produtos_Id_Adicional         int64        `query:"produtos_id_adicional" json:"produtos_id_adicional"`
	Produtos_Combos_Categorias_Id int64        `query:"produtos_combos_categorias_id" json:"produtos_combos_categorias_id"`
	Produtos_Combos_id            int64        `query:"produtos_combos_id" json:"produtos_combos_id"`
	Produtos_Id_Principal         int64        `query:"produtos_id_principal" json:"produtos_id_principal"`
	Consulta                      string       `query:"consulta" json:"consulta"`
	Produto_principal             string       `query:"produto_principal" json:"produto_principal"`
	Departamentos_Descricao       string       `db:"produtos_departamentos_descricao" query:"produtos_departamentos_descricao" json:"produtos_departamentos_descricao"`
	Categorias_Descricao          string       `db:"produtos_categorias_descricao" query:"produtos_categorias_descricao" json:"produtos_categorias_descricao"`
	Marcas_Descricao              string       `db:"produtos_marcas_descricao" query:"produtos_marcas_descricao" json:"produtos_marcas_descricao"`
	Produtos_Favoritos_id         int64        `query:"produtos_favoritos_id" json:"produtos_favoritos_id"`
	Produtos_Imagens_Id           int64        `query:"produtos_imagens_id" json:"produtos_imagens_id"`
	Count_Produtos                int          `query:"count_produtos" json:"count_produtos"`
	Produtos_Ids                  []int64      `query:"produtos_ids" json:"produtos_ids"`
	Imagens_Galeria_Ids           []int64      `query:"imagens_galeria_ids" json:"imagens_galeria_ids"`
	List_Imagens_Galeria          []ImgCatalog `query:"list_imagens_galeria" json:"list_imagens_galeria"`
}

type ImgCatalog struct {
	Produtos_Id         int64  `query:"produtos_id" json:"produtos_id"`
	Produtos_Imagens_id int64  `query:"produtos_imagens_id" json:"produtos_imagens_id"`
	Imagens_Id          int64  `query:"imagens_id" json:"imagens_id"`
	Video_Incorporado   string `query:"video_incorporado" json:"video_incorporado"`
	Imagens_Descricao   string `query:"imagens_descricao" json:"imagens_descricao"`
	Imagens_Diretorio   string `query:"imagens_diretorio" json:"imagens_diretorio"`
}

type RegistrationData struct {
	Brand      []Brand
	Department []Department
	Category   []Category
}

//marca
type Brand struct {
	ID           int64  `query:"id" json:"id"`
	Parceiros_Id string `query:"parceiros_id" json:"parceiros_id"`
	Descricao    string `query:"descricao" json:"descricao"`
	Img          string `query:"img" json:"img"`
	He_Ativo     bool   `query:"he_ativo" json:"he_ativo"`
}

// Departamento
type Department struct {
	ID           int64  `query:"id" json:"id"`
	Parceiros_Id string `query:"parceiros_id" json:"parceiros_id"`
	Descricao    string `query:"descricao" json:"descricao"`
	Img          string `query:"img" json:"img"`
	He_Ativo     bool   `query:"he_ativo" json:"he_ativo"`
}

// Categoria
type Category struct {
	ID                        int64  `query:"id" json:"id"`
	Parceiros_Id              string `query:"parceiros_id" json:"parceiros_id"`
	Produtos_Departamentos_Id int64  `query:"produtos_departamentos_id" json:"produtos_departamentos_id"`
	Descricao                 string `query:"descricao" json:"descricao"`
	Img                       string `query:"img" json:"img"`
	He_Ativo                  bool   `query:"he_ativo" json:"he_ativo"`
}

type ParamsSearch struct {
	Produtos_Id       int64  `query:"produtos_id" json:"produtos_id"`
	Pessoas_Id        string `query:"pessoas_id" json:"pessoas_id"`
	Parceiros_Id      string `query:"parceiros_id" json:"parceiros_id"`
	Codigo_De_Barras  string `query:"codigo_de_barras" json:"codigo_de_barras"`
	Descricao         string `query:"descricao" json:"descricao"`
	Condicao          string `query:"condicao" json:"condicao"`
	He_Ativo          bool   `query:"he_ativo" json:"he_ativo"`
	Pagina            int    `query:"pagina" json:"pagina"`
	Categorias_Ids    string `query:"categorias_ids" json:"categorias_ids"`
	Departamentos_Ids string `query:"departamentos_ids" json:"departamentos_ids"`
	Marcas_Ids        string `query:"marcas_ids" json:"marcas_ids"`
}

type ProductComposite struct {
	ID                            int64 `query:"id" json:"id"`
	Parceiros_Id                  int64 `query:"parceiros_id" json:"parceiros_id"`
	Produtos_Id_Principal         int64 `query:"produtos_id_principal" json:"produtos_id_principal"`
	Produtos_Id_Adicional         int64 `query:"produtos_id_adicional" json:"produtos_id_adicional"`
	Produtos_Combos_Categorias_Id int64 `query:"produtos_combos_categorias_id" json:"produtos_combos_categorias_id"`
	He_Ativo                      int64 `query:"he_ativo" json:"he_ativo"`
}

type ProductCompositeCategory struct {
	ID                     int64     `query:"id" json:"id"`
	Parceiros_Id           int64     `query:"parceiros_id" json:"parceiros_id"`
	Produtos_Id_Principal  int64     `query:"produtos_id_principal" json:"produtos_id_principal"`
	Principal_Produtos_Id  int64     `query:"principal_produtos_id" json:"principal_produtos_id"`
	Descricao              string    `query:"descricao" json:"descricao"`
	Tipo_Da_Selecao        string    `query:"tipo_da_selecao" json:"tipo_da_selecao"`
	Quantidade_Max         string    `query:"quantidade_max" json:"quantidade_max"`
	He_Obrigatorio_Selecao int64     `query:"he_obrigatorio_selecao" json:"he_obrigatorio_selecao"`
	He_Adicional           int64     `query:"he_adicional" json:"he_adicional"`
	He_Ativo               int64     `query:"he_ativo" json:"he_ativo"`
	Produtos_Categorias    []Product `query:"produtos_categoria" json:"produtos_categoria"`
}

type ParamsProductComposite struct {
	Product
	Pessoas_Id            string                     `query:"pessoas_id" json:"pessoas_id"`
	Produtos_Id           string                     `query:"produtos_id" json:"produtos_id"`
	Parceiros_Id          int64                      `query:"parceiros_id" json:"parceiros_id"`
	Produtos_Id_Principal int64                      `query:"produtos_id_principal" json:"produtos_id_principal"`
	Principal_Produtos_Id int64                      `query:"principal_produtos_id" json:"principal_produtos_id"`
	Categorias            []ProductCompositeCategory `query:"categorias" json:"categorias"`
	He_Ativo              int64                      `query:"he_ativo" json:"he_ativo"`
	Consulta              string                     `query:"consulta" json:"consulta"`
	Produto_Principal_Id  string                     `query:"produto_principal_id" json:"produto_principal_id"`
}

type Favorite struct {
	ID                  int64  `query:"id" json:"id"`
	Produtos_Id         int64  `query:"produtos_id" json:"produtos_id"`
	Pessoas_Usuarios_id string `query:"pessoas_usuarios_id" json:"pessoas_usuarios_id"`
	He_Ativo            int64  `query:"he_ativo" json:"he_ativo"`
	Pessoas_Id          string `query:"pessoas_id" json:"pessoas_id"`
}
