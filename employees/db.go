package employees

import (
	"errors"

	"github.com/api-qop-v2/address"
	"github.com/jmoiron/sqlx"
)

func CreateEmployeesTx(idUser int64, person address.PersonWeb, tx *sqlx.Tx) (id int64, err error) {
	query := `
		INSERT INTO public.parceiros_funcionarios(parceiros_id, pessoas_usuarios_id, he_adm, he_ativo)
			VALUES ($1, $2, $3, $4)
		RETURNING id;
		`

	heAdm := 1
	if !person.He_Adm {
		heAdm = 0
	}
	args := []interface{}{
		person.Parceiros_Id,
		idUser,
		heAdm,
		1,
	}

	err = tx.Get(&id, query, args...)
	if err != nil {
		return
	}

	return id, err
}

func AlterEmployeesTx(person address.PersonWeb, tx *sqlx.Tx) (err error) {
	query := `
	UPDATE public.parceiros_funcionarios
	SET 
	he_adm=$1, 
	he_ativo=$2
	WHERE id = $3;
		`

	heAdm := 1
	if !person.He_Adm {
		heAdm = 0
	}
	heAtivo := 1
	if !person.He_Ativo {
		heAtivo = 0
	}

	args := []interface{}{
		heAdm,
		heAtivo,
		person.Id,
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func findByFilterDBTx(filter Filter, tx *sqlx.Tx) (user string, err error) {

	if filter.ID > 0 && tx != nil {
		user = "Seens we have a user :)"
	} else {
		err = errors.New("There's no user")
	}

	return
}
