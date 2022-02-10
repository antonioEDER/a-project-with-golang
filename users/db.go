package users

import (
	"github.com/jmoiron/sqlx"
)

func SearchUserMigrateDBTx(param LoginParams, tx *sqlx.Tx) (users []PersonWeb, err error) {
	query := `
	SELECT 
		id
	FROM public.usuarios_migracao
	WHERE public.usuarios_migracao.pessoas_usuarios_id = $1`

	args := []interface{}{
		param.ID,
	}

	err = tx.Select(&users, query, args...)
	return
}

func CheckPassWordDBTx(param LoginParams, tx *sqlx.Tx) (users []PersonWeb, err error) {
	query := `
	SELECT
		public.pessoas_usuarios.id,
		public.pessoas_usuarios.email,
		public.pessoas_usuarios.senha
	FROM
		public.pessoas_usuarios

	WHERE public.pessoas_usuarios.email = $1

	`
	args := []interface{}{
		param.Email,
	}

	err = tx.Select(&users, query, args...)
	return
}

func LoginDBTx(param LoginParams, validateByPassword bool, tx *sqlx.Tx) (users []PersonWeb, err error) {
	query := `
	SELECT
		public.pessoas_usuarios.id,
		public.pessoas_usuarios.email,
		public.pessoas_usuarios.he_adm,

		public.pessoas.id as pessoas_id,
		public.pessoas.nome,
		
     	COALESCE(public.pessoas_juridicas.tipos_pessoas_id, public.pessoas_fisicas.tipos_pessoas_id) as tipos_pessoas_id,
		
		COALESCE(public.pessoas_fisicas.cpf, '') as cpf,
		COALESCE(public.pessoas_fisicas.data_nascimento, now()) as data_nascimento,
		
		public.pessoas_contatos.tipo as tipo_contato,
		public.pessoas_contatos.contato as celular,

		COALESCE(public.enderecos.cep, '') as cep,
		COALESCE(public.enderecos.logradouro, '') as logradouro,
		COALESCE(public.enderecos.bairro, '') as bairro,
		COALESCE(public.enderecos.complemento, '') as complemento,
		COALESCE(public.enderecos.numero, '') as numero,
		COALESCE(public.cidades.nome, '') as cidade,
		COALESCE(public.estados.uf, '') as uf

	FROM
		public.pessoas
			LEFT JOIN   public.pessoas_fisicas ON  public.pessoas_fisicas.pessoas_id = public.pessoas.id
			LEFT JOIN   public.pessoas_juridicas ON  public.pessoas_juridicas.pessoas_id = public.pessoas.id
			
			LEFT JOIN   public.pessoas_usuarios ON  public.pessoas_usuarios.pessoas_id = public.pessoas.id
			LEFT JOIN   public.pessoas_contatos ON  public.pessoas_contatos.pessoas_id = public.pessoas.id
			LEFT JOIN   public.enderecos ON  public.enderecos.pessoas_id = public.pessoas.id
			LEFT JOIN   public.cidades  on public.cidades.id = public.enderecos.cidades_id
			LEFT JOIN   public.estados  on public.estados.id = public.cidades.estados_id
	WHERE public.pessoas_usuarios.email = $1 LIMIT 1`
	// AND public.tipos_pessoas.tipo = $2
	args := []interface{}{
		param.Email,
		// param.Tipo,
	}

	err = tx.Select(&users, query, args...)
	return
}

func searchCodeValidateTx(user User, tx *sqlx.Tx) (confirmedUser []User, err error) {

	query := `
	SELECT pessoas_usuarios.id,
	pessoas_usuarios.email,
	pessoas_usuarios.pessoas_id,
	pessoas_usuarios.senha,
	pessoas_usuarios.he_ativo,
	public.tipos_pessoas.tipo,

	COALESCE( cod_confirmacao, '0') as cod_confirmacao
		FROM public.pessoas_usuarios
		LEFT JOIN public.pessoas ON  public.pessoas_usuarios.pessoas_id = public.pessoas.id
		LEFT JOIN public.pessoas_fisicas ON  public.pessoas_fisicas.pessoas_id = public.pessoas.id
		LEFT JOIN public.tipos_pessoas ON  public.tipos_pessoas.id = public.pessoas_fisicas.tipos_pessoas_id
		LEFT JOIN public.pessoas_contatos ON  public.pessoas_contatos.pessoas_id = public.pessoas.id
		WHERE pessoas_usuarios.email = $1
		AND pessoas_usuarios.cod_confirmacao = $2;
	`

	args := []interface{}{
		user.Email,
		user.Cod_Confirmacao,
	}

	err = tx.Select(&confirmedUser, query, args...)

	if err != nil {
		return
	}

	return
}

func activeUserTx(idPerson string, tx *sqlx.Tx) (err error) {

	queryActivePerson := `
		UPDATE public.pessoas
			SET he_ativo = $1
		WHERE id = $2 ;
		`

	queryActiveUser := `
	UPDATE public.pessoas_usuarios
		SET he_ativo = $1
	WHERE pessoas_id = $2 ;
	`

	queryActiveFisical := `
	UPDATE public.pessoas_fisicas
		SET he_ativo = $1
	WHERE pessoas_id = $2 ;
	`

	args := []interface{}{
		1,
		idPerson,
	}

	_, err = tx.Exec(queryActivePerson, args...)
	if err != nil {
		return err
	}

	_, err = tx.Exec(queryActiveUser, args...)
	if err != nil {
		return err
	}

	_, err = tx.Exec(queryActiveFisical, args...)
	if err != nil {
		return err
	}

	return
}

