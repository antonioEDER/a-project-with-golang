package images

import (
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
)

func findByFilterDBTx(filter Filter, tx *sqlx.Tx) (user string, err error) {

	if filter.ID > 0 && tx != nil {
		user = "Seens we have a user :)"
	} else {
		err = errors.New("There's no user")
	}

	return
}

func SearchImagesTx(i ImageGeneric, tx *sqlx.Tx) (imagens []ImageGeneric, err error) {

	query := `SELECT 
			 id,
			 COALESCE(codigo_de_barras, '') codigo_de_barras ,
			 COALESCE(descricao, '') descricao ,
			 COALESCE(diretorio, '') as diretorio,
			 COALESCE(ativar, 0) as ativar
		FROM public.imagens WHERE ativar = 1`

	and := ` AND upper(imagens.descricao) like '%' || $1 || '%' `
	and += ` AND upper(imagens.codigo_de_barras) = $2 `

	codigoDeBarras := strings.ToUpper(i.Codigo_De_Barras)
	descricao := strings.ToUpper(i.Descricao)

	args := []interface{}{
		descricao,
		codigoDeBarras,
	}

	if i.Codigo_De_Barras == "" && i.Descricao != "" {
		and = ` AND upper(imagens.descricao) like '%' || $1 || '%' `
		args = []interface{}{
			descricao,
		}
	}

	if i.Codigo_De_Barras != "" && i.Descricao == "" {
		and = ` AND upper(imagens.codigo_de_barras) = $1 `
		args = []interface{}{
			codigoDeBarras,
		}
	}

	if i.ID != 0 {
		and = ` AND imagens.id = $1 `
		args = []interface{}{
			i.ID,
		}
	}

	query += and

	err = tx.Select(&imagens, query, args...)
	if err != nil {
		return
	}

	return
}

func CreateImageTx(i ImageGeneric, tx *sqlx.Tx) (id int64, err error) {
	query := `INSERT INTO public.imagens(
		descricao,
		codigo_de_barras,
		diretorio,
		ativar
	    )
		VALUES ($1, $2, $3, $4)
		RETURNING id;
		`

	args := []interface{}{
		i.Descricao,
		i.Codigo_De_Barras,
		i.Diretorio,
		1,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func DeleteImagesCatalogTx(i ImageGeneric, tx *sqlx.Tx) (err error) {
	query := `
	UPDATE public.produtos_imagens
	SET 
	he_ativo=0
	WHERE id = $1;
		`

	args := []interface{}{
		i.Produtos_Imagens_Id,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}
