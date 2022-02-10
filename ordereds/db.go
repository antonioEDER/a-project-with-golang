package ordereds

import (
	"errors"
	"fmt"
	"time"

	"github.com/api-qop-v2/partners"
	"github.com/jmoiron/sqlx"
)

const queryPublic = `
SELECT
    pedidos.id,
    COALESCE(pedidos.parceiros_id, 0) as parceiros_id, 
    COALESCE(pedidos.pessoas_usuarios_id, 0) as pessoas_usuarios_id,
    COALESCE(pedidos.pedidos_categorias_id, 0) as pedidos_categorias_id,
    COALESCE(pedidos.pedidos_status_id, 0) as pedidos_status_id, 
    COALESCE(pedidos.enderecos_id, 0) as enderecos_id, 
    COALESCE(pedidos.valor_total,  0) as valor_total, 
    COALESCE(pedidos.forma_pagamento,  '') as forma_pagamento, 
    COALESCE(pedidos.troco, '') as troco,
    COALESCE(pedidos.servico_pagamento, '') as servico_pagamento,
    COALESCE(pedidos.taxa_entrega, 0) as taxa_entrega, 
	COALESCE(pedidos.tipo_veiculo, '') as tipo_veiculo, 
	COALESCE(pedidos.placa_veiculo, '') as placa_veiculo, 
	COALESCE(pedidos.cor_veiculo, '') as cor_veiculo,
	COALESCE(pedidos.entregador, '') as entregador,
    COALESCE(pedidos.visualizado, 0) as visualizado,
    COALESCE(pedidos.he_ativo, 0) as he_ativo,
	COALESCE(pedidos.prazo_entrega, '') as prazo_entrega,

    (SELECT pedidos.created_at at time zone 'UTC' at time zone $1) as created_at,

    COALESCE(pedidos_itens.id, 0) as pedidos_itens_id,
    COALESCE(pedidos_itens.pedidos_id , 0) as pedidos_id,
    COALESCE(pedidos_itens.valor , 0) as valor,
	COALESCE(pedidos_itens.vinculo_para_produto_composto, '') as vinculo_para_produto_composto,
    
    COALESCE(produtos.id , 0) as produtos_id,
	COALESCE(produtos.produtos_departamentos_id, 0) as produtos_departamentos_id,
	COALESCE(produtos.produtos_marcas_id, 0) as produtos_marcas_id,
	COALESCE(produtos.produtos_categorias_id, 0) as produtos_categorias_id,
	COALESCE(produtos.codigo_de_barras, '0') as codigo_de_barras,
	COALESCE(produtos.descricao, '') as descricao,
	COALESCE(produtos.quantidade, 0) as quantidade,
	COALESCE(produtos.unidade_medida, '') as unidade_medida,
	COALESCE(produtos.video_incorporado, '') as video_incorporado,
	COALESCE(produtos.he_promocao, 0) as he_promocao,

	COALESCE(produtos.valor_promocao, 0) as valor_promocao,
	COALESCE(produtos.he_combo, 0) as he_combo,

	COALESCE(produtos.he_acompanhamento, 0) as he_acompanhamento,
	COALESCE(produtos.informacao_adicional, '') as informacao_adicional,
	COALESCE(produtos.peso, 0) as peso,
	COALESCE(produtos.largura, 0) as largura,
	COALESCE(produtos.altura, 0) as altura,
	COALESCE(produtos.comprimento, 0) as comprimento,

    COALESCE(pedidos_status.descricao, '') as pedidos_status_descricao,
	COALESCE(pedidos_status.id, 0) as pedidos_status_id,
    COALESCE(pedidos_status.cor, '') as pedidos_status_cor,

    COALESCE(imagens.id, 0) as imagens_id,
	COALESCE(imagens.descricao, '') as imagens_descricao,
	COALESCE(imagens.diretorio, '') as  imagens_diretorio,

	COALESCE(pedidos_categorias.descricao, '') as pedidos_categorias_descricao,
    COALESCE(pedidos_categorias.id, 0) as pedidos_categorias_id

    FROM
    pedidos_itens
    JOIN pedidos ON pedidos.id = pedidos_itens.pedidos_id
    JOIN produtos ON produtos.id = pedidos_itens.produtos_id
    JOIN pedidos_status ON pedidos_status.id = pedidos.pedidos_status_id
    LEFT JOIN imagens ON produtos.imagens_id = imagens.id
    JOIN pedidos_categorias pedidos_categorias ON pedidos_categorias.id = pedidos.pedidos_categorias_id`

