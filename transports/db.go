package transports

import (
	"github.com/jmoiron/sqlx"
)

func SearchTransportsByActiveForPartnerTx(transport Transport, tx *sqlx.Tx) (resultAll []Transport, resultActive []Transport, err error) {

	args := []interface{}{
		transport.Parceiros_Id,
	}

	queryAll := `
	SELECT 
		transportadoras.id,
		transportadoras.parceiros_id, 
		transportadoras.pessoas_id, 
		transportadoras.he_entrega_propria, 
		transportadoras.he_ativo,
		pessoas_juridicas.fantasia,
		pessoas_juridicas.razao_social
	FROM public.transportadoras
		JOIN pessoas_juridicas ON pessoas_juridicas.pessoas_id = public.transportadoras.pessoas_id
	WHERE transportadoras.he_ativo = 1 
	AND (transportadoras.he_entrega_propria = 0 OR transportadoras.parceiros_id = $1)
	`

	queryActive := `
	SELECT 
		parceiros_transportadoras.id, 
		parceiros_transportadoras.parceiros_id, 
		parceiros_transportadoras.transportadoras_id, 
		parceiros_transportadoras.he_ativo,
		pessoas_juridicas.fantasia,
		pessoas_juridicas.razao_social
	FROM public.parceiros_transportadoras
	JOIN transportadoras ON transportadoras.id = parceiros_transportadoras.transportadoras_id
	JOIN pessoas ON pessoas.id = transportadoras.pessoas_id
    JOIN pessoas_juridicas ON pessoas_juridicas.pessoas_id = pessoas.id
    WHERE parceiros_transportadoras.parceiros_id = $1 AND parceiros_transportadoras.he_ativo = 1`

	err = tx.Select(&resultAll, queryAll, args...)
	err = tx.Select(&resultActive, queryActive, args...)

	if err != nil {
		return
	}

	return
}

