package images

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/api-qop-v2/common"
	"github.com/api-qop-v2/config"
	"github.com/eucatur/go-toolbox/database"
	"github.com/eucatur/go-toolbox/env"
	"github.com/labstack/echo"
)

func DeleteImagesCatalog(i ImageGeneric) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
	err = DeleteImagesCatalogTx(i, tx)
	if err != nil {
		return
	}
	err = tx.Commit()

	return
}

func SearchImages(i ImageGeneric) (imagens []ImageGeneric, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
	imagens, err = SearchImagesTx(i, tx)
	if err != nil {
		return
	}
	err = tx.Commit()

	return
}

func CreateImage(i ImageGeneric, c echo.Context) (img []ImageGeneric, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	file, err := c.FormFile("file")
	if err != nil {
		return
	}

	i.Codigo_De_Barras = c.FormValue("codigo_de_barras")
	i.Descricao = c.FormValue("descricao")
	i.Filename = c.FormValue("filename")
	fileExtension := path.Ext(file.Filename)
	fileName := fmt.Sprintf(`%s%s`, i.Filename, fileExtension)

	dirBucket := "/imagens_produtos/"

	src, err := file.Open()
	if err != nil {
		return
	}

	defer src.Close()

	// Destination
	PathFileTemp := env.MustString(common.PathFileTemp)

	dst, err := os.Create(PathFileTemp + dirBucket + fileName)
	if err != nil {
		return
	}

	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return
	}

	i.Diretorio = fileName
	id, err := CreateImageTx(i, tx)
	if err != nil {
		return
	}

	var fileImg ImageGeneric
	fileImg.ID = id
	img, err = SearchImagesTx(fileImg, tx)
	if err != nil {
		return
	}

	err = tx.Commit()

	return
}
