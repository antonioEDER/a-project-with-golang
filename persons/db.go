package persons

import (
	"fmt"
	"strings"

	"github.com/api-qop-v2/address"
	"github.com/labstack/echo"

	"github.com/jmoiron/sqlx"
)

func searchTypesPersonTx(person address.PersonWeb, tx *sqlx.Tx) (personWeb []address.PersonWeb, err error) {

	query := `
		SELECT id, tipo
		FROM public.tipos_pessoas 
		WHERE 
		tipo = $1;
		`

	args := []interface{}{
		person.Tipo,
	}

	err = tx.Select(&personWeb, query, args...)

	if err != nil {
		return
	}

	return
}

func searchPersonExistsTx(person address.PersonWeb, tx *sqlx.Tx) (personWeb []address.PersonWeb, err error) {

	query := `
		SELECT 
			id,
			email,
			he_ativo
		FROM public.pessoas_usuarios 
		WHERE 
			email = $1;
		`

	args := []interface{}{
		person.Email,
	}

	err = tx.Select(&personWeb, query, args...)

	if err != nil {
		return
	}

	return
}

func CreatePersonDBTx(person address.PersonWeb, active uint64, tx *sqlx.Tx) (id int64, err error) {

	query := `
		INSERT INTO public.pessoas (
			nome, 
			he_ativo
	    ) VALUES ($1, $2)
		RETURNING id;
		`

	args := []interface{}{
		person.Nome,
		active,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func CreatePersonAddressDBTx(idPerson int64, personAddress address.PersonWeb, tx *sqlx.Tx) (id int64, err error) {

	cities, err := address.SearchCity(personAddress.Uf, personAddress.Cidade, tx)
	if err != nil {
		return
	}
	if len(cities) <= 0 {
		err = echo.NewHTTPError(401, "Não encontramos uma cidade")
		return
	}

	query := `
		INSERT INTO public.enderecos (
			cidades_id, 
			pessoas_id, 
			cep, 
			logradouro, 
			bairro, 
			complemento, 
			numero, 
			latitude, 
			longitude, 
			area_abrangencia, 
			he_principal, 
			he_principal_parceiro
		)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id;
	`

	args := []interface{}{
		cities[0].Id,
		idPerson,
		personAddress.Cep,
		personAddress.Logradouro,
		personAddress.Bairro,
		personAddress.Complemento,
		personAddress.Numero,
		personAddress.Latitude,
		personAddress.Longitude,
		personAddress.Area_Abrangencia,
		1,
		personAddress.He_Principal_Parceiro,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func CreatePersonMigrationDBTx(idUser int64, tx *sqlx.Tx) (id int64, err error) {
	query := `
		INSERT INTO public.usuarios_migracao(
			pessoas_usuarios_id
		)
		VALUES ($1)
		RETURNING id;
		`

	args := []interface{}{
		idUser,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func CreatePersonUserDBTx(idPerson int64, codValidate string, person address.PersonWeb, tx *sqlx.Tx) (id int64, err error) {
	query := `
		INSERT INTO public.pessoas_usuarios(
			pessoas_id, 
			email, 
			senha, 
			uid, 
			rede_social, 
			he_ativo,
			cod_confirmacao
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
		`

	activeUser := 0
	if person.Uid != "" {
		activeUser = 1
	}
	args := []interface{}{
		idPerson,
		person.Email,
		person.Senha,
		person.Uid,
		person.Nome_Rede_Social,
		activeUser,
		fmt.Sprint(codValidate),
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func CreatePersonStatusTx(idPerson int64, person address.PersonWeb, tx *sqlx.Tx) (id int64, err error) {
	query := `
		INSERT INTO public.pessoas_status(
			pessoas_id, 
			descricao, 
			obs,
			he_ativo)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
		`

	args := []interface{}{
		idPerson,
		person.Descricao,
		person.Obs,
		1,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func CreateLegalPersonTx(idPerson int64, person address.PersonWeb, tx *sqlx.Tx) (id int64, err error) {
	query := `
		INSERT INTO public.pessoas_juridicas(
			pessoas_id, 
			razao_social, 
			fantasia, 
			cnpj, 
			he_ativo, 
			tipos_pessoas_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
		`

	args := []interface{}{
		idPerson,
		person.Razao_Social,
		person.Fantasia,
		person.Cnpj,
		1,
		person.Tipos_Pessoas_Id,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func CreatePhysicalPersonTx(idPerson int64, person address.PersonWeb, active uint64, tx *sqlx.Tx) (id int64, err error) {
	query := `
		INSERT INTO public.pessoas_fisicas(
			pessoas_id, 
			cpf, 
			data_nascimento,
			he_ativo, 
			tipos_pessoas_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
		`

	args := []interface{}{
		idPerson,
		person.Cpf,
		person.Data_Nascimento,
		active,
		person.Tipos_Pessoas_Id,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func CreateContactPersonTx(idPerson int64, person address.PersonWeb, activeUser uint64, tx *sqlx.Tx) (id int64, err error) {
	query := `
		INSERT INTO public.pessoas_contatos(
			pessoas_id,
			tipo,
			contato,
			he_ativo)
		VALUES 
			($1, $2, $3, $4)
		RETURNING id;
		`

	if person.Uid != "" {
		activeUser = 1
	}

	args := []interface{}{
		idPerson,
		"TELEFONE",
		person.Celular,
		activeUser,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func AlterUserPersonTx(person address.PersonWeb, tx *sqlx.Tx) (err error) {

	queryPersonUser := `
		UPDATE public.pessoas_fisicas
			SET 
				data_nascimento = $1
		WHERE pessoas_id = $2;
		`

	queryPersonUserCPF := `
		UPDATE public.pessoas_fisicas
			SET 
				cpf = $1
		WHERE pessoas_id = $2;
		`

	queryPersonContact := `
		UPDATE public.pessoas_contatos
			SET contato = $1
		WHERE pessoas_id = $2;
		`

	queryPerson := `
		UPDATE public.pessoas
			SET nome = $1
		WHERE id = $2;
		`

	argsPersonUser := []interface{}{
		person.Data_Nascimento,
		person.Id,
	}

	argsPersonContact := []interface{}{
		person.Contato,
		person.Id,
	}

	argsPerson := []interface{}{
		person.Nome,
		person.Id,
	}

	argsPersonCPF := []interface{}{
		person.Cpf,
		person.Id,
	}

	_, err = tx.Exec(queryPersonUser, argsPersonUser...)
	if err != nil {
		return err
	}

	_, err = tx.Exec(queryPersonContact, argsPersonContact...)
	if err != nil {
		return err
	}

	_, err = tx.Exec(queryPerson, argsPerson...)
	if err != nil {
		return err
	}

	_, _ = tx.Exec(queryPersonUserCPF, argsPersonCPF...)

	return
}

func AlterContactPersonTx(idPerson int64, person address.PersonWeb, tx *sqlx.Tx) (err error) {
	query := `
	  UPDATE public.pessoas_contatos
	    SET
			contato  = $1,
			he_ativo = $2
		WHERE id = $3;
		`

	active := 0

	if person.He_Ativo {
		active = 1
	}
	args := []interface{}{
		person.Contato,
		active,
		person.Pessoas_Contatos_Id,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}

	return err
}

func AlterPhysicalPersonTx(idPerson int64, person address.PersonWeb, tx *sqlx.Tx) (err error) {

	active := 0
	if person.He_Ativo {
		active = 1
	}

	query := ``
	queryUp := ` UPDATE public.pessoas_fisicas `

	querySet := ` SET
	data_nascimento = $1, 
	he_ativo = $2 `

	queryCpf := `, cpf = $3 `

	queryWhereNotCpf := ` WHERE pessoas_id = $3 `
	queryWhereYesCpf := ` WHERE pessoas_id = $4 `

	args := []interface{}{}

	if person.Cpf == "" {
		query = queryUp + querySet + queryWhereNotCpf
		args = []interface{}{
			person.Data_Nascimento,
			active,
			idPerson,
		}
	} else {
		query = queryUp + querySet + queryCpf + queryWhereYesCpf
		args = []interface{}{
			person.Data_Nascimento,
			active,
			person.Cpf,
			idPerson,
		}
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func AlterLegalPersonTx(idPerson int64, person address.PersonWeb, tx *sqlx.Tx) (err error) {
	query := `
	UPDATE public.pessoas_juridicas
		SET
			razao_social = $1, 
			fantasia = $2, 
			cnpj = $3, 
			he_ativo  = $4
		WHERE pessoas_id = $5
		`
	active := 0
	if person.He_Ativo {
		active = 1
	}

	args := []interface{}{
		person.Razao_Social,
		person.Fantasia,
		person.Cnpj,
		active,
		idPerson,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func AlterPersonUserPassOrEmailDBTx(idPerson int64, person address.PersonWeb, tx *sqlx.Tx) (err error) {
	queryAll := `
		UPDATE public.pessoas_usuarios
			SET
				email = $1,
				senha = $2,
				he_ativo = $3 
		WHERE pessoas_id = $4`

	queryEmail := `
		UPDATE public.pessoas_usuarios
			SET
				email = $1,
				he_ativo = $2 
		WHERE pessoas_id = $3`

	queryPass := `
		UPDATE public.pessoas_usuarios
			SET
			    senha = $1,
				he_ativo = $2 
		WHERE pessoas_id = $3`

	args := []interface{}{}
	query := ``

	active := 0
	if person.He_Ativo {
		active = 1
	}

	if person.Email != "" && person.Senha != "" {
		query = queryAll
		args = []interface{}{
			person.Email,
			person.Senha,
			active,
			idPerson,
		}
	}
	if person.Email != "" && person.Senha == "" {
		query = queryEmail
		args = []interface{}{
			person.Email,
			active,
			idPerson,
		}
	}
	if person.Email == "" && person.Senha != "" {
		query = queryPass
		args = []interface{}{
			person.Senha,
			active,
			idPerson,
		}
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func AlterPersonUserDBTx(idPerson int64, person address.PersonWeb, tx *sqlx.Tx) (err error) {
	queryUp := ` UPDATE public.pessoas_usuarios
		SET
			email = $1, 
			he_ativo = $2 `
	queryPass := ` , senha = $4 `
	queryWhere := ` WHERE pessoas_id = $3 `
	active := 0
	if person.He_Ativo {
		active = 1
	}

	args := []interface{}{}
	query := ""

	if person.Senha == "" {
		args = []interface{}{
			person.Email,
			active,
			idPerson,
		}
		query = queryUp + queryWhere
	} else {
		args = []interface{}{
			person.Email,
			active,
			idPerson,
			person.Senha,
		}
		query = queryUp + queryPass + queryWhere
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func AlterPersonAddressDBTx(idPerson int64, personAddress address.PersonWeb, tx *sqlx.Tx) (err error) {

	cities, err := address.SearchCity(personAddress.Uf, personAddress.Cidade, tx)
	if err != nil {
		return
	}
	if len(cities) <= 0 {
		err = echo.NewHTTPError(401, "Não encontramos uma cidade")
		return
	}

	query := `
	UPDATE public.enderecos
		SET
			cidades_id = $1, 
			cep = $2, 
			logradouro = $3, 
			bairro = $4, 
			complemento = $5, 
			numero = $6, 
			latitude = $7, 
			longitude = $8, 
			area_abrangencia = $9, 
			he_principal = $10, 
			he_principal_parceiro = $11
		WHERE pessoas_id = $12
	`

	args := []interface{}{
		cities[0].Id,
		personAddress.Cep,
		personAddress.Logradouro,
		personAddress.Bairro,
		personAddress.Complemento,
		personAddress.Numero,
		personAddress.Latitude,
		personAddress.Longitude,
		personAddress.Area_Abrangencia,
		1,
		personAddress.He_Principal_Parceiro,
		idPerson,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func CreateTransportTx(transport address.PersonWeb, tx *sqlx.Tx) (id int64, err error) {
	query := `
	INSERT INTO public.transportadoras(
		parceiros_id,
		pessoas_id,
		he_entrega_propria,
		he_ativo)
		VALUES ($1, $2, $3, $4)
		RETURNING id
		`
	active := 0
	if transport.He_Ativo {
		active = 1
	}

	heOwnDelivery := 0
	if transport.He_Entrega_Propria {
		heOwnDelivery = 1
	}

	args := []interface{}{
		transport.Parceiros_Id,
		transport.Pessoas_Id,
		heOwnDelivery,
		active,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func CreatePartnerTransportTx(transport address.PersonWeb, tx *sqlx.Tx) (id int64, err error) {
	query := `
	INSERT INTO public.parceiros_transportadoras(
		parceiros_id,
		transportadoras_id,
		he_ativo)
		VALUES ($1, $2, $3)
		RETURNING id
		`
	active := 0
	if transport.He_Ativo {
		active = 1
	}

	args := []interface{}{
		transport.Parceiros_Id,
		transport.Transportadoras_Id,
		active,
	}

	err = tx.Get(&id, query, args...)

	if err != nil {
		return
	}

	return id, err
}

func CreatePersonPartnersTx(person address.PersonWeb, tx *sqlx.Tx) (id int64, err error) {
	query := `
	INSERT INTO public.parceiros(
		ramo_atividades_id, 
		pessoas_id, 
		img, 
		responsavel, 
		segunda, 
		terca, 
		quarta, 
		quinta, 
		sexta, 
		sabado, 
		domingo, 
		email_pag_seguro, 
		token_pag_seguro, 
		email_pic_pay, 
		token_pic_pay, 
		chave_pix, 
		receber_pedido_por_email, 
		finalizar_com_orcamento, 
		finalizar_com_receber_em_casa, 
		finalizar_com_retirar_na_loja, 
		usa_correios, 
		monitorar_por_sse, 
		he_pago_na_entrega, 
		he_pago_pelo_site, 
		he_servico_pagamento, 
		he_fechado, 
		he_ativo,
		he_venda_produto_digital)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
			$21, $22, $23, $24, $25, $26, $27, $28
		)
		RETURNING id
		`
	active := 0
	if person.He_Ativo {
		active = 1
	}

	heSse := 0
	if person.Monitorar_Por_Sse {
		heSse = 1
	}

	args := []interface{}{
		person.Ramo_Atividades_Id,
		person.Pessoas_Id,
		person.Img,
		person.Responsavel,
		person.Segunda,
		person.Terca,
		person.Quarta,
		person.Quinta,
		person.Sexta,
		person.Sabado,
		person.Domingo,
		person.Email_Pag_Seguro,
		person.Token_Pag_Seguro,
		person.Email_Pic_Pay,
		person.Token_Pic_Pay,
		person.Chave_Pix,
		person.Receber_Pedido_Por_Email,
		person.Finalizar_Com_Orcamento,
		person.Finalizar_Com_Receber_Em_Casa,
		person.Finalizar_Com_Retirar_Na_Loja,
		person.Usa_Correios,
		heSse,
		person.He_Pago_Na_Entrega,
		person.He_Pago_Pelo_Site,
		person.He_Servico_Pagamento,
		person.He_Fechado,
		active,
		person.He_Venda_Produto_Digital,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func CreatePlanForSaleTx(active bool, person address.PersonWeb, tx *sqlx.Tx) (id int64, err error) {
	query := `
		INSERT INTO public.planos_por_vendas(
			parceiro_id, 
			valor_incial, 
			valor_final,
			percentagem,
			he_ativo)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
		`

	he_ativo := 0
	if active {
		he_ativo = 1
	}

	for _, plane := range person.Plano_Vendas {
		args := []interface{}{
			person.Parceiros_Id,
			plane.Valor_Inicial,
			plane.Valor_Final,
			plane.Percentagem,
			he_ativo,
		}
		err = tx.Get(&id, query, args...)
		if err != nil {
			return
		}
	}

	return id, err
}

func CreatePlanForProductTx(active bool, person address.PersonWeb, tx *sqlx.Tx) (id int64, err error) {
	query := `
		INSERT INTO public.planos_por_produtos_cadastrados(
			parceiro_id, 
			qtd_inicial, 
			qtd_final,
			valor,
			he_ativo)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
		`
	he_ativo := 0
	if active {
		he_ativo = 1
	}

	for _, plane := range person.Plano_Produto {
		args := []interface{}{
			person.Parceiros_Id,
			plane.Qtd_Inicial,
			plane.Qtd_Final,
			plane.Valor,
			he_ativo,
		}

		err = tx.Get(&id, query, args...)
		if err != nil {
			return
		}
	}

	return id, err
}

func AlterPersonDBTx(person address.PersonWeb, tx *sqlx.Tx) (id int64, err error) {

	query := `
	UPDATE  public.pessoas 
		SET
			nome = $1, 
			he_ativo  = $2
	WHERE id = $3;
		`

	active := 1
	if !person.He_Ativo {
		active = 0
	}
	args := []interface{}{
		person.Nome,
		active,
		person.Pessoas_Id,
	}

	_, err = tx.Exec(query, args...)

	if err != nil {
		return
	}

	return id, err
}

func AlterPersonPartnersTx(person address.PersonWeb, tx *sqlx.Tx) (id int64, err error) {
	query := `
	UPDATE  public.parceiros
		SET 
			ramo_atividades_id = $1, 
			pessoas_id = $2, 
			img = $3, 
			responsavel = $4, 
			segunda = $5, 
			terca = $6, 
			quarta = $7, 
			quinta = $8, 
			sexta = $9, 
			sabado = $10, 
			domingo = $11, 
			email_pag_seguro = $12, 
			token_pag_seguro = $13, 
			email_pic_pay = $14, 
			token_pic_pay = $15, 
			chave_pix = $16, 
			receber_pedido_por_email = $17, 
			finalizar_com_orcamento = $18, 
			finalizar_com_receber_em_casa = $19, 
			finalizar_com_retirar_na_loja = $20, 
			usa_correios = $21, 
			monitorar_por_sse = $22, 
			he_pago_na_entrega = $23, 
			he_pago_pelo_site = $24, 
			he_servico_pagamento = $25, 
			he_fechado = $26, 
			he_ativo = $27,
			he_venda_produto_digital = $28
		WHERE id = $29;
		`
	active := 0
	if person.He_Ativo {
		active = 1
	}

	heSse := 0
	if person.Monitorar_Por_Sse {
		heSse = 1
	}

	args := []interface{}{
		person.Ramo_Atividades_Id,
		person.Pessoas_Id,
		person.Img,
		person.Responsavel,
		person.Segunda,
		person.Terca,
		person.Quarta,
		person.Quinta,
		person.Sexta,
		person.Sabado,
		person.Domingo,
		person.Email_Pag_Seguro,
		person.Token_Pag_Seguro,
		person.Email_Pic_Pay,
		person.Token_Pic_Pay,
		person.Chave_Pix,
		person.Receber_Pedido_Por_Email,
		person.Finalizar_Com_Orcamento,
		person.Finalizar_Com_Receber_Em_Casa,
		person.Finalizar_Com_Retirar_Na_Loja,
		person.Usa_Correios,
		heSse,
		person.He_Pago_Na_Entrega,
		person.He_Pago_Pelo_Site,
		person.He_Servico_Pagamento,
		person.He_Fechado,
		active,
		person.He_Venda_Produto_Digital,
		person.Parceiros_Id,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return id, err
}

func AlterPersonStatusTx(idPerson int64, person address.PersonWeb, tx *sqlx.Tx) (err error) {
	query := `
	UPDATE public.pessoas_status
		SET
			descricao = $1, 
			obs = $2,
			he_ativo = $3
	WHERE pessoas_id = $4;`

	args := []interface{}{
		person.Descricao,
		person.Obs,
		1,
		idPerson,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func AlterLinkPartnerTransportTx(idPerson int64, person address.PersonWeb, tx *sqlx.Tx) (err error) {
	listTransport := strings.Join(person.Lista_Transportadoras_Ids, ",")

	queryUp := `
	UPDATE public.parceiros_transportadoras
		SET
			he_ativo = 0
		WHERE parceiros_id = $1 AND transportadoras_id NOT IN (` + listTransport + `)`

	argsUpdate := []interface{}{
		person.Parceiros_Id,
	}
	_, err = tx.Exec(queryUp, argsUpdate...)
	if err != nil {
		return
	}

	var id int64

	for _, idTransport := range person.Lista_Transportadoras_Ids {
		queryInsert := `
		INSERT INTO public.parceiros_transportadoras (parceiros_id, transportadoras_id, he_ativo)
		SELECT $1, $2, 1
			WHERE NOT EXISTS (
				SELECT 
					public.parceiros_transportadoras.id 
				FROM public.parceiros_transportadoras 
			    WHERE public.parceiros_transportadoras.transportadoras_id = $2
				AND  public.parceiros_transportadoras.parceiros_id = $1
				AND  public.parceiros_transportadoras.he_ativo = 1
			)`

		argsInsert := []interface{}{
			person.Parceiros_Id,
			idTransport,
		}

		_ = tx.Get(&id, queryInsert, argsInsert...)

	}
	return err
}

func AlterPlanForProductTx(active bool, person address.PersonWeb, tx *sqlx.Tx) (err error) {
	query := `
	UPDATE public.planos_por_produtos_cadastrados
		SET
			qtd_inicial = $1, 
			qtd_final = $2,
			valor = $3,
			he_ativo = $4
	WHERE id = $5;
		`
	he_ativo := 0
	if active {
		he_ativo = 1
	}

	for _, plane := range person.Plano_Produto {
		args := []interface{}{
			plane.Qtd_Inicial,
			plane.Qtd_Final,
			plane.Valor,
			he_ativo,
			plane.Id,
		}

		_, err = tx.Exec(query, args...)
		if err != nil {
			return
		}
	}

	return err
}

func AlterPlanForSaleTx(active bool, person address.PersonWeb, tx *sqlx.Tx) (err error) {
	query := `
	UPDATE public.planos_por_vendas
		SET
			valor_incial = $1, 
			valor_final = $2,
			percentagem = $3,
			he_ativo = $4
		WHERE id = $5;
		`

	he_ativo := 0
	if active {
		he_ativo = 1
	}

	for _, plane := range person.Plano_Vendas {
		args := []interface{}{
			plane.Valor_Inicial,
			plane.Valor_Final,
			plane.Percentagem,
			he_ativo,
			plane.Id,
		}
		_, err = tx.Exec(query, args...)
		if err != nil {
			return
		}
	}

	return err
}

func AlterPersonToAddTokenPushTx(person address.PersonWeb, tx *sqlx.Tx) (err error) {

	query := `
		UPDATE public.pessoas_usuarios
			SET token_push_web = $1
		WHERE pessoas_id = $2;
		`

	args := []interface{}{
		person.Token_Push_Web,
		person.Id,
	}

	if person.Token_Push_App != "" {
		query = `
		UPDATE public.pessoas_usuarios
			SET token_push_app = $1
		WHERE pessoas_id = $2;
		`
		args = []interface{}{
			person.Token_Push_App,
			person.Id,
		}
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}

	return
}
