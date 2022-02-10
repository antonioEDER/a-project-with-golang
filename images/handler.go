package images

import (
	"github.com/eucatur/go-toolbox/handler"
	"github.com/labstack/echo"
)

func SearchImagesHandler(c echo.Context) (err error) {
	i := *c.Get(handler.PARAMETERS).(*ImageGeneric)

	imagens, err := SearchImages(i)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{
		"imagens": imagens,
	})
}

func CreateImagesHandler(c echo.Context) (err error) {
	i := *c.Get(handler.PARAMETERS).(*ImageGeneric)

	img, err := CreateImage(i, c)
	if err != nil || len(img) == 0 {
		return
	}

	return c.JSON(201, echo.Map{"imagem": img[0]})

}

func DisableImageHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func CreateImageHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func DeleteImagesCatalogHandler(c echo.Context) (err error) {
	i := *c.Get(handler.PARAMETERS).(*ImageGeneric)

	err = DeleteImagesCatalog(i)
	if err != nil {
		return
	}

	return c.JSON(201, "Sucesso")

}
