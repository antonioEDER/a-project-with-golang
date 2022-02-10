package employees

import (
	gotoolboxtime "github.com/eucatur/go-toolbox/time"
)

type Employee struct {
	ID                  int64                   `query:"id" json:"id"`
	Parceiros_Id        int64                   `query:"parceiros_id" json:"parceiros_id"`
	Pessoas_Usuarios_Id int64                   `query:"pessoas_usuarios_id" json:"pessoas_usuarios_id"`
	He_Adm              bool                    `query:"he_adm" json:"he_adm"`
	He_Ativo            bool                    `query:"he_ativo" json:"he_ativo"`
	CREATED_AT          gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY          gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT           gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY           gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED             gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}
