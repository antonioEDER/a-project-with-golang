package products

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

const queryProductPublic = `
SELECT
	parceiros.id as parceiros_id,
	produtos.id,
	COALESCE(imagens.id, 0) AS imagens_id,
	COALESCE(imagens.descricao, '') AS imagens_descricao,
	COALESCE(imagens.diretorio, '') AS imagens_diretorio,
	produtos.produtos_departamentos_id,
	produtos.produtos_marcas_id,
	COALESCE(produtos.produtos_categorias_id, 0) as produtos_categorias_id,
	produtos.codigo_de_barras,
	produtos.descricao,
	produtos.valor,
	produtos.quantidade,
	produtos.unidade_medida,
	COALESCE(produtos.video_incorporado, '') as video_incorporado,
	
	produtos.he_ativo,
	produtos.he_promocao,
	COALESCE(produtos.valor_promocao, 0) as valor_promocao,
	COALESCE(produtos.he_combo, 0) as he_combo,
	produtos.he_acompanhamento,
	COALESCE(produtos.informacao_adicional, '') as informacao_adicional,
    COALESCE(produtos.peso, '0') AS peso,
	COALESCE(produtos.largura, 0) as largura,
	COALESCE(produtos.altura, 0) as altura,
	COALESCE(produtos.comprimento, 0) as comprimento,
	produtos_departamentos.descricao as produtos_departamentos_descricao,
	produtos_categorias.descricao as produtos_categorias_descricao,
	produtos_marcas.descricao as produtos_marcas_descricao

FROM public.produtos
JOIN produtos_departamentos ON  produtos_departamentos.id = produtos.produtos_departamentos_id
LEFT JOIN produtos_categorias ON produtos_categorias.id = produtos.produtos_categorias_id
LEFT JOIN produtos_marcas ON  produtos_marcas.id = produtos.produtos_marcas_id
LEFT JOIN imagens ON produtos.imagens_id = imagens.id
JOIN parceiros ON produtos.parceiros_id = parceiros.id `

const queryProduct = ` SELECT 
	COALESCE(imagens.id, 0) AS imagens_id,
	COALESCE(imagens.descricao, '') AS imagens_descricao,
	COALESCE(imagens.diretorio, '') AS imagens_diretorio,
	produtos.id,
	produtos.parceiros_id,
	produtos.produtos_departamentos_id,
	produtos.produtos_marcas_id,
	COALESCE(produtos.produtos_categorias_id, 0) as produtos_categorias_id,
	produtos.codigo_de_barras,
	produtos.descricao,
	produtos.valor,
	produtos.quantidade,
	produtos.unidade_medida,
	COALESCE(produtos.video_incorporado, '') as video_incorporado,
	produtos.he_ativo,
	produtos.he_promocao,
	COALESCE(produtos.valor_promocao, 0) as valor_promocao,
	COALESCE(produtos.he_combo, 0) as he_combo,
	produtos.he_acompanhamento,
	COALESCE(produtos.informacao_adicional, '') as informacao_adicional,
    COALESCE(produtos.peso, '0') AS peso,
	COALESCE(produtos.largura, 0) as largura,
	COALESCE(produtos.altura, 0) as altura,
	COALESCE(produtos.comprimento, 0) as comprimento
	FROM public.produtos 
	LEFT JOIN imagens ON produtos.imagens_id = imagens.id `

