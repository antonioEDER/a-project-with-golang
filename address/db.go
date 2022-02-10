package address

import (
	"strconv"

	"github.com/jmoiron/sqlx"
)

func SearchCoordenatesFromTwoUsersTx(userIdOne int64, userIdTwo int64, tx *sqlx.Tx) (geoOrigin string, geoDestiny string, err error) {
	var address []Address
	query := `SELECT 
				enderecos.id, 
				COALESCE(enderecos.latitude, '') as latitude,
				COALESCE(enderecos.longitude, '') as longitude
			FROM enderecos 
			WHERE enderecos.id in ( $1 , $2 )`

	args := []interface{}{
		userIdOne,
		userIdTwo,
	}
	err = tx.Select(&address, query, args...)
	if err != nil {
		return
	}

	geoOrigin = address[0].Latitude + "," + address[0].Longitude
	geoDestiny = address[1].Latitude + "," + address[1].Longitude

	return
}

func searchCityTx(uf string, cidade string, tx *sqlx.Tx) (cities []City, err error) {
	states := []State{}
	queryStates := `
		SELECT 
			id,
			COALESCE(paises_id, 0) as paises_id,
			COALESCE(nome, '') as nome,
			COALESCE(ddd, '') as ddd,
			COALESCE(uf, '') as  uf
		FROM public.estados
		WHERE uf = $1;`

	queryCity := `
		SELECT 
			id, 
			COALESCE(estados_id, 0) as estados_id,
			COALESCE(nome, '') as nome
		FROM public.cidades
		WHERE nome = $1 and estados_id = $2;
	`

	argsStates := []interface{}{
		uf,
	}
	err = tx.Select(&states, queryStates, argsStates...)

	argsCity := []interface{}{
		cidade,
		states[0].Id,
	}
	err = tx.Select(&cities, queryCity, argsCity...)

	return
}

func SearchAddresTx(personId string, tx *sqlx.Tx) (address []AddressWeb, err error) {

	query := `
		SELECT  
			public.enderecos.id, 
			COALESCE(public.enderecos.he_principal, 0) as he_principal,
			COALESCE(public.enderecos.logradouro, '') as logradouro,
			COALESCE(public.enderecos.numero, '') as numero,
			COALESCE(public.enderecos.bairro, '') as bairro,
			COALESCE(public.enderecos.complemento, '') as complemento,
			COALESCE(public.enderecos.latitude, '') as latitude,
			COALESCE(public.enderecos.longitude, '') as longitude,
			COALESCE(public.enderecos.cep, '') as cep,
			COALESCE(public.cidades.nome, '') as cidade,
			COALESCE(public.estados.nome, '') as estado,
			COALESCE(public.estados.uf, '') as uf
			FROM public.enderecos
			LEFT JOIN public.cidades  on public.cidades.id = public.enderecos.cidades_id
			LEFT JOIN public.estados  on public.estados.id = public.cidades.estados_id
			
		WHERE public.enderecos.pessoas_id = $1;
	`

	args := []interface{}{
		personId,
	}
	err = tx.Select(&address, query, args...)

	return
}

func CreateAddressTx(address PersonWeb, tx *sqlx.Tx) (id int64, err error) {
	cities, err := SearchCity(address.Uf, address.Cidade, tx)
	if err != nil {
		return
	}

	personsID := address.Pessoas_Id

	queryAll := `
	  UPDATE public.enderecos
	  SET
			he_principal = $2
		WHERE pessoas_id = $1
	`

	argsAll := []interface{}{
		personsID,
		0,
	}

	_, err = tx.Exec(queryAll, argsAll...)
	if err != nil {
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
			he_principal_parceiro, 
			created_at, 
			created_by
		)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id;
	`

	args := []interface{}{
		cities[0].Id,
		personsID,
		address.Cep,
		address.Logradouro,
		address.Bairro,
		address.Complemento,
		address.Numero,
		address.Latitude,
		address.Longitude,
		address.Area_Abrangencia,
		1,
		address.He_Principal_Parceiro,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func AlterAddressTx(address PersonWeb, tx *sqlx.Tx) (err error) {
	cities, err := SearchCity(address.Uf, address.Cidade, tx)
	if err != nil {
		return
	}

	personsID := address.Pessoas_Id

	addressID, err := strconv.ParseInt(address.Enderecos_Id, 10, 64)
	if err != nil {
		return
	}

	queryAll := `
	  UPDATE public.enderecos
	  SET
			he_principal = $2
		WHERE pessoas_id = $1
	`

	argsAll := []interface{}{
		personsID,
		0,
	}

	_, err = tx.Exec(queryAll, argsAll...)
	if err != nil {
		return
	}

	query := `
	  UPDATE public.enderecos
	  SET
			cidades_id = $1, 
			pessoas_id = $2, 
			cep = $3, 
			logradouro = $4, 
			bairro = $5, 
			complemento = $6, 
			numero = $7, 
			latitude = $8, 
			longitude = $9, 
			area_abrangencia = $10, 
			he_principal = $11, 
			he_principal_parceiro = $12
		WHERE id = $13
	`

	args := []interface{}{
		cities[0].Id,
		personsID,
		address.Cep,
		address.Logradouro,
		address.Bairro,
		address.Complemento,
		address.Numero,
		address.Latitude,
		address.Longitude,
		address.Area_Abrangencia,
		1,
		address.He_Principal_Parceiro,
		addressID,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}
