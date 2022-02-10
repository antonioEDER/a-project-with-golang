package products

import (
	"fmt"
	"strings"

	"github.com/api-qop-v2/common"
	"github.com/eucatur/go-toolbox/env"
	"github.com/eucatur/go-toolbox/handler"
	"github.com/eucatur/go-toolbox/jwt"
	"github.com/labstack/echo"
)

func SearchCategorySuggestionsByProximityHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchProductCompoundToCSVHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Product)
	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}
	pessoasId := fmt.Sprint(claims["PessoasId"])

	p.Pessoas_Id = pessoasId

	csv, err := SearchProductCompoundToCSV(p)
	if err != nil {
		return err
	}

	return c.JSON(200, echo.Map{
		"csv": csv,
	})
}

func SearchProductToCSVHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Product)
	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}
	pessoasId := fmt.Sprint(claims["PessoasId"])

	p.Pessoas_Id = pessoasId

	csv, err := SearchProductToCSV(p)
	if err != nil {
		return err
	}

	return c.JSON(200, echo.Map{
		"csv": csv,
	})
}

func SearchDepartmentsBrandsCategoryHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Product)

	departments, brands, categorys, err := SearchDepartmentsBrandsCategory(fmt.Sprint(p.Parceiros_Id))
	if err != nil {
		return err
	}

	return c.JSON(200, echo.Map{
		"departamentos": departments,
		"marcas":        brands,
		"categorias":    categorys,
	})
}

func GenerateCsvProductsForPartnerHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func GenerateCsvProductCompoundForPartnerHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchProductsForPartnerHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*ParamsSearch)
	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}
	pessoasId := fmt.Sprint(claims["PessoasId"])

	p.Pessoas_Id = pessoasId

	sumProduct, products, err := SearchProductsForPartner(p)
	if err != nil {
		return err
	}

	return c.JSON(200, echo.Map{
		"produtos":   products,
		"quantidade": sumProduct,
	})
}

func SearchProductDimensionsByIdForPartnerHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchProductsByDescriptionHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchProductsByNameForPartnerHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchProductsByFiltersHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*ParamsSearch)

	products, err := SearchProductsByFilters(p)
	if err != nil {
		return err
	}

	return c.JSON(200, echo.Map{
		"produtos": products,
	})
}

func SearchDataforProducsHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func CreateCompositeProductHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*ParamsProductComposite)

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}
	pessoasId := fmt.Sprint(claims["PessoasId"])

	p.Pessoas_Id = pessoasId

	id, err := CreateCompositeProduct(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"id": id})

}

func AlterCompositeProductHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*ParamsProductComposite)

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}
	pessoasId := fmt.Sprint(claims["PessoasId"])

	p.Pessoas_Id = pessoasId

	err = AlterCompositeProduct(p)
	if err != nil {
		return
	}

	return c.JSON(201, "Sucesso")

}

func SearchCompoundProductForPartnerHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*ParamsSearch)
	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}
	pessoasId := fmt.Sprint(claims["PessoasId"])

	p.Pessoas_Id = pessoasId

	products, err := SearchCompoundProductForPartner(p)
	if err != nil {
		return err
	}

	return c.JSON(200, echo.Map{
		"produtos": products,
	})
}

func SearchProductCompositeByIdHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*ParamsProductComposite)

	productComposite, err := SearchProductCompositeById(p)
	if err != nil {
		return
	}

	return c.JSON(200, echo.Map{"produto_composto": productComposite})
}

func CreateProductHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Product)

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}
	pessoasId := fmt.Sprint(claims["PessoasId"])

	p.Pessoas_Id = pessoasId

	id, err := CreateProduct(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"id": id})

}

func SearchBrandHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchCategoryHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchDepartmentHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func CreateBrandHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Brand)

	id, err := CreateBrand(p)
	if err != nil {
		return err
	}

	return c.JSON(200, echo.Map{
		"id": id,
	})
}

func CreateDepartmentHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Department)

	id, err := CreateDepartment(p)
	if err != nil {
		return err
	}

	return c.JSON(200, echo.Map{
		"id": id,
	})
}

func CreateCategoryHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Category)

	id, err := CreateCategory(p)
	if err != nil {
		return err
	}

	return c.JSON(200, echo.Map{
		"id": id,
	})
}

func AlterBrandHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Brand)

	err = AlterBrand(p)
	if err != nil {
		return err
	}

	return c.JSON(200, "sucesso")
}

func AlterCategoryHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Category)

	err = AlterCategory(p)
	if err != nil {
		return err
	}

	return c.JSON(200, "sucesso")
}

func AlterDepartmentHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Department)

	err = AlterDepartment(p)
	if err != nil {
		return err
	}

	return c.JSON(200, "sucesso")
}

func SearchProductByBarCodeHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func AlterProductHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Product)

	token := strings.TrimSpace(c.Request().Header.Get("QOP-Api-Token"))
	claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
	if err != nil {
		return c.JSON(401, echo.Map{"erro": "Token inválido"})
	}
	pessoasId := fmt.Sprint(claims["PessoasId"])

	p.Pessoas_Id = pessoasId

	err = AlterProduct(p)
	if err != nil {
		return
	}

	return c.JSON(201, "Sucesso")
}

func CreateFavoriteHandler(c echo.Context) (err error) {
	f := *c.Get(handler.PARAMETERS).(*Favorite)
	idPerson := fmt.Sprint(c.Get("PessoasId"))
	if err != nil {
		return
	}

	f.Pessoas_Id = idPerson

	id, err := CreateFavorite(f)
	if err != nil {
		return
	}

	return c.JSON(200, echo.Map{
		"id": id,
	})
}

func AlterFavoriteHandler(c echo.Context) (err error) {
	f := *c.Get(handler.PARAMETERS).(*Favorite)
	idPerson := fmt.Sprint(c.Get("PessoasId"))
	if err != nil {
		return
	}

	f.Pessoas_Id = idPerson

	err = AlterFavorite(f)
	if err != nil {
		return
	}

	return c.JSON(200, "Sucesso")
}

func SearchFavoritesHandler(c echo.Context) (err error) {
	f := *c.Get(handler.PARAMETERS).(*Favorite)
	idPerson := fmt.Sprint(c.Get("PessoasId"))
	if err != nil {
		return
	}

	f.Pessoas_Id = idPerson

	favorites, err := SearchFavorites(f)
	if err != nil {
		return
	}

	return c.JSON(200, echo.Map{"favoritos": favorites})
}

func SearchCubageOfProductsByDimensionHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchProductForCatalogHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Product)

	products, err := SearchProductsForUSerFromId(p)
	if err != nil {
		return err
	}

	imagems, err := SearchImagensProductsForCatalog(p)
	if err != nil {
		return err
	}

	return c.JSON(200, echo.Map{
		"produtos": products,
		"imagems":  imagems,
	})
}

func SearchProductsForUSerHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*ParamsSearch)

	products, err := SearchProductsForUSer(p)
	if err != nil {
		return err
	}

	return c.JSON(200, echo.Map{
		"produtos": products,
	})
}
