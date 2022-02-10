package products

import (
	"github.com/api-qop-v2/auth"

	"github.com/eucatur/go-toolbox/handler"

	"github.com/labstack/echo"
)

// AddRoutes adiciona as rotas
func AddRoutes(e *echo.Echo) {
	product := e.Group("v2/product", auth.Middleware)
	product.GET("/partner", SearchProductsForPartnerHandler, handler.MiddlewareBindAndValidate(&ParamsSearch{}))
	product.POST("", CreateProductHandler, handler.MiddlewareBindAndValidate(&Product{}))
	product.PUT("", AlterProductHandler, handler.MiddlewareBindAndValidate(&Product{}))

	product.POST("/brand", CreateBrandHandler, handler.MiddlewareBindAndValidate(&Brand{}))
	product.POST("/category", CreateCategoryHandler, handler.MiddlewareBindAndValidate(&Category{}))
	product.POST("/department", CreateDepartmentHandler, handler.MiddlewareBindAndValidate(&Department{}))

	product.PUT("/brand", AlterBrandHandler, handler.MiddlewareBindAndValidate(&Brand{}))
	product.PUT("/category", AlterCategoryHandler, handler.MiddlewareBindAndValidate(&Category{}))
	product.PUT("/department", AlterDepartmentHandler, handler.MiddlewareBindAndValidate(&Department{}))

	product.GET("/partner/compound", SearchCompoundProductForPartnerHandler, handler.MiddlewareBindAndValidate(&ParamsSearch{}))
	product.PUT("/compound", AlterCompositeProductHandler, handler.MiddlewareBindAndValidate(&ParamsProductComposite{}))
	product.POST("/compound", CreateCompositeProductHandler, handler.MiddlewareBindAndValidate(&ParamsProductComposite{}))

	product.POST("/favorite", CreateFavoriteHandler, handler.MiddlewareBindAndValidate(&Favorite{}))
	product.PUT("/favorite", AlterFavoriteHandler, handler.MiddlewareBindAndValidate(&Favorite{}))
	product.GET("/favorite", SearchFavoritesHandler, handler.MiddlewareBindAndValidate(&Favorite{}))

	product.GET("/to-name", SearchProductsByNameForPartnerHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	product.GET("/to-barcodes", SearchProductByBarCodeHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	product.GET("/to-description", SearchProductsByDescriptionHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	product.GET("/compound/generate-csv", GenerateCsvProductCompoundForPartnerHandler, handler.MiddlewareBindAndValidate(&Filter{}))

	product.GET("/generate-csv", GenerateCsvProductsForPartnerHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	product.GET("/brand", SearchBrandHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	product.GET("/category", SearchCategoryHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	product.GET("/department", SearchDepartmentHandler, handler.MiddlewareBindAndValidate(&Filter{}))
	product.GET("/generate-csv-products", SearchProductToCSVHandler, handler.MiddlewareBindAndValidate(&Product{}))
	product.GET("/generate-csv-compound", SearchProductCompoundToCSVHandler, handler.MiddlewareBindAndValidate(&Product{}))

	product.GET("/proximity-to-category", SearchCategorySuggestionsByProximityHandler, handler.MiddlewareBindAndValidate(&Filter{}))

	public := e.Group("v2/product/public")
	public.GET("/partner", SearchProductsForUSerHandler, handler.MiddlewareBindAndValidate(&ParamsSearch{}))
	public.GET("/departments-brands-categories", SearchDepartmentsBrandsCategoryHandler, handler.MiddlewareBindAndValidate(&Product{}))
	public.GET("/filters", SearchProductsByFiltersHandler, handler.MiddlewareBindAndValidate(&ParamsSearch{}))
	public.GET("/compound-to-id", SearchProductCompositeByIdHandler, handler.MiddlewareBindAndValidate(&ParamsProductComposite{}))
	public.GET("/to-catalog", SearchProductForCatalogHandler, handler.MiddlewareBindAndValidate(&Product{}))

}