func SearchDepartmentsTx(idPartner string, tx *sqlx.Tx) (departments []Department, err error) {

	query := `
	SELECT id, parceiros_id, descricao, he_ativo
	FROM public.produtos_departamentos
		WHERE 
		parceiros_id = $1;
		`

	args := []interface{}{
		idPartner,
	}

	err = tx.Select(&departments, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchBrandsTx(idPartner string, tx *sqlx.Tx) (brands []Brand, err error) {

	query := `
	SELECT id, parceiros_id, descricao, he_ativo
		FROM public.produtos_marcas
		WHERE 
		parceiros_id = $1;
		`

	args := []interface{}{
		idPartner,
	}

	err = tx.Select(&brands, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchCategorysTx(idPartner string, tx *sqlx.Tx) (categorys []Category, err error) {

	query := `
	SELECT id, parceiros_id, produtos_departamentos_id, descricao, he_ativo
		FROM public.produtos_categorias
		WHERE 
		parceiros_id = $1;
		`

	args := []interface{}{
		idPartner,
	}

	err = tx.Select(&categorys, query, args...)

	if err != nil {
		return
	}

	return
}

func CreateBrandTx(p Brand, tx *sqlx.Tx) (id int64, err error) {

	query := `
		INSERT INTO public.produtos_marcas(
			parceiros_id, descricao, he_ativo)
		VALUES ($1, $2, $3)
		RETURNING id;
		`

	args := []interface{}{
		p.Parceiros_Id,
		p.Descricao,
		1,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func CreateCategoryTx(p Category, tx *sqlx.Tx) (id int64, err error) {

	query := `
		INSERT INTO public.produtos_categorias(
			parceiros_id, produtos_departamentos_id, descricao, he_ativo)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
		`

	args := []interface{}{
		p.Parceiros_Id,
		p.Produtos_Departamentos_Id,
		p.Descricao,
		1,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func CreateDepartmentTx(p Department, tx *sqlx.Tx) (id int64, err error) {

	query := `
		INSERT INTO public.produtos_departamentos(
			parceiros_id, descricao, he_ativo)
	   	VALUES ($1, $2, $3)
		RETURNING id;
		`

	args := []interface{}{
		p.Parceiros_Id,
		p.Descricao,
		1,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func AlterBrandTx(p Brand, tx *sqlx.Tx) (err error) {

	query := `
		UPDATE public.produtos_marcas
			SET descricao=$1
		WHERE id = $2;
		`

	args := []interface{}{
		p.Descricao,
		p.ID,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func AlterCategoryTx(p Category, tx *sqlx.Tx) (err error) {

	query := `
		UPDATE public.produtos_categorias
			SET descricao=$1
		WHERE id = $2;
		`

	args := []interface{}{
		p.Descricao,
		p.ID,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func AlterDepartmentTx(p Department, tx *sqlx.Tx) (err error) {

	query := `
		UPDATE public.produtos_departamentos
			SET descricao=$1
		WHERE id = $2;
		`

	args := []interface{}{
		p.Descricao,
		p.ID,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func CreateCompositeProductTx(p ParamsProductComposite, tx *sqlx.Tx) (id int64, err error) {

	queryUpProd := `UPDATE public.produtos SET  he_combo = $1 WHERE id = $2`
	argsUpProd := []interface{}{
		1,
		p.Produtos_Id_Principal,
	}

	_, err = tx.Exec(queryUpProd, argsUpProd...)
	if err != nil {
		return
	}

	for _, category := range p.Categorias {
		queryCat := `
		INSERT INTO public.produtos_combos_categorias(
			parceiros_id,
			principal_produtos_id,
			descricao,
			tipo_da_selecao,
			quantidade_max,
			he_obrigatorio_selecao,
			he_adicional,
			he_ativo)
			VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8)
		RETURNING id;
		`
		argsCat := []interface{}{
			p.Parceiros_Id,
			category.Produtos_Id_Principal,
			category.Descricao,
			category.Tipo_Da_Selecao,
			category.Quantidade_Max,
			category.He_Obrigatorio_Selecao,
			category.He_Adicional,
			category.He_Ativo,
		}

		var idCategory int64
		err = tx.Get(&idCategory, queryCat, argsCat...)
		if err != nil {
			return
		}

		for _, products := range category.Produtos_Categorias {
			queryProd := `
			INSERT INTO public.produtos_combos(
				parceiros_id,
				produtos_id_principal,
				produtos_id_adicional,
				produtos_combos_categorias_id,
				he_ativo)
				VALUES (
				$1,
				$2,
				$3,
				$4,
				$5)
			RETURNING id;
			`
			argsProd := []interface{}{
				p.Parceiros_Id,
				category.Produtos_Id_Principal,
				products.ID,
				idCategory,
				products.He_Ativo,
			}

			err = tx.Get(&id, queryProd, argsProd...)
			if err != nil {
				return
			}
		}
	}

	return id, err
}

func SearchProductCompoundToCSVTx(p Product, tx *sqlx.Tx) (pts []Product, err error) {

	query := `SELECT
    COALESCE(produtos.id, 0) AS id,
    replace(COALESCE(produtos.descricao, '_'), ',', '.') AS descricao,
    COALESCE(produtos.valor, '00.00') AS valor,
    COALESCE(produtos.quantidade, '0') AS quantidade,
    COALESCE(produtos.peso, '0') AS peso,
    COALESCE(produtos.unidade_medida, '_') AS unidade_medida,
    COALESCE(produtos.codigo_de_barras, '_') AS codigo_de_barras,
    COALESCE(produtos.valor_promocao, '00.00') AS valor_promocao,
    COALESCE(produtos.he_acompanhamento, '0') AS he_acompanhamento,
    COALESCE(produtos.he_combo, '0') AS he_combo
    FROM
	produtos_combos
		JOIN produtos ON produtos.id = produtos_combos.produtos_id_principal 
		JOIN imagens ON imagens.id = produtos.imagens_id
        JOIN parceiros ON produtos.parceiros_id = parceiros.id
	WHERE parceiros.id = $1
	GROUP BY produtos.id`

	args := []interface{}{
		p.Parceiros_Id,
	}

	err = tx.Select(&pts, query, args...)

	if err != nil {
		return
	}

	return pts, err
}

func SearchProductToCsvTx(p Product, tx *sqlx.Tx) (pts []Product, err error) {

	query := `SELECT
    COALESCE(produtos.id, 0) AS id,
    replace(COALESCE(produtos.descricao, '_'), ',', '.') AS descricao,
    COALESCE(produtos.valor, '00.00') AS valor,
    COALESCE(produtos.quantidade, '0') AS quantidade,
    COALESCE(produtos.peso, '0') AS peso,
    COALESCE(produtos.unidade_medida, '_') AS unidade_medida,
    COALESCE(produtos.codigo_de_barras, '_') AS codigo_de_barras,
    COALESCE(produtos.valor_promocao, '00.00') AS valor_promocao,
    COALESCE(produtos.he_acompanhamento, '0') AS he_acompanhamento,
    COALESCE(produtos.he_combo, '0') AS he_combo,
    COALESCE(produtos_departamentos.id, 0) AS produtos_departamentos_id,
    COALESCE(produtos_departamentos.descricao, '_') AS produtos_departamentos_descricao,
    COALESCE(produtos_categorias.id, 0) AS produtos_combos_categorias_id,
    COALESCE(produtos_categorias.descricao, '_') AS produtos_categorias_descricao,
    COALESCE(imagens.id, 0) AS produtos_imagens_id,
    COALESCE(imagens.descricao, '_') AS imagens_descricao,
    COALESCE(produtos_marcas.id, 0) AS produtos_marcas_id,
    COALESCE(produtos_marcas.descricao, '_') AS produtos_marcas_descricao
    FROM
    produtos
        JOIN
    produtos_departamentos ON produtos_departamentos.id = produtos.produtos_departamentos_id
        LEFT JOIN
    produtos_categorias ON produtos_categorias.id = produtos.produtos_categorias_id
        LEFT JOIN
    produtos_marcas ON produtos_marcas.id = produtos.produtos_marcas_id
        JOIN
    imagens ON produtos.imagens_id = imagens.id
        JOIN
    parceiros ON produtos.parceiros_id = parceiros.id
	
	WHERE parceiros.id = $1`

	args := []interface{}{
		p.Parceiros_Id,
	}

	err = tx.Select(&pts, query, args...)

	if err != nil {
		return
	}

	return pts, err
}

func CreateImagensProductForCatalogTx(p Product, tx *sqlx.Tx) (id int64, err error) {

	query := `INSERT INTO public.produtos_imagens(produtos_id, imagens_id, he_ativo)
	   VALUES ($1, $2, $3)
		RETURNING id;`

	args := []interface{}{
		p.ID,
		p.Imagens_Id,
		1,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func CreateProductTx(p Product, tx *sqlx.Tx) (id int64, err error) {

	query := `INSERT INTO public.produtos(
		parceiros_id,
		produtos_departamentos_id,
		produtos_marcas_id,
		produtos_categorias_id,
		codigo_de_barras,
		descricao,
		valor,
		quantidade,
		unidade_medida,
		video_incorporado,
		he_ativo,
		he_promocao,
		valor_promocao,
		he_acompanhamento,
		informacao_adicional,
		peso,
		largura,
		altura,
		comprimento,
		imagens_id)
	   VALUES (
		   	$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12,
			$13,
			$14,
			$15,
			$16,
			$17,
			$18,
			$19,
			$20
		)
		RETURNING id;
		`

	args := []interface{}{
		p.Parceiros_Id,
		p.Produtos_Departamentos_Id,
		p.Produtos_Marcas_Id,
		p.Produtos_Categorias_Id,
		p.Codigo_De_Barras,
		p.Descricao,
		p.Valor,
		p.Quantidade,
		p.Unidade_Medida,
		p.Video_Incorporado,
		1,
		p.He_Promocao,
		p.Valor_Promocao,
		p.He_Acompanhamento,
		p.Informacao_Adicional,
		p.Peso,
		p.Largura,
		p.Altura,
		p.Comprimento,
		p.Imagens_Id,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func AlterProductTx(p Product, tx *sqlx.Tx) (err error) {

	query := `
	UPDATE public.produtos
	  SET
		produtos_departamentos_id = $1,
		produtos_marcas_id = $2,
		produtos_categorias_id = $3,
		codigo_de_barras = $4,
		descricao = $5,
		valor = $6,
		quantidade = $7,
		unidade_medida = $8,
		video_incorporado = $9,
		he_ativo = $10,
		he_promocao = $11,
		valor_promocao = $12,
		he_acompanhamento = $13,
		informacao_adicional = $14,
		peso = $15,
		largura = $16,
		altura = $17,
		comprimento = $18,
		imagens_id = $19
		WHERE id = $20 AND parceiros_id = $21`

	args := []interface{}{
		p.Produtos_Departamentos_Id,
		p.Produtos_Marcas_Id,
		p.Produtos_Categorias_Id,
		p.Codigo_De_Barras,
		p.Descricao,
		p.Valor,
		p.Quantidade,
		p.Unidade_Medida,
		p.Video_Incorporado,
		1,
		p.He_Promocao,
		p.Valor_Promocao,
		p.He_Acompanhamento,
		p.Informacao_Adicional,
		p.Peso,
		p.Largura,
		p.Altura,
		p.Comprimento,
		p.Imagens_Id,
		p.ID,
		p.Parceiros_Id,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func SearchProductsForPartnerTx(p ParamsSearch, tx *sqlx.Tx) (products []Product, err error) {

	query := queryProduct

	pgEnd := 0
	offset := p.Pagina

	args := []interface{}{}

	codigoDeBarras := strings.ToUpper(p.Codigo_De_Barras)
	descricao := strings.ToUpper(p.Descricao)

	if p.Descricao == "" && p.Codigo_De_Barras == "" {
		query += ` WHERE produtos.parceiros_id = $1 GROUP BY imagens.id, produtos.id LIMIT 10 OFFSET $2 `
		if offset > 0 {
			pgEnd = offset * 10
		}
		args = []interface{}{
			p.Parceiros_Id,
			pgEnd,
		}
	} else if p.Descricao != "" {
		query += ` WHERE produtos.parceiros_id = $1 AND upper(produtos.descricao) like '%' || $2 || '%' GROUP BY imagens.id, produtos.id`
		args = []interface{}{
			p.Parceiros_Id,
			descricao,
		}
	} else if p.Codigo_De_Barras != "" {
		query += ` WHERE produtos.parceiros_id = $1 AND upper(produtos.codigo_de_barras) like '%' || $2 || '%' GROUP BY imagens.id, produtos.id`
		args = []interface{}{
			p.Parceiros_Id,
			codigoDeBarras,
		}
	}

	err = tx.Select(&products, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchCompoundProductForPartnerTx(p ParamsSearch, tx *sqlx.Tx) (products []ParamsProductComposite, err error) {

	var compounds []ParamsProductComposite

	queryProdutosCombos := `
	SELECT * FROM (
		SELECT 
		(
			SELECT DISTINCT produtos_id_principal FROM produtos_combos WHERE produtos.id = produtos_combos.produtos_id_principal) AS produto_principal_id,
			 produtos.id,
			 produtos.parceiros_id,
			 produtos.produtos_departamentos_id,
			 produtos.produtos_marcas_id,

			 COALESCE(produtos.produtos_categorias_id, 0) as produtos_categorias_id,
			 produtos.codigo_de_barras,
			 produtos.descricao,

			 produtos.valor,
			 produtos.quantidade,
			 produtos.unidade_medida,
			 COALESCE(produtos.video_incorporado, '') as video_incorporado,
			 produtos.he_ativo,

			 produtos.he_promocao,
			 COALESCE(produtos.valor_promocao, 0) as valor_promocao,

			 COALESCE(produtos.he_combo, 0) as he_combo,
			 produtos.he_acompanhamento,
			 COALESCE(produtos.informacao_adicional, '') as informacao_adicional,

			 COALESCE(produtos.peso, '0') AS peso,
			 COALESCE(produtos.largura, 0) as largura,
			 COALESCE(produtos.altura, 0) as altura,
			 COALESCE(produtos.comprimento, 0) as comprimento,

			 COALESCE(imagens.id, 0) AS imagens_id,
			 COALESCE(imagens.descricao, '') AS imagens_descricao,
			 COALESCE(imagens.diretorio, '') AS imagens_diretorio
 
			 FROM public.produtos
			 LEFT JOIN imagens ON produtos.imagens_id = imagens.id
			WHERE parceiros_id = $1
		 ) 
		consulta WHERE produto_principal_id IS NOT null
	`
	argsProdutosCombos := []interface{}{
		p.Parceiros_Id,
	}

	err = tx.Select(&compounds, queryProdutosCombos, argsProdutosCombos...)
	if err != nil {
		return
	}

	for indexCompund, compound := range compounds {

		products = append(products, compound)

		queryCategory := `
		SELECT 
			id,
			parceiros_id,
			principal_produtos_id, 
			descricao,
			tipo_da_selecao,
			quantidade_max, 
			he_obrigatorio_selecao,
			he_adicional, 
			he_ativo
		FROM public.produtos_combos_categorias
		WHERE principal_produtos_id = $1
		AND parceiros_id = $2
		AND he_ativo = 1
	`
		argsCategory := []interface{}{
			compound.Produto_Principal_Id,
			p.Parceiros_Id,
		}

		err = tx.Select(&compound.Categorias, queryCategory, argsCategory...)
		if err != nil {
			return
		}

		for indexCategory, category := range compound.Categorias {

			products[indexCompund].Categorias = append(products[indexCompund].Categorias, category)

			queryCategory := `
				SELECT 
				public.produtos.id,
				public.produtos.parceiros_id,
				public.produtos.produtos_departamentos_id,
				public.produtos.produtos_marcas_id,
				COALESCE(produtos.produtos_categorias_id, 0) as produtos_categorias_id,
				public.produtos.codigo_de_barras,
				public.produtos.descricao,
				public.produtos.valor,
				public.produtos.quantidade,
				public.produtos.unidade_medida,
				COALESCE(produtos.video_incorporado, '') as video_incorporado,
				public.produtos.he_ativo,
				public.produtos.he_promocao,
				COALESCE(produtos.valor_promocao, 0) as valor_promocao,
				COALESCE(produtos.he_combo, 0) as he_combo,
				public.produtos.he_acompanhamento,
				COALESCE(produtos.informacao_adicional, '') as informacao_adicional,
				COALESCE(produtos.peso, '0') AS peso,
				COALESCE(produtos.largura, 0) as largura,
				COALESCE(produtos.altura, 0) as altura,
				COALESCE(produtos.comprimento, 0) as comprimento,
				
				public.produtos_combos.id AS produtos_combos_id,
				public.produtos_combos.produtos_id_principal,
				public.produtos_combos.produtos_combos_categorias_id,
			   public. produtos_combos.parceiros_id
			   FROM public.produtos
			   JOIN produtos_combos on produtos_combos.produtos_id_adicional = public.produtos.id
				WHERE produtos_combos.produtos_combos_categorias_id = $1
				AND produtos_combos.he_ativo = 1
				AND produtos_combos.parceiros_id = $2
			`
			argsCategory := []interface{}{
				category.ID,
				p.Parceiros_Id,
			}

			err = tx.Select(&category.Produtos_Categorias, queryCategory, argsCategory...)
			if err != nil {
				return
			}

			for _, product := range category.Produtos_Categorias {
				products[indexCompund].Categorias[indexCategory].Produtos_Categorias = append(products[indexCompund].Categorias[indexCategory].Produtos_Categorias, product)
			}

		}
	}

	return
}

func SearchProductCompositeByIdTx(p ParamsProductComposite, tx *sqlx.Tx) (products []ParamsProductComposite, err error) {

	var compounds []ParamsProductComposite

	queryProdutosCombos := `
	SELECT * FROM (
		SELECT 
		(
			SELECT DISTINCT produtos_id_principal FROM produtos_combos WHERE produtos.id = produtos_combos.produtos_id_principal) AS produto_principal_id,
			produtos.id,
			produtos.parceiros_id,
			produtos.produtos_departamentos_id,
			produtos.produtos_marcas_id,
			
			COALESCE(produtos.produtos_categorias_id, 0) as produtos_categorias_id,
			produtos.codigo_de_barras,
			produtos.descricao,
			 
			produtos.valor,
			produtos.quantidade,
			produtos.unidade_medida,
			COALESCE(produtos.video_incorporado, '') as video_incorporado,
			produtos.he_ativo,
			 
			produtos.he_promocao,
			COALESCE(produtos.valor_promocao, 0) as valor_promocao,
			COALESCE(produtos.he_combo, 0) as he_combo,
			produtos.he_acompanhamento,
			COALESCE(produtos.informacao_adicional, '') as informacao_adicional,
			 
			COALESCE(produtos.peso, '0') AS peso,
			COALESCE(produtos.largura, 0) as largura,
			COALESCE(produtos.altura, 0) as altura,
			COALESCE(produtos.comprimento, 0) as comprimento,

			COALESCE(imagens.id, 0) AS imagens_id,
			COALESCE(imagens.descricao, '') AS imagens_descricao,
			COALESCE(imagens.diretorio, '') AS imagens_diretorio

			FROM public.produtos
			LEFT JOIN imagens ON produtos.imagens_id = imagens.id

			WHERE produtos.id = $1
		 ) 
		consulta WHERE produto_principal_id IS NOT null
	`
	argsProdutosCombos := []interface{}{
		p.Produtos_Id,
	}

	err = tx.Select(&compounds, queryProdutosCombos, argsProdutosCombos...)
	if err != nil {
		return
	}

	for indexCompund, compound := range compounds {

		products = append(products, compound)

		queryCategory := `
		SELECT 
			id,
			parceiros_id,
			principal_produtos_id, 
			descricao,
			tipo_da_selecao,
			quantidade_max, 
			he_obrigatorio_selecao,
			he_adicional, 
			he_ativo
		FROM public.produtos_combos_categorias
		WHERE principal_produtos_id = $1
		AND he_ativo = 1
	`
		argsCategory := []interface{}{
			compound.Produto_Principal_Id,
		}

		err = tx.Select(&compound.Categorias, queryCategory, argsCategory...)
		if err != nil {
			return
		}

		for indexCategory, category := range compound.Categorias {

			products[indexCompund].Categorias = append(products[indexCompund].Categorias, category)

			queryCategory := `
				SELECT 
				public.produtos.id,
				public.produtos.parceiros_id,
				public.produtos.produtos_departamentos_id,
				public.produtos.produtos_marcas_id,
				COALESCE(produtos.produtos_categorias_id, 0) as produtos_categorias_id,
				public.produtos.codigo_de_barras,
				public.produtos.descricao,
				public.produtos.valor,
				public.produtos.quantidade,
				public.produtos.unidade_medida,
				COALESCE(produtos.video_incorporado, '') as video_incorporado,
				public.produtos.he_ativo,
				public.produtos.he_promocao,
				COALESCE(produtos.valor_promocao, 0) as valor_promocao,
				COALESCE(produtos.he_combo, 0) as he_combo,
				public.produtos.he_acompanhamento,
				COALESCE(produtos.informacao_adicional, '') as informacao_adicional,
				COALESCE(produtos.peso, '0') AS peso,
				COALESCE(produtos.largura, 0) as largura,
				COALESCE(produtos.altura, 0) as altura,
				COALESCE(produtos.comprimento, 0) as comprimento,
				
				public.produtos_combos.id AS produtos_combos_id,
				public.produtos_combos.produtos_id_principal,
				public.produtos_combos.produtos_combos_categorias_id,
			   public. produtos_combos.parceiros_id
			   FROM public.produtos
			   JOIN produtos_combos on produtos_combos.produtos_id_adicional = public.produtos.id
				WHERE produtos_combos.produtos_combos_categorias_id = $1
				AND produtos_combos.he_ativo = 1
			`
			argsCategory := []interface{}{
				category.ID,
			}

			err = tx.Select(&category.Produtos_Categorias, queryCategory, argsCategory...)
			if err != nil {
				return
			}

			for _, product := range category.Produtos_Categorias {
				products[indexCompund].Categorias[indexCategory].Produtos_Categorias = append(products[indexCompund].Categorias[indexCategory].Produtos_Categorias, product)
			}

		}
	}

	return
}

func AlterCompositeProductTx(p ParamsProductComposite, tx *sqlx.Tx) (err error) {

	queryUpProd := `UPDATE public.produtos SET  he_combo = $1 WHERE id = $2`
	argsUpProd := []interface{}{
		0,
		p.Produtos_Id_Principal,
	}
	_, err = tx.Exec(queryUpProd, argsUpProd...)
	if err != nil {
		return
	}

	queryUpCombo := `UPDATE public.produtos_combos_categorias SET  he_ativo = $1 WHERE principal_produtos_id = $2`
	argsUpCombo := []interface{}{
		0,
		p.Produtos_Id_Principal,
	}
	_, err = tx.Exec(queryUpCombo, argsUpCombo...)
	if err != nil {
		return
	}

	queryUpCategory := `UPDATE public.produtos_combos SET  he_ativo = $1 WHERE produtos_id_principal = $2`
	argsUpCategory := []interface{}{
		0,
		p.Produtos_Id_Principal,
	}
	_, err = tx.Exec(queryUpCategory, argsUpCategory...)
	if err != nil {
		return
	}

	return
}

func SearchProductsForUSerTx(p ParamsSearch, tx *sqlx.Tx) (products []Product, err error) {

	query := queryProductPublic

	pgEnd := 0
	offset := p.Pagina

	args := []interface{}{}

	descricao := strings.ToUpper(p.Descricao)

	if p.Descricao == "" {
		query += ` WHERE parceiros.id = $1
				AND produtos.he_acompanhamento = 0
				AND produtos.he_ativo = 1
			LIMIT 10 OFFSET $2 `
		if offset > 0 {
			pgEnd = offset * 10
		}
		args = []interface{}{
			p.Parceiros_Id,
			pgEnd,
		}
	} else {
		query += `
			WHERE parceiros.id = $1
				AND produtos.he_acompanhamento = 0
				AND produtos.he_ativo = 1
				AND upper(produtos.descricao) LIKE '%' || $2 || '%'
		`
		args = []interface{}{
			p.Parceiros_Id,
			descricao,
		}
	}

	err = tx.Select(&products, query, args...)
	if err != nil {
		return
	}

	return
}

func SearchCountProductsTx(p ParamsSearch, tx *sqlx.Tx) (countProducts int, err error) {
	var products []Product

	query := `SELECT COUNT(id) as count_produtos
		FROM public.produtos
		WHERE produtos.parceiros_id = $1
		AND produtos.he_ativo = 1`

	args := []interface{}{
		p.Parceiros_Id,
	}

	err = tx.Select(&products, query, args...)
	if err != nil {
		return
	}

	if len(products) > 0 {
		countProducts = products[0].Count_Produtos
	} else {
		countProducts = 0
	}

	return
}

func SearchProductsForUSerFromFiltersTx(p ParamsSearch, tx *sqlx.Tx) (products []Product, err error) {

	query := queryProductPublic + ` WHERE parceiros.id = $1
			AND produtos.he_acompanhamento = 0
			AND produtos.he_ativo = 1 `

	args := []interface{}{
		p.Parceiros_Id,
	}

	if p.Marcas_Ids != "" {
		query += ` AND produtos_marcas.id IN (` + p.Marcas_Ids + `) `
	}

	if p.Departamentos_Ids != "" {
		query += ` AND produtos_departamentos.id IN (` + p.Departamentos_Ids + `) `
	}

	if p.Categorias_Ids != "" {
		query += ` AND produtos_categorias.id IN (` + p.Categorias_Ids + `) `
	}

	err = tx.Select(&products, query, args...)
	if err != nil {
		return
	}

	return
}

func CreateFavoriteTx(f Favorite, tx *sqlx.Tx) (id int64, err error) {

	query := `
		INSERT INTO public.produtos_favoritos(
			produtos_id,
			pessoas_usuarios_id, 
			he_ativo)
		VALUES ($1, $2, $3)
		RETURNING id;
		`

	args := []interface{}{
		f.Produtos_Id,
		f.Pessoas_Usuarios_id,
		1,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func AlterFavoriteTx(f Favorite, tx *sqlx.Tx) (err error) {

	query := `UPDATE public.produtos_favoritos
	SET 
	 	 produtos_id = $1,
		 pessoas_usuarios_id = $2,
		 he_ativo = $3
	WHERE id = $4;`

	args := []interface{}{
		f.Produtos_Id,
		f.Pessoas_Usuarios_id,
		f.He_Ativo,
		f.ID,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func SearchUserFromIdPersonTx(u Favorite, tx *sqlx.Tx) (user []Favorite, err error) {

	query := `SELECT 
		id,
		pessoas_id
		FROM public.pessoas_usuarios 
		WHERE pessoas_id = $1;`

	args := []interface{}{
		u.Pessoas_Id,
	}

	err = tx.Select(&user, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchFavoritesTx(f Favorite, tx *sqlx.Tx) (favorites []Product, err error) {

	query := `SELECT
			produtos_favoritos.id as produtos_favoritos_id,
			parceiros.id as parceiros_id,
			produtos.id,
			produtos.produtos_departamentos_id,
			produtos.produtos_marcas_id,
			COALESCE(produtos.produtos_categorias_id, 0) as produtos_categorias_id,
			produtos.codigo_de_barras,
			produtos.descricao,
			produtos.valor,
			produtos.quantidade,
			produtos.unidade_medida,
			COALESCE(produtos.video_incorporado, '') as video_incorporado,
			produtos.he_ativo,
			produtos.he_promocao,
			COALESCE(produtos.valor_promocao, 0) as valor_promocao,
			COALESCE(produtos.he_combo, 0) as he_combo,
			produtos.he_acompanhamento,
			COALESCE(produtos.informacao_adicional, '') as informacao_adicional,
			COALESCE(produtos.peso, 0) as peso,
			COALESCE(produtos.largura, 0) as largura,
			COALESCE(produtos.altura, 0) as altura,
			COALESCE(produtos.comprimento, 0) as comprimento,
			produtos_departamentos.descricao as produtos_departamentos_descricao,
			produtos_categorias.descricao as produtos_categorias_descricao,
			produtos_marcas.descricao as produtos_marcas_descricao,
			COALESCE(imagens.id, 0) AS imagens_id,
			COALESCE(imagens.descricao, '') AS imagens_descricao,
			COALESCE(imagens.diretorio, '') AS imagens_diretorio
		FROM public.produtos_favoritos
		JOIN produtos ON  produtos.id = produtos_favoritos.produtos_id
		JOIN produtos_departamentos ON  produtos_departamentos.id = produtos.produtos_departamentos_id
		LEFT JOIN produtos_categorias ON produtos_categorias.id = produtos.produtos_categorias_id
		LEFT JOIN produtos_marcas ON  produtos_marcas.id = produtos.produtos_marcas_id
		LEFT JOIN imagens ON produtos.imagens_id = imagens.id
		JOIN parceiros ON produtos.parceiros_id = parceiros.id
		WHERE produtos_favoritos.pessoas_usuarios_id = $1 and produtos_favoritos.he_ativo = 1;`

	args := []interface{}{
		f.Pessoas_Usuarios_id,
	}

	err = tx.Select(&favorites, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchProductsForUSerFromIdTx(p Product, tx *sqlx.Tx) (products []Product, err error) {
	query := queryProductPublic

	args := []interface{}{}

	query += ` WHERE parceiros.id = $1
				AND produtos.id = $2`

	args = []interface{}{
		p.Parceiros_Id,
		p.ID,
	}

	err = tx.Select(&products, query, args...)
	if err != nil {
		return
	}

	return
}

func SearchImagensProductsForCatalogTx(p Product, tx *sqlx.Tx) (imgs []ImgCatalog, err error) {

	query := `SELECT
    produtos.id AS  produtos_id,
    COALESCE(produtos.video_incorporado, '') as video_incorporado,

	produtos_imagens.id as produtos_imagens_id,
    
    imagens.id AS imagens_id,
    imagens.descricao AS imagens_descricao,
    imagens.diretorio AS imagens_diretorio

        FROM
            produtos
		JOIN parceiros ON produtos.parceiros_id = parceiros.id
		JOIN produtos_imagens ON produtos_imagens.produtos_id = produtos.id
		JOIN imagens ON imagens.id = produtos_imagens.imagens_id

        WHERE parceiros.id = $1 AND produtos.id = $2 AND produtos_imagens.he_ativo = 1`

	args := []interface{}{
		p.Parceiros_Id,
		p.ID,
	}

	err = tx.Select(&imgs, query, args...)
	if err != nil {
		return
	}

	return
}