func findByFilterDBTx(filter Filter, tx *sqlx.Tx) (user string, err error) {

	if filter.ID > 0 && tx != nil {
		user = "Seens we have a user :)"
	} else {
		err = errors.New("There's no user")
	}

	return
}

func CreateOrderTx(o Order, tx *sqlx.Tx) (id int64, err error) {

	now := time.Now()
	loc, _ := time.LoadLocation("UTC")
	o.CREATED_AT = now.In(loc)

	query := `
	INSERT INTO public.pedidos(
		parceiros_id,
		pessoas_usuarios_id,
		pedidos_categorias_id,
		pedidos_status_id,
		enderecos_id,
		valor_total,
		forma_pagamento,
		troco,
		servico_pagamento,
		taxa_entrega,
		tipo_veiculo,
		placa_veiculo,
		cor_veiculo,
		entregador,
		he_ativo,
		prazo_entrega,
		created_at)
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
	    $17)
		RETURNING id;`

	taxaEntrega := float64(0)
	if o.Pedidos_Categorias_Id == 2 {
		taxaEntrega = o.Taxa_Entrega
	}

	args := []interface{}{
		o.Parceiros_Id,
		o.Pessoas_Usuarios_Id,
		o.Pedidos_Categorias_Id,
		o.Pedidos_Status_Id,
		o.Enderecos_Id,
		o.Valor_Total,
		o.Forma_Pagamento,
		o.Troco,
		o.Servico_Pagamento,
		taxaEntrega,
		o.Tipo_Veiculo,
		o.Placa_Veiculo,
		o.Cor_Veiculo,
		o.Entregador,
		1,
		o.Prazo_Entrega,
		o.CREATED_AT,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func CreateListProductsFromOrderTx(idOrder int64, i OrderItens, tx *sqlx.Tx) (id int64, err error) {

	query := `INSERT INTO public.pedidos_itens(
		pedidos_id,
		produtos_id,
		valor,
		vinculo_para_produto_composto,
		he_ativo)
	   VALUES (
		$1,
		$2,
		$3,
		$4,
		$5)
		RETURNING id;`

	args := []interface{}{
		i.Pedidos_Id,
		i.Produtos_Id,
		i.Valor,
		i.Vinculo_Para_Produto_Composto,
		1,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func CreateListProductsCompoundFromOrderTx(idOrder int64, o Order, p Produtos, pc ProdutoCompostoAdd, tx *sqlx.Tx) (id int64, err error) {

	query := `INSERT INTO public.pedidos_produtos_combos(
		parceiros_id,
		produtos_id,
		principal_produtos_id,
		pedidos_id,
		produtos_categorias_id,
		vinculo_para_produto_composto,
		identificador_adicional,
		he_ativo,
		pessoas_usuarios_id
	)
	   VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9)
		RETURNING id;`

	args := []interface{}{
		o.Parceiros_Id,
		pc.Produtos_Id,
		p.ID,
		idOrder,
		pc.Produtos_Categorias_Id,
		p.Vinculo_Para_Produto_Composto,
		pc.Identificador_Adicional,
		1,
		o.Pessoas_Usuarios_Id,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func SearchOrderSpecificActiveTx(o Order, tx *sqlx.Tx) (orders []Order, err error) {

	query := queryPublic + ` WHERE pedidos.id = $2 and pedidos.pessoas_usuarios_id = $3 AND pedidos.he_ativo = $4; `

	args := []interface{}{
		o.Time_Zone,
		o.Pedidos_Id,
		o.Pessoas_Usuarios_Id,
		1,
	}

	err = tx.Select(&orders, query, args...)
	if err != nil {
		return
	}

	return
}

func SearchOrderToResendProductDigitalTx(o Order, tx *sqlx.Tx) (orders []Order, err error) {

	query := queryPublic + ` WHERE pedidos.id = $2 AND pedidos.he_ativo = $3; `

	args := []interface{}{
		"America/Porto_Velho",
		o.Pedidos_Id,
		1,
	}

	err = tx.Select(&orders, query, args...)
	if err != nil {
		return
	}

	return
}

func SearchProductCompositeByIdTx(p ParamsProductComposite, tx *sqlx.Tx) (products []ParamsProductComposite, err error) {

	queryProdutosCombos := `SELECT
			COALESCE(imagens.id, 0) AS imagens_id,
			COALESCE(imagens.descricao, '') AS imagens_descricao,
			COALESCE(imagens.diretorio, '') AS imagens_diretorio,

			COALESCE(pedidos_produtos_combos.vinculo_para_produto_composto, '') as vinculo_para_produto_composto,
			COALESCE(pedidos_produtos_combos.produtos_categorias_id, 0) as produtos_categorias_id,
			COALESCE(produtos.id, 0) as produtos_id,
			COALESCE(produtos.produtos_departamentos_id , 0) as produtos_departamentos_id,
			COALESCE(produtos.produtos_marcas_id , 0) as produtos_marcas_id,
			COALESCE(produtos.codigo_de_barras , '') as codigo_de_barras,
			COALESCE(produtos.descricao , '') as descricao,
			COALESCE(produtos.quantidade , 0) as quantidade,
			COALESCE(produtos.unidade_medida , '') as unidade_medida,
			COALESCE(produtos.video_incorporado , '') as video_incorporado,
			COALESCE(produtos.he_ativo  , 0) as he_ativo,
			COALESCE(produtos.he_promocao  , 0) as he_promocao,
			COALESCE(produtos.valor_promocao, 0) as valor_promocao,
			COALESCE(produtos.valor , 0) as valor,

			COALESCE(produtos.he_combo, 0) as he_combo,
			
			COALESCE(produtos.he_acompanhamento, 0) as he_acompanhamento,

			COALESCE(produtos.informacao_adicional, '') as informacao_adicional,

			COALESCE(produtos.peso, 0) as peso,
			COALESCE(produtos.largura, 0) as largura,
			COALESCE(produtos.altura, 0) as altura,
			COALESCE(produtos.comprimento, 0) as comprimento
		FROM
			pedidos_produtos_combos
				JOIN
			produtos ON produtos.id = pedidos_produtos_combos.principal_produtos_id
			LEFT JOIN imagens ON produtos.imagens_id = imagens.id
		WHERE
			pedidos_produtos_combos.pedidos_id = $1
		GROUP BY 
		imagens.id,
		pedidos_produtos_combos.vinculo_para_produto_composto,
		pedidos_produtos_combos.produtos_categorias_id,
		produtos.id
		`
	argsProdutosCombos := []interface{}{
		p.Pedidos_Id,
	}

	var compounds []ParamsProductComposite

	err = tx.Select(&compounds, queryProdutosCombos, argsProdutosCombos...)
	if err != nil {
		return
	}

	for indexCompund, compound := range compounds {

		products = append(products, compound)

		queryCategory := `
			SELECT
				COALESCE(imagens.id, 0) AS imagens_id,
				COALESCE(imagens.descricao, '') AS imagens_descricao,
				COALESCE(imagens.diretorio, '') AS imagens_diretorio,
				COALESCE(pedidos_produtos_combos.vinculo_para_produto_composto, '') as vinculo_para_produto_composto,
				COALESCE(pedidos_produtos_combos.produtos_categorias_id, 0) as produtos_categorias_id,
				COALESCE(produtos_combos_categorias.descricao, '') as produtos_combos_categorias_descricao

				FROM
					pedidos_produtos_combos
						JOIN
					produtos_combos_categorias ON produtos_combos_categorias.id = pedidos_produtos_combos.produtos_categorias_id
						JOIN
					produtos ON produtos.id = pedidos_produtos_combos.principal_produtos_id
						LEFT JOIN
				imagens ON produtos.imagens_id = imagens.id
				WHERE
					pedidos_produtos_combos.pedidos_id = $1
				AND pedidos_produtos_combos.vinculo_para_produto_composto = $2
				AND pedidos_produtos_combos.he_ativo = 1
				GROUP BY 
					pedidos_produtos_combos.produtos_categorias_id,
					imagens.id,
					pedidos_produtos_combos.vinculo_para_produto_composto,
					pedidos_produtos_combos.produtos_categorias_id,
					produtos_combos_categorias.descricao
				`
		argsCategory := []interface{}{
			p.Pedidos_Id,
			compound.Vinculo_Para_Produto_Composto,
		}

		err = tx.Select(&compound.Categorias, queryCategory, argsCategory...)
		if err != nil {
			return
		}

		for indexCategory, category := range compound.Categorias {

			products[indexCompund].Categorias = append(products[indexCompund].Categorias, category)

			queryCategory := `SELECT COALESCE(identificador_adicional, '') as identificador_adicional
			FROM 
				pedidos_produtos_combos
			WHERE pedidos_id = $1 AND produtos_categorias_id = $2
			AND pedidos_produtos_combos.he_ativo = 1
			GROUP BY identificador_adicional`

			var additionals []AdditionalIdentifier

			argsCategory := []interface{}{
				p.Pedidos_Id,
				category.Produtos_Categorias_Id,
			}

			err = tx.Select(&additionals, queryCategory, argsCategory...)
			if err != nil {
				return
			}

			for _, additional := range additionals {
				category.Produtos_Categorias = []Product{}

				queryCategory := `
					SELECT
						COALESCE(imagens.id, 0) AS imagens_id,
						COALESCE(imagens.descricao, '') AS imagens_descricao,
						COALESCE(imagens.diretorio, '') AS imagens_diretorio,
						COUNT(produtos.id) qtd_produto,

						COALESCE(produtos.descricao, '') as descricao,
						COALESCE(produtos.id, 0) as produtos_id,
						COALESCE(produtos.unidade_medida, '') as unidade_medida,
						COALESCE(produtos.he_promocao, 0) as he_promocao,
						COALESCE(produtos.valor_promocao, 0) as valor_promocao,

						COALESCE(produtos.valor, 0) as valor,
						COALESCE(produtos_combos_categorias.descricao, '') as produtos_combos_categorias_descricao,
						COALESCE(produtos_combos_categorias.id, 0) as produtos_combos_categorias_id,
						COALESCE(produtos_combos_categorias.principal_produtos_id, 0) as principal_produtos_id,
						COALESCE(produtos_combos_categorias.he_adicional , 0) as he_adicional,
						COALESCE(pedidos_produtos_combos.vinculo_para_produto_composto, '') as vinculo_para_produto_composto,
						COALESCE(pedidos_produtos_combos.identificador_adicional, '') as identificador_adicional
					FROM
						pedidos_produtos_combos
							JOIN
						produtos_combos ON produtos_combos.produtos_id_adicional = pedidos_produtos_combos.produtos_id
							JOIN
						produtos_combos_categorias ON produtos_combos_categorias.id = produtos_combos.produtos_combos_categorias_id
							JOIN
						produtos ON produtos.id = pedidos_produtos_combos.produtos_id
							LEFT JOIN imagens ON produtos.imagens_id = imagens.id
					WHERE 
						pedidos_produtos_combos.pedidos_id = $1
					AND produtos_combos_categorias.id = $2
					AND pedidos_produtos_combos.vinculo_para_produto_composto = $3
					AND pedidos_produtos_combos.identificador_adicional = $4
					GROUP BY 
					produtos.id,
					produtos.descricao,
					imagens.id,
					produtos_combos_categorias.id,
					pedidos_produtos_combos.vinculo_para_produto_composto,
					pedidos_produtos_combos.identificador_adicional
				`
				argsCategory := []interface{}{
					p.Pedidos_Id,
					category.Produtos_Categorias_Id,
					category.Vinculo_Para_Produto_Composto,
					additional.Identificador_Adicional,
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
	}

	return
}

func SearchOrderForPartnerTx(o Order, tx *sqlx.Tx) (orders []Order, err error) {

	query := `SELECT
			pedidos.id,
			COALESCE(pedidos.parceiros_id, 0) as parceiros_id, 
			COALESCE(pedidos.pessoas_usuarios_id , 0) as pessoas_usuarios_id, 
			COALESCE(pedidos.pedidos_categorias_id , 0) as pedidos_categorias_id, 
			COALESCE(pedidos.pedidos_status_id , 0) as pedidos_status_id, 
			COALESCE(pedidos.enderecos_id  , 0) as enderecos_id, 
			COALESCE(pedidos.valor_total , 0) as valor_total,
			COALESCE(pedidos.forma_pagamento , '') as forma_pagamento,

			COALESCE(pedidos.troco, '') as troco,
			COALESCE(pedidos.servico_pagamento, '') as servico_pagamento,
			COALESCE(pedidos.taxa_entrega, 0) as taxa_entrega,
			COALESCE(pedidos.tipo_veiculo, '') as tipo_veiculo, 
			COALESCE(pedidos.placa_veiculo, '') as placa_veiculo, 
			COALESCE(pedidos.cor_veiculo, '') as cor_veiculo,
			COALESCE(pedidos.entregador, '') as entregador,
			COALESCE(pedidos.visualizado, 0) as visualizado,
			COALESCE(pedidos.he_ativo,0) as he_ativo,
			COALESCE(pedidos.prazo_entrega, '') as prazo_entrega,

			(SELECT pedidos.created_at at time zone 'UTC' at time zone $1) as created_at,

			COALESCE(pedidos_itens.id,0) as pedidos_itens_id,
			COALESCE(pedidos_itens.pedidos_id, 0) as pedidos_id, 
			COALESCE(pedidos_itens.valor, 0) as valor,

			COALESCE(pedidos_itens.vinculo_para_produto_composto, '') as vinculo_para_produto_composto, 
			
			produtos.id as produtos_id,

			COALESCE(produtos.produtos_departamentos_id, 0) as produtos_departamentos_id,
			COALESCE(produtos.produtos_marcas_id , 0) as produtos_marcas_id,
			COALESCE(produtos.produtos_categorias_id , 0) as produtos_categorias_id,
			COALESCE(produtos.codigo_de_barras , '') as codigo_de_barras,
			COALESCE(produtos.descricao , '') as descricao,
			COALESCE(produtos.quantidade  , 0) as quantidade,
			COALESCE(produtos.unidade_medida  , '') as unidade_medida,
			COALESCE(produtos.video_incorporado , '') as video_incorporado,
			COALESCE(produtos.he_ativo  , 0) as he_ativo,
			COALESCE(produtos.he_promocao  , 0) as he_promocao,

			COALESCE(produtos.valor_promocao, 0) as valor_promocao,
			COALESCE(produtos.he_combo, 0) as he_combo,
			COALESCE(produtos.he_acompanhamento,0) as he_acompanhamento,
			COALESCE(produtos.informacao_adicional, '') as informacao_adicional,

			COALESCE(produtos.peso, 0) as peso,
			COALESCE(produtos.largura, 0) as largura,
			COALESCE(produtos.altura, 0) as altura,
			COALESCE(produtos.comprimento, 0) as comprimento,

			COALESCE(pedidos_status.descricao, '') as pedidos_status_descricao,
			COALESCE(pedidos_status.id, 0) as pedidos_status_id,
			COALESCE(pedidos_status.cor, '') as pedidos_status_cor,

			COALESCE(pedidos_categorias.descricao, '') as pedidos_categorias_descricao,
			COALESCE(pedidos_categorias.id, 0) as pedidos_categorias_id,
			COALESCE(pedidos_categorias.pedidos_status_ids::varchar(255), '') as pedidos_status_ids,
			
			COALESCE(imagens.id, 0) as imagens_id,
			COALESCE(imagens.descricao, '') as imagens_descricao,
			COALESCE(imagens.diretorio, '') as  imagens_diretorio,

			COALESCE(pessoas.id, 0) as pessoas_id,
			COALESCE(public.pessoas_usuarios.id, 0) as pessoas_usuarios_id,
			COALESCE(public.pessoas_usuarios.email , '') as email,

			COALESCE(public.pessoas.nome, '') as nome,
			
			COALESCE(public.pessoas_fisicas.cpf, '') as cpf,
			COALESCE(public.pessoas_fisicas.data_nascimento, now()) as data_nascimento,
			
			COALESCE(public.pessoas_contatos.contato, '') as celular,

			COALESCE(public.enderecos.cep , '') as cep, 
			COALESCE(public.enderecos.logradouro, '') as logradouro, 
			COALESCE(public.enderecos.bairro, '') as bairro, 

			COALESCE(public.enderecos.complemento, '') as complemento,
			COALESCE(public.enderecos.numero , '') as numero,
			COALESCE(public.cidades.nome, '') as cidade, 
			COALESCE(public.estados.uf, '') as uf
		FROM
			pedidos_itens
		JOIN pedidos ON pedidos.id = pedidos_itens.pedidos_id
		JOIN produtos ON produtos.id = pedidos_itens.produtos_id
		LEFT JOIN imagens ON produtos.imagens_id = imagens.id
		JOIN public.pessoas_usuarios ON  public.pessoas_usuarios.id = public.pedidos.pessoas_usuarios_id
		JOIN public.pessoas on public.pessoas.id = public.pessoas_usuarios.pessoas_id
		JOIN  public.pessoas_fisicas ON  public.pessoas_fisicas.pessoas_id = public.pessoas.id
		LEFT JOIN   public.pessoas_contatos ON  public.pessoas_contatos.pessoas_id = public.pessoas.id
		LEFT JOIN   public.enderecos ON  public.enderecos.pessoas_id = public.pessoas.id
		LEFT JOIN   public.cidades  on public.cidades.id = public.enderecos.cidades_id
		LEFT JOIN   public.estados  on public.estados.id = public.cidades.estados_id

		JOIN pedidos_status ON pedidos_status.id = pedidos.pedidos_status_id
		JOIN pedidos_categorias ON pedidos_categorias.id = pedidos.pedidos_categorias_id
	    WHERE pedidos.id = $2 AND enderecos.he_principal = 1`

	args := []interface{}{
		o.Time_Zone,
		o.Pedidos_Id,
	}

	err = tx.Select(&orders, query, args...)
	if err != nil {
		return
	}

	return
}

func SearchOrderForUserTx(o Order, tx *sqlx.Tx) (orders []Order, err error) {

	query := queryPublic + ` WHERE pedidos.pessoas_usuarios_id = $2 AND pedidos.he_ativo = $3 ORDER BY id desc; `

	args := []interface{}{
		o.Time_Zone,
		o.Pessoas_Usuarios_Id,
		1,
	}

	err = tx.Select(&orders, query, args...)
	if err != nil {
		return
	}

	return
}

func CancelOrderTx(order Order, tx *sqlx.Tx) (err error) {
	query := `
		UPDATE pedidos set he_ativo = 0, pedidos_status_id = 13, updated_at = NOW() 
			WHERE id = $1
		`

	args := []interface{}{
		order.Pedidos_Id,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func SearchCategoriesOrderTx(tx *sqlx.Tx) (orders []OrderStatus, err error) {

	query := `SELECT 
		id, 
		COALESCE(descricao, '') AS descricao,
		COALESCE(cor, '') AS cor,
		COALESCE(icone, '') AS icone,
		COALESCE(he_ativo, 0) AS he_ativo
		
	FROM public.pedidos_status WHERE he_ativo = 1;`

	args := []interface{}{}

	err = tx.Select(&orders, query, args...)
	if err != nil {
		return
	}

	return
}

func SearchStatusOrderTx(tx *sqlx.Tx) (orders []OrderCategory, err error) {

	query := `SELECT id, descricao, pedidos_status_ids::varchar(255), he_ativo
	FROM public.pedidos_categorias
	WHERE  he_ativo = 1`

	args := []interface{}{}

	err = tx.Select(&orders, query, args...)
	if err != nil {
		return
	}

	return
}

func SearchOrdersForFiltersTx(filters string, args []interface{}, tx *sqlx.Tx) (orders []Order, err error) {

	query := `SELECT
			pedidos.id,
			pedidos.parceiros_id, 
			pedidos.pessoas_usuarios_id, 
			pedidos.pedidos_categorias_id, 
			pedidos.pedidos_status_id, 
			pedidos.enderecos_id, 
			pedidos.valor_total, 
			pedidos.forma_pagamento, 
			
			COALESCE(pedidos.troco, '') as troco, 
			COALESCE(pedidos.servico_pagamento, '') as servico_pagamento,
			pedidos.taxa_entrega,
			COALESCE(pedidos.tipo_veiculo, '') as tipo_veiculo, 
			COALESCE(pedidos.placa_veiculo, '') as placa_veiculo, 
			COALESCE(pedidos.cor_veiculo, '') as cor_veiculo, 
			COALESCE(pedidos.entregador, '') as entregador,
			pedidos.visualizado, 
			pedidos.he_ativo,
			COALESCE(pedidos.prazo_entrega, '') as prazo_entrega,

			(SELECT pedidos.created_at at time zone 'UTC' at time zone $1) as created_at,

			pedidos_status.descricao as pedidos_status_descricao,
			pedidos_status.id as pedidos_status_id,
			COALESCE(pedidos_status.cor, '') as pedidos_status_cor,


			pedidos_categorias.descricao as pedidos_categorias_descricao,
			pedidos_categorias.id as pedidos_categorias_id,
			pedidos_categorias.pedidos_status_ids::varchar(255) as pedidos_status_ids,

			pessoas.id as pessoas_id,
			public.pessoas_usuarios.id as pessoas_usuarios_id,
			public.pessoas_usuarios.email,

			public.pessoas.id as pessoas_id,
			public.pessoas.nome,
			
			COALESCE(public.pessoas_fisicas.cpf, '') as cpf,
			COALESCE(public.pessoas_fisicas.data_nascimento, now()) as data_nascimento,
			
			public.pessoas_contatos.contato as celular,

			public.enderecos.cep, 
			public.enderecos.logradouro, 
			public.enderecos.bairro, 
			COALESCE(public.enderecos.complemento, '') as complemento,
			public.enderecos.numero,
			public.cidades.nome as cidade, 
			public.estados.uf
		FROM
			pedidos
		JOIN public.pessoas_usuarios ON  public.pessoas_usuarios.id = public.pedidos.pessoas_usuarios_id
		JOIN public.pessoas on public.pessoas.id = public.pessoas_usuarios.pessoas_id
		JOIN  public.pessoas_fisicas ON  public.pessoas_fisicas.pessoas_id = public.pessoas.id
		LEFT JOIN   public.pessoas_contatos ON  public.pessoas_contatos.pessoas_id = public.pessoas.id
		LEFT JOIN   public.enderecos ON  public.enderecos.pessoas_id = public.pessoas.id
		LEFT JOIN   public.cidades  on public.cidades.id = public.enderecos.cidades_id
		LEFT JOIN   public.estados  on public.estados.id = public.cidades.estados_id

		JOIN pedidos_status ON pedidos_status.id = pedidos.pedidos_status_id
		JOIN pedidos_categorias ON pedidos_categorias.id = pedidos.pedidos_categorias_id ` + filters

	err = tx.Select(&orders, query, args...)
	if err != nil {
		return
	}

	return
}

func AlterOrdersAcceptedTx(order Order, tx *sqlx.Tx) (err error) {
	query := `
		UPDATE pedidos 
			SET visualizado = 1, 
			he_ativo = 1 
		WHERE id = $1 AND  parceiros_id = $2
		`

	args := []interface{}{
		order.Pedidos_Id,
		order.Parceiros_Id,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func AlterOrderStatusTx(order Order, pt partners.Partners, tx *sqlx.Tx) (err error) {
	query := `
	UPDATE pedidos SET pedidos_status_id = $1, he_ativo = 1 WHERE id = $2 AND  parceiros_id = $3
		`

	args := []interface{}{
		order.Pedidos_Status_Id,
		order.Pedidos_Id,
		pt.Id,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func SearchOrdersForIConditionsTx(o Order, tx *sqlx.Tx) (orders []Order, err error) {

	query := `SELECT
			pedidos.id,
			pedidos.parceiros_id, 
			pedidos.pessoas_usuarios_id, 
			pedidos.pedidos_categorias_id, 
			pedidos.pedidos_status_id, 
			pedidos.enderecos_id, 
			pedidos.valor_total, 
			pedidos.forma_pagamento, 
			COALESCE(pedidos.troco, '') as troco,
			COALESCE(pedidos.servico_pagamento, '') as servico_pagamento,
			pedidos.taxa_entrega, 
			COALESCE(pedidos.tipo_veiculo, '') as tipo_veiculo, 
			COALESCE(pedidos.placa_veiculo, '') as placa_veiculo, 
			COALESCE(pedidos.cor_veiculo, '') as cor_veiculo,
			COALESCE(pedidos.entregador, '') as entregador,
			pedidos.visualizado, 
			pedidos.he_ativo,
			COALESCE(pedidos.prazo_entrega, '') as prazo_entrega,

			pedidos_status.descricao as pedidos_status_descricao,
			pedidos_status.id as pedidos_status_id,
			COALESCE(pedidos_status.cor, '') as pedidos_status_cor,

			pedidos_categorias.descricao as pedidos_categorias_descricao,
			pedidos_categorias.id as pedidos_categorias_id,
			pedidos_categorias.pedidos_status_ids::varchar(255) as pedidos_status_ids,

			pessoas.id as pessoas_id,
			public.pessoas_usuarios.id as pessoas_usuarios_id,
			public.pessoas_usuarios.email,

			public.pessoas.id as pessoas_id,
			public.pessoas.nome,
			
			COALESCE(public.pessoas_fisicas.cpf, '') as cpf,
			COALESCE(public.pessoas_fisicas.data_nascimento, now()) as data_nascimento,
			
			public.pessoas_contatos.contato as celular,

			public.enderecos.cep, 
			public.enderecos.logradouro, 
			public.enderecos.bairro, 
			COALESCE(public.enderecos.complemento, '') as complemento,
			public.enderecos.numero,
			public.cidades.nome as cidade, 
			public.estados.uf
		FROM
			pedidos
		JOIN public.pessoas_usuarios ON  public.pessoas_usuarios.id = public.pedidos.pessoas_usuarios_id
		JOIN public.pessoas on public.pessoas.id = public.pessoas_usuarios.pessoas_id
		JOIN  public.pessoas_fisicas ON  public.pessoas_fisicas.pessoas_id = public.pessoas.id
		LEFT JOIN   public.pessoas_contatos ON  public.pessoas_contatos.pessoas_id = public.pessoas.id
		LEFT JOIN   public.enderecos ON  public.enderecos.pessoas_id = public.pessoas.id
		LEFT JOIN   public.cidades  on public.cidades.id = public.enderecos.cidades_id
		LEFT JOIN   public.estados  on public.estados.id = public.cidades.estados_id

		JOIN pedidos_status ON pedidos_status.id = pedidos.pedidos_status_id
		JOIN pedidos_categorias ON pedidos_categorias.id = pedidos.pedidos_categorias_id
		WHERE pedidos.id = $1`

	args := []interface{}{
		fmt.Sprint(o.Pedidos_Id),
	}

	err = tx.Select(&orders, query, args...)
	if err != nil {
		return
	}

	return
}
