package images

import (
	"github.com/api-qop-v2/auth"
	"github.com/eucatur/go-toolbox/handler"

	"github.com/labstack/echo"
)

// AddRoutes adiciona as rotas
func AddRoutes(e *echo.Echo) {

	auth := e.Group("v2/imagens", auth.Middleware)
	auth.GET("/product", SearchImagesHandler, handler.MiddlewareBindAndValidate(&ImageGeneric{}))
	auth.POST("/product", CreateImagesHandler, handler.MiddlewareBindAndValidate(&ImageGeneric{}))
	auth.DELETE("/product/catalog", DeleteImagesCatalogHandler, handler.MiddlewareBindAndValidate(&ImageGeneric{}))

}
