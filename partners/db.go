package partners

import (
	"errors"
	"strconv"

	"github.com/api-qop-v2/address"
	"github.com/jmoiron/sqlx"
)

const querySELECT = ` SELECT `

const queryPartners = `
	parceiros.id,
	parceiros.he_ativo,
	parceiros.responsavel,
	parceiros.receber_pedido_por_email,
	parceiros.img, 
	COALESCE(parceiros.chave_pix, '') as chave_pix,
	COALESCE(parceiros.segunda, '') as segunda,
	COALESCE(parceiros.terca, '') as terca,
	COALESCE(parceiros.quarta, '') as quarta,
	COALESCE(parceiros.quinta, '') as quinta,
	COALESCE(parceiros.sexta, '') as sexta,
	COALESCE(parceiros.sabado, '') as sabado,
	COALESCE(parceiros.domingo, '') as domingo,
	parceiros.he_fechado,
	parceiros.monitorar_por_sse,
	parceiros.he_servico_pagamento,
	COALESCE (parceiros.he_venda_produto_digital, 0) as he_venda_produto_digital,
	
	parceiros.he_pago_pelo_site,
	parceiros.he_pago_na_entrega,
	parceiros.usa_correios,

	COALESCE ((parceiros.token_pag_seguro <> 'null' OR parceiros.token_pag_seguro <> null), true, false) AS mostrar_pagSeguro,
	COALESCE ((parceiros.token_pic_pay <> 'null' OR parceiros.token_pic_pay <> null), true, false) AS mostrar_picPay,
	COALESCE ((parceiros.chave_pix <> 'null' OR parceiros.chave_pix <> null), true, false) AS mostrar_pix,

	parceiros.finalizar_com_orcamento,
	parceiros.finalizar_com_receber_em_casa,
	parceiros.finalizar_com_retirar_na_loja,

	ramo_atividades.id as ramo_atividades_id,
	ramo_atividades.descricao as ramo_atividades_tipo,

	pessoas_contatos.contato,
	pessoas_juridicas.razao_social,
	COALESCE(pessoas_juridicas.cnpj, '') as cnpj,
	pessoas_juridicas.fantasia, 
	pessoas.id as pessoas_id,
	enderecos.id as enderecos_id,
	enderecos.cidades_id,
	COALESCE(enderecos.cep, '') as cep,
	enderecos.logradouro,
	enderecos.bairro,
	COALESCE(public.enderecos.complemento, '') as complemento,
	enderecos.numero,
	enderecos.latitude,
	enderecos.longitude,
	enderecos.area_abrangencia,
	enderecos.he_principal,
	enderecos.he_principal_parceiro,

	cidades.nome as cidade,
	estados.nome as estado,
	estados.uf as uf,

	pessoas_usuarios.email,
	pessoas_usuarios.id as usuarios_id

	FROM parceiros
	INNER JOIN pessoas ON pessoas.id = parceiros.pessoas_id
	INNER JOIN pessoas_usuarios ON pessoas_usuarios.pessoas_id =  pessoas.id 
	INNER JOIN pessoas_contatos ON  pessoas_contatos.pessoas_id =  pessoas.id 
	INNER JOIN pessoas_juridicas ON  pessoas_juridicas.pessoas_id =  pessoas.id 
	INNER JOIN enderecos ON enderecos.pessoas_id = pessoas.id
	INNER JOIN ramo_atividades ON ramo_atividades.id = parceiros.ramo_atividades_id

	INNER JOIN cidades ON cidades.id = enderecos.cidades_id
	INNER JOIN estados ON estados.id = cidades.estados_id

	WHERE enderecos.he_principal_parceiro = 1
   `

const queryPartnersPrivate = querySELECT + ` 
	COALESCE(parceiros.email_pag_seguro, '') as email_pag_seguro,
	COALESCE(parceiros.token_pag_seguro, '') as token_pag_seguro,
	COALESCE(parceiros.email_pic_pay, '') as email_pic_pay,
	COALESCE(parceiros.token_pic_pay, '') as token_pic_pay, 
	COALESCE(pessoas_usuarios.senha, '') as senha,
	` + queryPartners

