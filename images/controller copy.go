package images

// import (
// 	"context"
// 	"fmt"
// 	"io"
// 	"os"
// 	"path"
// 	"time"

// 	"cloud.google.com/go/storage"
// 	"github.com/api-qop-v2/common"
// 	"github.com/api-qop-v2/config"
// 	"github.com/eucatur/go-toolbox/database"
// 	"github.com/eucatur/go-toolbox/env"
// 	"github.com/labstack/echo"
// 	"google.golang.org/api/option"
// )

// type Demo struct {
// 	client     *storage.Client
// 	bucketName string
// 	bucket     *storage.BucketHandle
// 	ctx        context.Context
// }

// func (d *Demo) ConnectionCloud(c echo.Context) (err error) {
// 	gcsBucketName := env.MustString(common.GcsBucketName)

// 	ctx := context.Background()

// 	client, err := storage.NewClient(ctx, option.WithCredentialsFile(gcsBucketName))
// 	if err != nil {
// 		return
// 	}
// 	// defer client.Close()

// 	bucket := "media-qop"

// 	d.ctx = ctx
// 	d.bucket = client.Bucket(bucket)
// 	d.client = client
// 	d.bucketName = bucket

// 	return
// }

// func (d *Demo) UploadedFile(fileName string, dirBucket string, c echo.Context) (err error) {
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return err
// 	}
// 	src, err := file.Open()
// 	if err != nil {
// 		return err
// 	}
// 	defer src.Close()

// 	PathFileTemp := env.MustString(common.PathFileTemp)
// 	dst, err := os.Create(PathFileTemp + fileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer dst.Close()
// 	// Copy
// 	_, err = io.Copy(dst, src)
// 	if err != nil {
// 		return err
// 	}

// 	ctx, cancel := context.WithTimeout(d.ctx, time.Second*50)
// 	defer cancel()

// 	sw := d.client.Bucket(d.bucketName).Object(dirBucket + fileName).NewWriter(ctx)

// 	f, err := os.Open(PathFileTemp + fileName)
// 	if err != nil {
// 		return fmt.Errorf("os.Open: %v", err)
// 	}
// 	defer f.Close()

// 	_, err = io.Copy(sw, f)
// 	if err != nil {
// 		return
// 	}

// 	err = sw.Close()
// 	if err != nil {
// 		return
// 	}

// 	return
// }

// func SearchImages(i ImageGeneric) (imagens []ImageGeneric, err error) {
// 	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
// 	imagens, err = SearchImagesTx(i, tx)
// 	if err != nil {
// 		return
// 	}
// 	err = tx.Commit()

// 	return
// }

// func CreateImage(i ImageGeneric, c echo.Context) (img []ImageGeneric, err error) {
// 	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return
// 	}

// 	i.Codigo_De_Barras = c.FormValue("codigo_de_barras")
// 	i.Descricao = c.FormValue("descricao")
// 	i.Filename = c.FormValue("filename")

// 	fileExtension := path.Ext(file.Filename)

// 	fileName := fmt.Sprintf(`%s%s`, i.Filename, fileExtension)

// 	conn := &Demo{}
// 	err = conn.ConnectionCloud(c)
// 	if err != nil {
// 		return
// 	}

// 	err = conn.UploadedFile(fileName, "imagens_produtos/", c)
// 	if err != nil {
// 		return
// 	}

// 	i.Diretorio = fileName
// 	id, err := CreateImageTx(i, tx)
// 	if err != nil {
// 		return
// 	}

// 	var fileImg ImageGeneric
// 	fileImg.ID = id
// 	img, err = SearchImagesTx(fileImg, tx)
// 	if err != nil {
// 		return
// 	}

// 	err = tx.Commit()

// 	return
// }

// func DeleteImagesCatalog(i ImageGeneric) (err error) {
// 	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
// 	err = DeleteImagesCatalogTx(i, tx)
// 	if err != nil {
// 		return
// 	}
// 	err = tx.Commit()

// 	return
// }