func NewPasswordTx(user User, tx *sqlx.Tx) (err error) {

	query := `
		UPDATE public.pessoas_usuarios
			SET senha = $1
		WHERE id = $2 ;
		`

	args := []interface{}{
		user.Senha,
		user.ID,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}

	return
}

// func SearchPersonsDBTx(tx *sqlx.Tx) (persons []Person, err error) {
// 	persons = []Person{}

// 	query := `SELECT
// 					id,
// 					nome,
// 					tipo,
// 					he_ativo,
// 					created_at,
// 					created_by,
// 					updated_at,
// 					updated_by,
// 					deleted
// 				FROM public.pessoas;`

// 	err = tx.Select(&persons, query)

// 	return
// }

func SearchUserFromIdUserTx(u User, tx *sqlx.Tx) (user []User, err error) {

	query := `SELECT 
		id,
		pessoas_id,
		email,
		senha,
		COALESCE(uid, '') as uid,

		COALESCE(rede_social, '') as rede_social,
		COALESCE(foto, '') as foto,
		COALESCE(token_push_web, '') as token_push_web,
	    COALESCE(token_push_app, '') as token_push_app,
		COALESCE(cod_confirmacao, '') as cod_confirmacao, 
		he_ativo 
		FROM public.pessoas_usuarios 
		WHERE id = $1;`

	args := []interface{}{
		u.ID,
	}

	err = tx.Select(&user, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchUserFromIdPersonTx(u User, tx *sqlx.Tx) (user []User, err error) {

	query := `SELECT 
		id,
		pessoas_id,
		email,
		senha,
		COALESCE(uid, '') as uid,
		COALESCE(rede_social, '') as rede_social,
		COALESCE(foto, '') as foto,
		COALESCE(token_push_web, '') as token_push_web,
	    COALESCE(token_push_app, '') as token_push_app,
		COALESCE(cod_confirmacao, '') as cod_confirmacao, 
		he_ativo 
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

func searchUserTx(user User, tx *sqlx.Tx) (confirmedUser []User, err error) {

	query := `SELECT 
		public.pessoas_usuarios.id,
		public.pessoas_usuarios.email,
		public.pessoas_usuarios.pessoas_id,
		public.pessoas.nome, 
		public.pessoas_usuarios.senha, 
		public.pessoas_usuarios.he_ativo,
		COALESCE(public.pessoas_usuarios.cod_confirmacao, '0') as cod_confirmacao
		FROM public.pessoas_usuarios
		JOIN public.pessoas on public.pessoas.id = public.pessoas_usuarios.pessoas_id
	WHERE email = $1;
	`

	args := []interface{}{
		user.Email,
	}

	err = tx.Select(&confirmedUser, query, args...)

	if err != nil {
		return
	}

	return
}

func SearchUserAllDataFromIdPersonTx(param User, tx *sqlx.Tx) (users []PersonWeb, err error) {
	query := `
	SELECT
		public.pessoas_usuarios.id,
		public.pessoas_usuarios.email,

		public.pessoas.id as pessoas_id,
		public.pessoas.nome,
		public.tipos_pessoas.id,
		
		COALESCE(public.pessoas_fisicas.cpf, '') as cpf,
		COALESCE(public.pessoas_fisicas.data_nascimento, now()) as data_nascimento,
		
		public.pessoas_contatos.tipo as tipo_contato,
		public.pessoas_contatos.contato as celular,

		public.enderecos.cep, 
		public.enderecos.logradouro, 
		public.enderecos.bairro,
		COALESCE(public.enderecos.complemento, '') as complemento, 
		public.enderecos.numero,
		public.cidades.nome as cidade, 
		public.estados.uf

	FROM
		public.pessoas
			LEFT JOIN   public.pessoas_fisicas ON  public.pessoas_fisicas.pessoas_id = public.pessoas.id
			LEFT JOIN   public.tipos_pessoas ON  public.tipos_pessoas.id = public.pessoas_fisicas.tipos_pessoas_id
			LEFT JOIN   public.pessoas_usuarios ON  public.pessoas_usuarios.pessoas_id = public.pessoas.id
			LEFT JOIN   public.pessoas_contatos ON  public.pessoas_contatos.pessoas_id = public.pessoas.id
			LEFT JOIN   public.enderecos ON  public.enderecos.pessoas_id = public.pessoas.id
			LEFT JOIN   public.cidades  on public.cidades.id = public.enderecos.cidades_id
			LEFT JOIN   public.estados  on public.estados.id = public.cidades.estados_id
		WHERE public.pessoas.id = $1
	`

	args := []interface{}{
		param.Pessoas_Id,
	}

	err = tx.Select(&users, query, args...)
	return
}

func CreateLeadsForOffersTx(user User, tx *sqlx.Tx) (id int64, err error) {

	query := `
		INSERT INTO public.ofertas_por_email(
		email)
		VALUES ($1)
		RETURNING id;
	`

	args := []interface{}{
		user.Email,
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
