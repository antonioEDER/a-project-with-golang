package products

import (
	"fmt"
	"strconv"

	"github.com/api-qop-v2/config"
	"github.com/api-qop-v2/partners"
	"github.com/eucatur/go-toolbox/database"
)

func SearchCategorySuggestionsByProximity() {
	return
}

func CreateBrand(p Brand) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	id, err = CreateBrandTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return

}

func SearchDepartmentsBrandsCategory(idPartner string) (departments []Department, brands []Brand, categorys []Category, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	departments, err = SearchDepartmentsTx(idPartner, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	brands, err = SearchBrandsTx(idPartner, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	categorys, err = SearchCategorysTx(idPartner, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return

}

func GenerateCsvProductsForPartner() {
	return
}

func GenerateCsvProductCompoundForPartner() {
	return
}

func SearchProductsForPartner(p ParamsSearch) (sumProduct int, productsAddImgs []Product, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
	partner, err := partners.SearchPartnerFromPersonIdTx(p.Pessoas_Id, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	p.Parceiros_Id = partner[0].Id

	products, err := SearchProductsForPartnerTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	productsAddImgs = []Product{}
	for _, pd := range products {
		imgsCatalogs, _ := SearchImagensProductsForCatalogTx(pd, tx)
		pd.List_Imagens_Galeria = imgsCatalogs
		productsAddImgs = append(productsAddImgs, pd)
	}

	sumProduct, err = SearchCountProducts(p)
	if err != nil {
		return
	}

	err = tx.Commit()

	return

}

func SearchProductDimensionsByIdForPartner() {
	return
}

func SearchProductsByDescription() {
	return
}

func SearchProductsByNameForPartner() {
	return
}

func SearchProductsByFilters(p ParamsSearch) (products []Product, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	products, err = SearchProductsForUSerFromFiltersTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return

}
func SearchProductCompoundToCSV(p Product) (csv string, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	partner, err := partners.SearchPartnerFromPersonIdTx(p.Pessoas_Id, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	ParceiroID, err := strconv.ParseInt(fmt.Sprint(partner[0].Id), 10, 64)
	if err != nil {
		return
	}

	p.Parceiros_Id = ParceiroID

	products, err := SearchProductCompoundToCSVTx(p, tx)
	if err != nil || len(products) == 0 {
		tx.Rollback()
		return
	}

	csv = ""
	csv += "id,"
	csv += "descricao,"
	csv += "valor,"
	csv += "quantidade,"
	csv += "peso,"
	csv += "unidade_medida,"
	csv += "codigo_de_barras,"
	csv += "valor_promocao,"
	csv += "he_acompanhamento,"
	csv += "he_combo,"
	csv += "\n"

	for _, pt := range products {
		csv += fmt.Sprintf("%#v,", pt.ID)
		csv += fmt.Sprintf("%#v,", pt.Descricao)
		csv += fmt.Sprintf("%#v,", pt.Valor)
		csv += fmt.Sprintf("%#v,", pt.Quantidade)
		csv += fmt.Sprintf("%#v,", pt.Peso)
		csv += fmt.Sprintf("%#v,", pt.Unidade_Medida)
		csv += fmt.Sprintf("%#v,", pt.Codigo_De_Barras)
		csv += fmt.Sprintf("%#v,", pt.Valor_Promocao)
		csv += fmt.Sprintf("%#v,", pt.He_Acompanhamento)
		csv += fmt.Sprintf("%#v,", pt.He_Combo)
		csv += "\n"
	}

	err = tx.Commit()

	return

}

func SearchProductToCSV(p Product) (csv string, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	partner, err := partners.SearchPartnerFromPersonIdTx(p.Pessoas_Id, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	ParceiroID, err := strconv.ParseInt(fmt.Sprint(partner[0].Id), 10, 64)
	if err != nil {
		return
	}

	p.Parceiros_Id = ParceiroID

	products, err := SearchProductToCsvTx(p, tx)
	if err != nil || len(products) == 0 {
		tx.Rollback()
		return
	}

	csv = ""
	csv += "id,"
	csv += "descricao,"
	csv += "valor,"
	csv += "quantidade,"
	csv += "peso,"
	csv += "unidade_medida,"
	csv += "codigo_de_barras,"
	csv += "valor_promocao,"
	csv += "he_acompanhamento,"
	csv += "he_combo,"
	csv += "produtos_departamentos_id,"
	csv += "produtos_departamentos_descricao,"
	csv += "produtos_categorias_id,"
	csv += "produtos_categorias_descricao,"
	csv += "produtos_imagens_id,"
	csv += "imagens_descricao,"
	csv += "produtos_marcas_id,"
	csv += "produtos_marcas_descricao,"
	csv += "\n"

	for _, pt := range products {
		csv += fmt.Sprintf("%#v,", pt.ID)
		csv += fmt.Sprintf("%#v,", pt.Descricao)
		csv += fmt.Sprintf("%#v,", pt.Valor)
		csv += fmt.Sprintf("%#v,", pt.Quantidade)
		csv += fmt.Sprintf("%#v,", pt.Peso)
		csv += fmt.Sprintf("%#v,", pt.Unidade_Medida)
		csv += fmt.Sprintf("%#v,", pt.Codigo_De_Barras)
		csv += fmt.Sprintf("%#v,", pt.Valor_Promocao)
		csv += fmt.Sprintf("%#v,", pt.He_Acompanhamento)
		csv += fmt.Sprintf("%#v,", pt.He_Combo)
		csv += fmt.Sprintf("%#v,", pt.Produtos_Departamentos_Id)
		csv += fmt.Sprintf("%#v,", pt.Departamentos_Descricao)
		csv += fmt.Sprintf("%#v,", pt.Produtos_Categorias_Id)
		csv += fmt.Sprintf("%#v,", pt.Categorias_Descricao)
		csv += fmt.Sprintf("%#v,", pt.Produtos_Imagens_Id)
		csv += fmt.Sprintf("%#v,", pt.Imagens_Descricao)
		csv += fmt.Sprintf("%#v,", pt.Produtos_Marcas_Id)
		csv += fmt.Sprintf("%#v,", pt.Marcas_Descricao)
		csv += "\n"
	}

	err = tx.Commit()

	return

}

func SearchDataforProducs() {
	return
}

func CreateCompositeProduct(p ParamsProductComposite) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	partner, err := partners.SearchPartnerFromPersonIdTx(p.Pessoas_Id, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	partnerID, err := strconv.ParseInt(fmt.Sprint(partner[0].Id), 10, 64)
	if err != nil {
		tx.Rollback()
		return
	}

	p.Parceiros_Id = partnerID

	id, err = CreateCompositeProductTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return

}

func AlterCompositeProduct(p ParamsProductComposite) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	partner, err := partners.SearchPartnerFromPersonIdTx(p.Pessoas_Id, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	partnerID, err := strconv.ParseInt(fmt.Sprint(partner[0].Id), 10, 64)
	if err != nil {
		tx.Rollback()
		return
	}

	p.Parceiros_Id = partnerID

	err = AlterCompositeProductTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	if p.He_Ativo == 1 {
		_, err = CreateCompositeProductTx(p, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}

	err = tx.Commit()

	return

}

func SearchCompoundProduct() {
	return
}

func SearchProductCompositeById(p ParamsProductComposite) (products []ParamsProductComposite, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	products, err = SearchProductCompositeByIdTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return

}

func CreateProduct(p Product) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	partner, err := partners.SearchPartnerFromPersonIdTx(p.Pessoas_Id, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	partnerID, err := strconv.ParseInt(fmt.Sprint(partner[0].Id), 10, 64)
	if err != nil {
		tx.Rollback()
		return
	}

	p.Parceiros_Id = partnerID

	id, err = CreateProductTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	p.ID = id
	for _, productId := range p.Imagens_Galeria_Ids {
		p.Imagens_Id = productId
		CreateImagensProductForCatalogTx(p, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}

	err = tx.Commit()

	return

}

func SearchBrand() {
	return
}

func SearchCategory() {
	return
}

func SearchDepartment() {
	return
}

func CreateDepartment(p Department) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	id, err = CreateDepartmentTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return

}

func CreateCategory(p Category) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	id, err = CreateCategoryTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return

}

func AlterBrand(p Brand) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	err = AlterBrandTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func AlterCategory(p Category) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	err = AlterCategoryTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func AlterDepartment(p Department) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	err = AlterDepartmentTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func SearchProductByBarCode() {
	return
}

func AlterProduct(p Product) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	partner, err := partners.SearchPartnerFromPersonIdTx(p.Pessoas_Id, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	partnerID, err := strconv.ParseInt(fmt.Sprint(partner[0].Id), 10, 64)
	if err != nil {
		tx.Rollback()
		return
	}

	p.Parceiros_Id = partnerID

	err = AlterProductTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	for _, productId := range p.Imagens_Galeria_Ids {
		p.Imagens_Id = productId
		CreateImagensProductForCatalogTx(p, tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}

	err = tx.Commit()

	return

}

func CreateFavorite(f Favorite) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	users, err := SearchUserFromIdPersonTx(f, tx)
	if err != nil {
		return
	}

	f.Pessoas_Usuarios_id = fmt.Sprint(users[0].ID)

	id, err = CreateFavoriteTx(f, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return

}

func AlterFavorite(f Favorite) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	users, err := SearchUserFromIdPersonTx(f, tx)
	if err != nil {
		return
	}

	f.Pessoas_Usuarios_id = fmt.Sprint(users[0].ID)

	err = AlterFavoriteTx(f, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func SearchFavorites(f Favorite) (favorites []Product, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	users, err := SearchUserFromIdPersonTx(f, tx)
	if err != nil {
		return
	}

	f.Pessoas_Usuarios_id = fmt.Sprint(users[0].ID)

	favorites, err = SearchFavoritesTx(f, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}
func SearchCubageOfProductsByDimension() {
	return
}

func SearchProductForCatalog() {
	return
}

func SearchCompoundProductForPartner(p ParamsSearch) (products []ParamsProductComposite, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
	partner, err := partners.SearchPartnerFromPersonIdTx(p.Pessoas_Id, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	p.Parceiros_Id = partner[0].Id

	products, err = SearchCompoundProductForPartnerTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return

}

func SearchProductsForUSer(p ParamsSearch) (products []Product, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	products, err = SearchProductsForUSerTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return

}

func SearchCountProducts(p ParamsSearch) (products int, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	products, err = SearchCountProductsTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return

}

func CalculateCubageProduct(p Product) (cubage CubageProduct, err error) {
	var burdenAmount float64

	var height float64
	var length float64
	var width float64
	var burden float64

	h, _ := strconv.ParseFloat(p.Altura, 64)
	l, _ := strconv.ParseFloat(p.Comprimento, 64)
	w, _ := strconv.ParseFloat(p.Largura, 64)
	b, _ := strconv.ParseFloat(p.Peso, 64)

	height += h
	length += l
	width += w
	burden += b

	burdenAmount += burden

	cubage.Dimensao_Altura = fmt.Sprint(height)
	cubage.Dimensao_Comprimento = fmt.Sprint(length)
	cubage.Dimensao_Largura = fmt.Sprint(width)
	cubage.Dimensao_Peso = fmt.Sprint(burdenAmount)

	return

}

func SearchProductsForUSerFromId(p Product) (products []Product, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	products, err = SearchProductsForUSerFromIdTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return

}

func SearchImagensProductsForCatalog(p Product) (products []ImgCatalog, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	products, err = SearchImagensProductsForCatalogTx(p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return

}