func SearchTransportsActiveTx(transport Transport, tx *sqlx.Tx) (result []Transport, err error) {

	query := `;
	SELECT 
		public.transportadoras.id,
		COALESCE(public.transportadoras.parceiros_id, 0) as parceiros_id,
		COALESCE(public.transportadoras.pessoas_id, 0) as pessoas_id,
		public.transportadoras.he_entrega_propria,
		public.transportadoras.he_ativo,

		COALESCE(public.enderecos.id, 0) as enderecos_id ,
		COALESCE(public.enderecos.latitude, '') as latitude,
		COALESCE(public.enderecos.longitude, '') as longitude,
		COALESCE(public.enderecos.cep, '') as cep,
		COALESCE(public.enderecos.logradouro, '') as logradouro,
		COALESCE(public.enderecos.numero, '') as numero,
		COALESCE(public.enderecos.bairro, '') as bairro,

		public.pessoas_juridicas.fantasia,
		public.pessoas_juridicas.razao_social,
		COALESCE(pessoas_juridicas.cnpj, '') as cnpj,

		COALESCE(public.pessoas_usuarios.id, 0) as usuarios_id,
		COALESCE(public.pessoas_usuarios.email , '') as email,

		COALESCE(public.pessoas_contatos.contato, '') as email,
		COALESCE(public.pessoas_contatos.id, 0) as pessoas_contatos_id,

		COALESCE(public.estados.uf , '') as uf,
		COALESCE(public.estados.nome, '') as estado,
		COALESCE(public.cidades.nome, '') as cidade

		FROM public.transportadoras
		LEFT JOIN public.pessoas on public.pessoas.id = public.transportadoras.pessoas_id
		LEFT JOIN public.pessoas_juridicas on public.pessoas_juridicas.pessoas_id = public.pessoas.id
		LEFT JOIN public.pessoas_usuarios on public.pessoas.id = public.pessoas_usuarios.pessoas_id
		LEFT JOIN public.pessoas_contatos on public.pessoas.id = public.pessoas_contatos.pessoas_id

		LEFT JOIN public.enderecos on public.pessoas.id = public.enderecos.pessoas_id
		LEFT JOIN public.cidades on public.enderecos.cidades_id = public.cidades.id
		LEFT JOIN public.estados on public.cidades.estados_id = public.estados.id
		LEFT JOIN public.paises on public.estados.paises_id = public.paises.id
	`

	args := []interface{}{}

	err = tx.Select(&result, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchTransportsByActiveForUserTx(transport Transport, tx *sqlx.Tx) (listTransports []Transport, err error) {

	query := `
		SELECT 
			COALESCE((transportadoras.id), 0) AS transportadoras_id,
			COALESCE((transportadoras.pessoas_id), 0) AS pessoas_id,
			COALESCE((transportadoras_servicos.codigo), '') AS codigo,
			COALESCE((transportadoras_servicos.id ), 0) AS transportadoras_servicos_id,
			COALESCE((transportadoras_servicos.descricao), '') AS descricao,
			COALESCE((transportadoras_servicos.prazo), '') AS prazo,
			COALESCE((enderecos.id), 0) AS enderecos_id,
			COALESCE((pessoas_juridicas.fantasia), '') AS fantasia,
			COALESCE((parceiros_transportadoras.parceiros_id ), 0) AS parceiros_id
	FROM parceiros_transportadoras
	left JOIN transportadoras ON transportadoras.id = parceiros_transportadoras.transportadoras_id
	left JOIN transportadoras_servicos on transportadoras_servicos.transportadoras_id = transportadoras.id
	JOIN pessoas ON pessoas.id = transportadoras.pessoas_id
	JOIN pessoas_juridicas ON pessoas_juridicas.pessoas_id = pessoas.id
	INNER JOIN enderecos on enderecos.pessoas_id = pessoas.id
	WHERE parceiros_transportadoras.parceiros_id = $1
	AND transportadoras.he_ativo = 1
		AND parceiros_transportadoras.he_ativo = 1
	ORDER BY transportadoras_servicos.descricao`

	args := []interface{}{
		transport.Parceiros_Id,
	}

	err = tx.Select(&listTransports, query, args...)
	if err != nil {
		return
	}

	return
}

func SearchServicesTransportsByActiveForUserTx(transport Transport, km int, tx *sqlx.Tx) (listServices []Transport, err error) {

	query := `SELECT 
				COALESCE((transportadoras_servicos.id), 0) AS id,
				COALESCE((servicos_por_valor.valor), 0) AS valor,
				COALESCE((servicos_por_km.km), 0) AS km,
				COALESCE((servicos_por_km.frete+servicos_por_valor.frete), 0) AS frete
			FROM transportadoras_servicos
			LEFT JOIN servicos_por_valor on transportadoras_servicos.id = servicos_por_valor.transportadoras_servicos_id
			LEFT JOIN servicos_por_km on transportadoras_servicos.id = servicos_por_km.transportadoras_servicos_id
			WHERE transportadoras_servicos.id = $1
				AND servicos_por_km.km <= $2
				AND transportadoras_servicos.he_ativo = 1
				AND servicos_por_km.he_ativo = 1 
				AND servicos_por_valor.he_ativo = 1
			ORDER BY servicos_por_valor.valor DESC LIMIT 1`
	// AND servicos_por_valor.valor <= $3
	// 		transport.Valor_Pedido,

	args := []interface{}{
		transport.Transportadoras_Servicos_Id,
		km,
	}

	err = tx.Select(&listServices, query, args...)
	if err != nil {
		return
	}

	return
}

func CreateTransportTx(transport Transport, tx *sqlx.Tx) (id int64, err error) {
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

	heOwnDelivery := 1
	if !transport.He_Entrega_Propria {
		heOwnDelivery = 0
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

func AlterTransportTx(transport Transport, tx *sqlx.Tx) (err error) {
	query := `
	UPDATE public.transportadoras
		SET
			he_entrega_propria = $1,
			he_ativo = $2
		WHERE  public.transportadoras.id = $3
		`

	heOwnDelivery := 1
	if !transport.He_Entrega_Propria {
		heOwnDelivery = 0
	}

	active := 0
	if transport.He_Ativo {
		active = 1
	}

	args := []interface{}{
		heOwnDelivery,
		active,
		transport.ID,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func CreateServiceTx(transport TransportService, tx *sqlx.Tx) (id int64, err error) {
	query := `
	INSERT INTO public.transportadoras_servicos(
		transportadoras_id,
		descricao,
		codigo,
		detalhes,
		prazo,
		he_ativo,
		parceiros_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	active := 0
	if transport.He_Ativo {
		active = 1
	}

	args := []interface{}{
		transport.Transportadoras_Id,
		transport.Descricao,
		transport.Codigo,
		transport.Detalhes,
		transport.Prazo,
		active,
		transport.Parceiros_Id,
	}

	if transport.Parceiros_Id == 0 {
		args = []interface{}{
			transport.Transportadoras_Id,
			transport.Descricao,
			transport.Codigo,
			transport.Detalhes,
			transport.Prazo,
			active,
			nil,
		}
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func AlterServiceTx(transport TransportService, tx *sqlx.Tx) (err error) {
	query := `
	  UPDATE public.transportadoras_servicos
		SET
			transportadoras_id = $1,
			descricao = $2,
			codigo = $3,
			detalhes = $4,
			prazo = $5,
			he_ativo = $6
		WHERE  public.transportadoras_servicos.id = $7	`

	active := 0
	if transport.He_Ativo {
		active = 1
	}

	args := []interface{}{
		transport.Transportadoras_Id,
		transport.Descricao,
		transport.Codigo,
		transport.Detalhes,
		transport.Prazo,
		active,
		transport.ID,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func SearchServicesTx(transport TransportService, tx *sqlx.Tx) (results []TransportService, err error) {

	query := `
		SELECT 
			id,
			transportadoras_id,
			descricao,
			codigo,
			COALESCE(detalhes, '') as detalhes,
			valor,
			prazo,
			he_ativo,
			created_at,
			created_by,
			updated_at,
			updated_by,
			deleted
		FROM public.transportadoras_servicos
	`

	args := []interface{}{}

	err = tx.Select(&results, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchServicesByIdPartnerTx(transport TransportService, tx *sqlx.Tx) (results []TransportService, err error) {

	query := `
		SELECT 
			id,
			COALESCE(parceiros_id, '0') as parceiros_id,
			transportadoras_id,
			descricao,
			codigo,
			COALESCE(detalhes, '') as detalhes,
			valor,
			prazo,
			he_ativo
		FROM public.transportadoras_servicos
		WHERE  public.transportadoras_servicos.transportadoras_id = $1
	`

	args := []interface{}{
		transport.Transportadoras_Id,
	}

	err = tx.Select(&results, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchServiceToMoneyByIdTx(service TransportServiceAmount, tx *sqlx.Tx) (results []TransportServiceAmount, err error) {

	query := `
		SELECT id,
		 COALESCE(public.servicos_por_valor.parceiros_id, '0') as parceiros_id,
		 transportadoras_servicos_id,
		 valor,
		 frete,
		 he_ativo
		FROM public.servicos_por_valor
		WHERE  public.servicos_por_valor.transportadoras_servicos_id = $1
	`

	args := []interface{}{
		service.Transportadoras_Servicos_Id,
	}

	err = tx.Select(&results, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchServiceToKmByIdTx(service TransportServiceKM, tx *sqlx.Tx) (results []TransportServiceKM, err error) {
	query := `
	SELECT 
		id,
		COALESCE(public.servicos_por_km.parceiros_id, '0') as parceiros_id,
		transportadoras_servicos_id,
		km,
		frete,
		he_ativo
	FROM public.servicos_por_km	
	WHERE  public.servicos_por_km.transportadoras_servicos_id = $1
`

	args := []interface{}{
		service.Transportadoras_Servicos_Id,
	}

	err = tx.Select(&results, query, args...)

	if err != nil {
		return
	}

	return
}

func CreateServiceToKMTx(service TransportServiceKM, tx *sqlx.Tx) (id int64, err error) {
	query := `
		INSERT INTO public.servicos_por_km(
			transportadoras_servicos_id, km, frete, he_ativo)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	active := 0
	if service.He_Ativo {
		active = 1
	}

	args := []interface{}{
		service.Transportadoras_Servicos_Id,
		service.Km,
		service.Frete,
		active,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func CreateServiceToMoneyTx(transport TransportServiceAmount, tx *sqlx.Tx) (id int64, err error) {
	query := `
	INSERT INTO public.servicos_por_valor(transportadoras_servicos_id, valor, frete, he_ativo)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`
	active := 0
	if transport.He_Ativo {
		active = 1
	}

	args := []interface{}{
		transport.Transportadoras_Servicos_Id,
		transport.Valor,
		transport.Frete,
		active,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func AlterServiceToKMTx(service TransportServiceKM, tx *sqlx.Tx) (err error) {
	query := `
	UPDATE public.servicos_por_km
		SET
		km = $1,
		frete = $2,
		he_ativo = $3
		WHERE  public.servicos_por_km.id = $4
		`

	active := 0
	if service.He_Ativo {
		active = 1
	}

	args := []interface{}{
		service.Km,
		service.Frete,
		active,
		service.ID,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func AlterServiceToMoneyTx(service TransportServiceAmount, tx *sqlx.Tx) (err error) {
	query := `
	UPDATE public.servicos_por_valor
		SET
		valor = $1,
		frete = $2,
		he_ativo = $3
		WHERE  public.servicos_por_valor.id = $4
		`

	active := 0
	if service.He_Ativo {
		active = 1
	}

	args := []interface{}{
		service.Valor,
		service.Frete,
		active,
		service.ID,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}