func SearchPartnersAllTx(p Partners, tx *sqlx.Tx) (partner []Partners, err error) {

	query := queryPartnersPrivate
	args := []interface{}{}

	if p.Fantasia != "" {
		query = query + ` AND pessoas_juridicas.fantasia like '%' || $1 || '%' `
		args = []interface{}{
			p.Fantasia,
		}
	}

	err = tx.Select(&partner, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchPartnerFromUserIdTx(userID string, tx *sqlx.Tx) (partner []Partners, err error) {

	query := queryPartnersPrivate
	args := []interface{}{
		userID,
	}

	query = query + ` AND pessoas_usuarios.id = $1 `

	err = tx.Select(&partner, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchPartnerTx(p Partners, tx *sqlx.Tx) (partner []Partners, err error) {

	query := queryPartnersPrivate
	args := []interface{}{
		p.Parceiros_Id,
	}

	query = query + ` AND parceiros.id = $1 `

	err = tx.Select(&partner, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchPartnerFromPersonIdTx(personID string, tx *sqlx.Tx) (partner []Partners, err error) {

	query := queryPartnersPrivate
	args := []interface{}{
		personID,
	}

	query = query + ` AND pessoas.id = $1 `

	err = tx.Select(&partner, query, args...)

	if err != nil {
		return
	}

	return
}
func findByFilterDBTx(filter Filter, tx *sqlx.Tx) (user string, err error) {

	if filter.ID > 0 && tx != nil {
		user = "Seens we have a user :)"
	} else {
		err = errors.New("There's no user")
	}

	return
}

func SearchIdUserFromPartnerWithEmployeeDataTx(userIDFromEmployee string, tx *sqlx.Tx) (idUserFromPartner string, err error) {

	queryEmployee := `SELECT parceiros_funcionarios.id, parceiros_funcionarios.parceiros_id, parceiros_funcionarios.pessoas_usuarios_id
	FROM public.parceiros_funcionarios
	WHERE  parceiros_funcionarios.pessoas_usuarios_id = $1`

	argsEmployee := []interface{}{
		userIDFromEmployee,
	}

	employee := []address.PersonWeb{}
	err = tx.Select(&employee, queryEmployee, argsEmployee...)
	if err != nil {
		return
	}

	partners := []Partners{}
	queryPartner := queryPartnersPrivate
	argsPartner := []interface{}{
		employee[0].Parceiros_Id,
	}
	queryPartner = queryPartner + ` AND parceiros.id = $1`

	err = tx.Select(&partners, queryPartner, argsPartner...)
	if err != nil {
		return
	}

	idUserFromPartner = partners[0].Usuarios_Id

	return
}

func SearchRangeActivityTx(tx *sqlx.Tx) (branch []RangeActivity, err error) {

	query := `SELECT id, descricao, icone, cor_icone, background, he_ativo
	FROM public.ramo_atividades;`

	args := []interface{}{}

	err = tx.Select(&branch, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchPlansTx(p Partners, tx *sqlx.Tx) (listPlansProducts []Plano_Produto, listPlansSales []Plano_Vendas, err error) {

	args := []interface{}{
		p.Id,
	}

	queryPlansProducts := `SELECT id, parceiro_id, qtd_inicial, qtd_final, valor, he_ativo
	FROM public.planos_por_produtos_cadastrados
	WHERE parceiro_id = $1`

	queryPlansSales := `SELECT id, parceiro_id, valor_incial, valor_final, percentagem, he_ativo
	FROM public.planos_por_vendas
	WHERE parceiro_id = $1;`

	err = tx.Select(&listPlansProducts, queryPlansProducts, args...)
	if err != nil {
		return
	}

	err = tx.Select(&listPlansSales, queryPlansSales, args...)
	if err != nil {
		return
	}

	return
}

func CreateRangeTx(branch RangeActivity, tx *sqlx.Tx) (id int64, err error) {
	query := `INSERT INTO public.ramo_atividades(
			descricao,
			icone,
			cor_icone,
			background,
			he_ativo
	    )
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
		`

	args := []interface{}{
		branch.Descricao,
		branch.Icone,
		branch.Cor_Icone,
		branch.Background,
		branch.He_Ativo,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func AlterRangeActivityTx(branch RangeActivity, tx *sqlx.Tx) (err error) {

	query := `
		UPDATE public.ramo_atividades
			SET 
			descricao = $1, 
			icone = $2, 
			cor_icone = $3, 
			background = $4, 
			he_ativo = $5
		WHERE id = $6 ;
		`

	args := []interface{}{
		branch.Descricao,
		branch.Icone,
		branch.Cor_Icone,
		branch.Background,
		branch.He_Ativo,
		branch.ID,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}

	return
}

func SearchEmployeesTx(partnerID string, tx *sqlx.Tx) (p []address.PersonWeb, err error) {

	query := `SELECT 
	public.pessoas_fisicas.cpf,
	public.pessoas.nome,
	public.pessoas_fisicas.data_nascimento,
	public.pessoas_usuarios.email,
	public.pessoas.id as pessoas_id,
	public.parceiros_funcionarios.id,
	public.parceiros_funcionarios.parceiros_id,
	public.parceiros_funcionarios.pessoas_usuarios_id,
	public.parceiros_funcionarios.he_adm,
	public.parceiros_funcionarios.he_ativo
	FROM public.parceiros_funcionarios
	JOIN public.pessoas_usuarios on public.pessoas_usuarios.id = public.parceiros_funcionarios.pessoas_usuarios_id
	JOIN public.pessoas on pessoas.id = public.pessoas_usuarios.pessoas_id
	JOIN public.pessoas_fisicas on  public.pessoas_fisicas.pessoas_id = public.pessoas.id
	WHERE public.parceiros_funcionarios.parceiros_id = $1`

	args := []interface{}{
		partnerID,
	}

	err = tx.Select(&p, query, args...)

	if err != nil {
		return
	}

	return
}

func AlterHoursOfOperationTx(pessoasId string, hours Partners, tx *sqlx.Tx) (err error) {

	query := `
		UPDATE public.parceiros
		SET 
			segunda = $1,
			terca = $2,
			quarta = $3,
			quinta = $4,
			sexta = $5,
			sabado = $6, 
			domingo = $7,
			monitorar_por_sse = $8,
			he_fechado = $9
		WHERE pessoas_id = $10
		`

	args := []interface{}{
		hours.Segunda,
		hours.Terca,
		hours.Quarta,
		hours.Quinta,
		hours.Sexta,
		hours.Sabado,
		hours.Domingo,
		hours.Monitorar_Por_Sse,
		hours.He_Fechado,
		pessoasId,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}

	return
}

func SearchPartnerTypesActivesPublicTx(tx *sqlx.Tx) (ranges []RangeActivity, err error) {

	query := ` SELECT id, descricao, icone, cor_icone, background, he_ativo
	FROM public.ramo_atividades WHERE he_ativo = 1 `

	args := []interface{}{}

	err = tx.Select(&ranges, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchPartnersAllForUserTx(p Partners, tx *sqlx.Tx) (partner []Partners, err error) {

	query := ` SELECT ($1 - enderecos.latitude::DECIMAL) as comercio_mais_proximo, `
	query = query + queryPartners

	latitude, err := strconv.ParseFloat(p.Latitude, 64)
	if err != nil {
		return
	}

	args := []interface{}{
		latitude,
		p.Uf,
		p.Cidade,
	}

	query += `
		AND
			enderecos.area_abrangencia = 'País'
		OR
			enderecos.area_abrangencia = 'Estado' AND estados.uf = $2
		OR
			cidades.nome = $3
		AND parceiros.he_ativo = 1 `

	if p.Ramo_Atividades_Id > 0 {
		query = query + ` AND ramo_atividades.id = $4 `

		args = []interface{}{
			latitude,
			p.Uf,
			p.Cidade,
			p.Ramo_Atividades_Id,
		}
	}

	query += ` ORDER BY comercio_mais_proximo ASC `

	err = tx.Select(&partner, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchPartnerForUserTx(p Partners, tx *sqlx.Tx) (partner []Partners, err error) {

	query := ` SELECT ` + queryPartners
	query = query + ` AND parceiros.id = $1 `

	args := []interface{}{
		p.Parceiros_Id,
	}

	err = tx.Select(&partner, query, args...)

	if err != nil {
		return
	}

	return
}

func SummaryPurchaseTx(f Filter, and string, tx *sqlx.Tx) (resumo []ResumoVendas, err error) {

	query := `
		SELECT
		COALESCE(
			SUM(CASE WHEN (pedidos.pedidos_categorias_id = 3 AND pedidos.pedidos_status_id = 9) 
			THEN pedidos.valor_total ELSE 0 END), 0)
		AS orcamentos_total,
		
		COALESCE(
			SUM(CASE WHEN (pedidos.forma_pagamento = 'credit' AND pedidos.pedidos_categorias_id <> 3  AND pedidos.pedidos_status_id = 10)
			THEN pedidos.valor_total ELSE 0 END), 0)
		AS cartoes_total,
		
		COALESCE(
			SUM(CASE WHEN(pedidos.forma_pagamento = 'money' AND pedidos.pedidos_categorias_id <> 3 AND pedidos.pedidos_status_id = 10)
			THEN pedidos.valor_total ELSE 0  END), 0) 
		AS dinheiro_total,
		
		COALESCE(
			SUM(CASE WHEN(pedidos.pedidos_categorias_id <> 3 AND pedidos.pedidos_status_id = 10)
			THEN pedidos.valor_total ELSE 0 END), 0) 
		AS valor_total

		FROM pedidos

		WHERE pedidos.he_ativo = 1
		` + and + `
			AND CAST((SELECT pedidos.created_at at time zone 'UTC' at time zone $1) AS DATE) 
			BETWEEN  
			CAST('` + f.Data_Inicio + `' AS DATE) 
			AND 
			CAST('` + f.Data_Final + `' AS DATE);`

	args := []interface{}{
		f.Time_Zone,
	}

	err = tx.Select(&resumo, query, args...)
	if err != nil {
		return
	}

	return
}

func SummaryPurchaseAllPartnerTx(f Filter, and string, tx *sqlx.Tx) (resumo []ResumoVendas, err error) {

	query := `SELECT
		parceiros.id,
		pessoas_juridicas.fantasia,
		COALESCE(pessoas_juridicas.cnpj, '') as cnpj,

		COALESCE(
			SUM(CASE WHEN (pedidos.pedidos_categorias_id = 3 AND pedidos.pedidos_status_id = 9) 
			THEN pedidos.valor_total ELSE 0 END), 0)
		AS orcamentos_total,

		COALESCE(
		SUM(CASE WHEN (pedidos.forma_pagamento = 'credit' AND pedidos.pedidos_categorias_id <> 3  and pedidos.pedidos_status_id = 10) 
		THEN pedidos.valor_total ELSE 0 END), 0)
		AS cartoes_total,

		COALESCE(
		SUM(CASE WHEN (pedidos.forma_pagamento = 'money' AND pedidos.pedidos_categorias_id <> 3 and pedidos.pedidos_status_id = 10) 
		THEN pedidos.valor_total ELSE 0 END), 0)
		AS dinheiro_total,

		COALESCE(
		SUM(CASE WHEN (pedidos.pedidos_categorias_id <> 3 and pedidos.pedidos_status_id = 10) 
		THEN pedidos.valor_total ELSE 0 END), 0)
		AS valor_total

		FROM pedidos
		
		INNER JOIN parceiros ON parceiros.id = pedidos.parceiros_id
		INNER JOIN pessoas_juridicas ON pessoas_juridicas.pessoas_id = parceiros.pessoas_id

		WHERE parceiros.he_ativo = 1
		` + and + `
		AND CAST((SELECT pedidos.created_at at time zone 'UTC' at time zone $1) AS DATE) 
		BETWEEN  
		CAST('` + f.Data_Inicio + `' AS DATE) 
		AND 
		CAST('` + f.Data_Final + `' AS DATE)
		
		GROUP BY  parceiros.id,
		pessoas_juridicas.fantasia,
		pessoas_juridicas.cnpj`

	args := []interface{}{
		f.Time_Zone,
	}

	err = tx.Select(&resumo, query, args...)

	if err != nil {
		return
	}

	return
}

func SummaryPurchaseStatusTx(f Filter, and string, tx *sqlx.Tx) (resumo []ResumoVendas, err error) {

	query := `SELECT
        pedidos_status.descricao as descricao,
        COUNT(pedidos.id) as qtd,
        COALESCE(SUM(pedidos.valor_total), 0) as valor_total,
        pedidos_status.id
        FROM pedidos_status
        JOIN pedidos on pedidos.pedidos_status_id = pedidos_status.id

		WHERE CAST((SELECT pedidos.created_at at time zone 'UTC' at time zone $1) AS DATE) 
		BETWEEN  
		CAST('` + f.Data_Inicio + `' AS DATE) 
		AND 
		CAST('` + f.Data_Final + `' AS DATE)
		` + and + `
        GROUP BY 
        pedidos_status.id,
        pedidos_status.descricao`

	args := []interface{}{
		f.Time_Zone,
	}

	err = tx.Select(&resumo, query, args...)

	if err != nil {
		return
	}

	return
}

func SummaryPurchaseFromPartnerTx(f Filter, and string, tx *sqlx.Tx) (resumo []ResumoVendas, err error) {

	query := `SELECT
			parceiros.id,
			pessoas_juridicas.fantasia,
			COALESCE(pessoas_juridicas.cnpj, '') as cnpj,

			COALESCE(
				SUM(CASE WHEN (pedidos.pedidos_categorias_id = 3 AND pedidos.pedidos_status_id = 9) 
				THEN pedidos.valor_total ELSE 0 END), 0)
			AS orcamentos_total,

			COALESCE(
			SUM(CASE WHEN (pedidos.pedidos_categorias_id = 2 AND pedidos.pedidos_categorias_id <> 3 AND pedidos.pedidos_status_id = 10) 
			THEN pedidos.valor_total ELSE 0 END), 0)
			AS receber_em_casa_total,

			COALESCE(
			SUM(CASE WHEN (pedidos.pedidos_categorias_id = 1 AND pedidos.pedidos_categorias_id <> 3 AND pedidos.pedidos_status_id = 10) 
			THEN pedidos.valor_total ELSE 0 END), 0)
			AS retirar_na_loja_total,
			
			COALESCE(
			SUM(CASE WHEN (pedidos.pedidos_categorias_id <> 3 AND pedidos.pedidos_status_id = 10) 
			THEN pedidos.valor_total ELSE 0 END), 0)
			AS valor_total

		FROM pedidos
		INNER JOIN parceiros ON parceiros.id = pedidos.parceiros_id
		INNER JOIN pessoas_juridicas ON pessoas_juridicas.pessoas_id = parceiros.pessoas_id

		WHERE parceiros.he_ativo = 1
		` + and + `
		AND CAST((SELECT pedidos.created_at at time zone 'UTC' at time zone $1) AS DATE) 
		BETWEEN  
		CAST('` + f.Data_Inicio + `' AS DATE) 
		AND 
		CAST('` + f.Data_Final + `' AS DATE)

		GROUP BY 
			parceiros.id,
			pessoas_juridicas.fantasia,
			pessoas_juridicas.cnpj`

	args := []interface{}{
		f.Time_Zone,
	}

	err = tx.Select(&resumo, query, args...)

	if err != nil {
		return
	}

	return
}

func SummaryPurchaseAllToGraphicTx(f Filter, and string, tx *sqlx.Tx) (resumo []ResumoVendas, err error) {

	query := `SELECT
		DATE(pedidos.created_at) as dia,

		COALESCE(
			SUM(CASE WHEN (pedidos.pedidos_categorias_id = 3 AND pedidos.pedidos_status_id = 9) 
			THEN pedidos.valor_total ELSE 0 END), 0)
		AS orcamentos_total,

		COALESCE(
		SUM(CASE WHEN (pedidos.pedidos_categorias_id = 2 AND pedidos.pedidos_categorias_id <> 3 and pedidos.pedidos_status_id = 10) 
		THEN pedidos.valor_total ELSE 0 END), 0)
		AS receber_em_casa_total,
		
		COALESCE(
		SUM(CASE WHEN (pedidos.pedidos_categorias_id = 1 AND pedidos.pedidos_categorias_id <> 3 and pedidos.pedidos_status_id = 10) 
		THEN pedidos.valor_total ELSE 0 END), 0)
		AS retirar_na_loja_total,
		
		COALESCE(
		SUM(CASE WHEN (pedidos.pedidos_categorias_id <> 3 and pedidos.pedidos_status_id = 10) 
		THEN pedidos.valor_total ELSE 0 END), 0)
		AS valor_total
		
		FROM pedidos
		
		WHERE pedidos.he_ativo = 1
		` + and + `
		AND CAST((SELECT pedidos.created_at at time zone 'UTC' at time zone $1) AS DATE) 
		BETWEEN  
		CAST('` + f.Data_Inicio + `' AS DATE) 
		AND 
		CAST('` + f.Data_Final + `' AS DATE)
		
		GROUP BY  DATE(pedidos.created_at), pedidos.he_ativo`

	args := []interface{}{
		f.Time_Zone,
	}

	err = tx.Select(&resumo, query, args...)

	if err != nil {
		return
	}

	return
}

func SummaryPurchasePartnerToGraphicTx(f Filter, and string, tx *sqlx.Tx) (resumo []ResumoVendas, err error) {

	query := `SELECT

		parceiros.id,
		pessoas_juridicas.fantasia,
		COALESCE(pessoas_juridicas.cnpj, '') as cnpj,

		COALESCE(
		SUM(CASE WHEN (pedidos.pedidos_categorias_id = 3 AND pedidos.pedidos_status_id = 9) 
		THEN pedidos.valor_total ELSE 0 END), 0)
		AS orcamentos_total,

		COALESCE(
		SUM(CASE WHEN (pedidos.pedidos_categorias_id = 2 AND pedidos.pedidos_categorias_id <> 3 and pedidos.pedidos_status_id = 10) 
		THEN pedidos.valor_total ELSE 0 END), 0)
		AS receber_em_casa_total,
		
		COALESCE(
		SUM(CASE WHEN (pedidos.pedidos_categorias_id = 1 AND pedidos.pedidos_categorias_id <> 3 and pedidos.pedidos_status_id = 10) 
		THEN pedidos.valor_total ELSE 0 END), 0)
		AS retirar_na_loja_total,
		
		COALESCE(
		SUM(CASE WHEN (pedidos.pedidos_categorias_id <> 3 and pedidos.pedidos_status_id = 10) 
		THEN pedidos.valor_total ELSE 0 END), 0)
		AS valor_total
		
		FROM pedidos

		INNER JOIN parceiros ON parceiros.id = pedidos.parceiros_id
		INNER JOIN pessoas_juridicas ON pessoas_juridicas.pessoas_id = parceiros.pessoas_id
		
		WHERE pedidos.he_ativo = 1
        ` + and + `
		AND CAST((SELECT pedidos.created_at at time zone 'UTC' at time zone $1) AS DATE) 
		BETWEEN  
		CAST('` + f.Data_Inicio + `' AS DATE) 
		AND 
		CAST('` + f.Data_Final + `' AS DATE)
		
		GROUP BY
		pedidos.he_ativo,
		parceiros.id,
		pessoas_juridicas.fantasia,
		pessoas_juridicas.cnpj`

	args := []interface{}{
		f.Time_Zone,
	}

	err = tx.Select(&resumo, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchInvoicePartnerADMTx(f Filter, and string, tx *sqlx.Tx) (resumo []ResumoVendas, err error) {

	query := `
		SELECT
			parceiros.id,
			pessoas_juridicas.fantasia,
			COALESCE(pessoas_juridicas.cnpj, '') as cnpj,
			pessoas_usuarios.email,

			COALESCE((SELECT planos_por_produtos_cadastrados.id
			FROM planos_por_produtos_cadastrados
			WHERE planos_por_produtos_cadastrados.parceiro_id = parceiros.id
			AND
			planos_por_produtos_cadastrados.he_ativo = 1
			LIMIT 1), 0) AS planos_produtos,

			COALESCE((SELECT planos_por_vendas.id
			FROM planos_por_vendas
			WHERE planos_por_vendas.parceiro_id = parceiros.id
			AND
			planos_por_vendas.he_ativo = 1
			LIMIT 1), 0) as planos_vendas,

			(SELECT COUNT(produtos.id) FROM produtos WHERE produtos.parceiros_id = parceiros.id) AS qtd_produtos,

			COALESCE((SELECT
				SUM(CASE WHEN(pedidos.pedidos_categorias_id = 3 AND pedidos.pedidos_status_id = 9) THEN pedidos.valor_total ELSE 0 END)
			FROM pedidos WHERE pedidos.parceiros_id = parceiros.id
			AND date_part('year', pedidos.created_at) = $1 
			AND date_part('month', pedidos.created_at) = $2
			AND pedidos.he_ativo = 1 ` + and + ` ), 0) AS total_orcamento,

			COALESCE((SELECT
				SUM(CASE WHEN (pedidos.pedidos_categorias_id <> 3 and pedidos.pedidos_status_id = 10) THEN pedidos.valor_total ELSE 0 END)
			FROM pedidos WHERE pedidos.parceiros_id = parceiros.id
			AND date_part('year', pedidos.created_at) = $1 
			AND date_part('month', pedidos.created_at) = $2
			AND pedidos.he_ativo = 1 ` + and + `), 0) AS total_vendas

			FROM  parceiros
			LEFT JOIN pedidos ON pedidos.parceiros_id = parceiros.id
			JOIN pessoas_usuarios ON  pessoas_usuarios.pessoas_id = parceiros.pessoas_id
			JOIN pessoas_juridicas ON  pessoas_juridicas.pessoas_id = pessoas_usuarios.pessoas_id
			
			WHERE parceiros.he_ativo = 1
			` + and + `
			GROUP BY
			parceiros.id,
			pessoas_juridicas.fantasia,
			pessoas_juridicas.cnpj,
			pessoas_usuarios.email`

	args := []interface{}{
		f.Ano,
		f.Mes,
	}

	err = tx.Select(&resumo, query, args...)
	if err != nil {
		return
	}

	return
}

func SearchPlansPartnerForValueTx(p Filter, tx *sqlx.Tx) (plan []Plano_Vendas, err error) {

	query := `SELECT id, 
	parceiro_id,
	 valor_incial,
	 valor_final,
	 percentagem,
	 he_ativo FROM planos_por_vendas 
	 WHERE parceiro_id = $1 AND he_ativo = 1`
	args := []interface{}{
		p.Parceiros_Id,
	}

	err = tx.Select(&plan, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchPlansPartnerForProductTx(f Filter, tx *sqlx.Tx) (plan []Plano_Produto, err error) {

	query := `SELECT id
	, parceiro_id
	, qtd_inicial
	, qtd_final
	, valor
	, he_ativo FROM planos_por_produtos_cadastrados 
	WHERE parceiro_id = $1 AND he_ativo = 1`
	args := []interface{}{
		f.Parceiros_Id,
	}

	err = tx.Select(&plan, query, args...)

	if err != nil {
		return
	}

	return
}

func CreatePotentialPartnersTx(f Filter, tx *sqlx.Tx) (id int64, err error) {

	query := `
		INSERT INTO public.parceiros_leads (
			nome, email, telefone
		)
		VALUES($1, $2, $3)
		RETURNING id;
	`

	args := []interface{}{
		f.Nome,
		f.Email,
		f.Telefone,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}
