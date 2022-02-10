package persons

import (
	gotoolboxtime "github.com/eucatur/go-toolbox/time"
)

const (
	// PackageName ...
	PackageName                        = "person"
	cachePrefix                        = PackageName
	envPersonCacheSeconds              = PackageName + "_cache_seconds"
	envPersonByCodeCacheSeconds        = PackageName + "_by_code_cache_seconds"
	RowPerPage                         = 100
	ErrorForeignKey                    = "foreign key constraint fails"
	MessageErrorForeignKey             = "Exclusão não permitida! Existem dados vinculados a esta pessoa."
	MessageErrorCNPJExists             = "CNPJ já existe no cadastro de pessoa."
	MessageErrorEmailWebExists         = "E-mail da Web já esta vinculado a outra pessoa."
	MessageErrorPersonAddressNotEqual  = "O ID de registro pessoaendereco, não é igual ao vinculado a conta."
	MessageErrorPhysicalPersonNotEqual = "O ID de registro pessoafisica, não é igual ao vinculado a conta."
	MessageErrorPhysicalPersonCPFEmpty = "O CPF de pessoa física deve ser informado."
	MessageErrorPhysicalPersonRGEmpty  = "O RG de pessoa física deve ser informado."
)

type Person struct {
	ID         int64                   `query:"id" json:"id"`
	Nome       string                  `query:"nome" json:"nome"`
	Tipo       string                  `query:"tipo" json:"tipo"`
	He_Ativo   bool                    `query:"he_ativo" json:"he_ativo"`
	CREATED_AT gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT  gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY  gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED    gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type PersonContact struct {
	ID         int64                   `query:"id" json:"id"`
	Pessoas_Id int64                   `query:"pessoas_id" json:"pessoas_id"`
	Tipo       string                  `query:"tipo" json:"tipo"`
	Contato    string                  `query:"contato" json:"contato"`
	Obs        string                  `query:"obs" json:"obs"`
	He_Ativo   bool                    `query:"he_ativo" json:"he_ativo"`
	CREATED_AT gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT  gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY  gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED    gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type PersonPhysical struct {
	ID              int64                   `query:"id" json:"id"`
	Pessoas_Id      int64                   `query:"pessoas_id" json:"pessoas_id"`
	Cpf             string                  `query:"cpf" json:"cpf"`
	Documento       string                  `query:"documento" json:"documento"`
	Tipo_Documento  string                  `query:"tipo_documento" json:"tipo_documento"`
	Data_Nascimento string                  `query:"data_nascimento" json:"data_nascimento"`
	He_Ativo        bool                    `query:"he_ativo" json:"he_ativo"`
	CREATED_AT      gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY      gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT       gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY       gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED         gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type PersonLegal struct {
	ID           int64                   `query:"id" json:"id"`
	Pessoas_Id   string                  `query:"pessoas_id" json:"pessoas_id"`
	Razao_Social string                  `query:"razao_social" json:"razao_social"`
	Fantasia     string                  `query:"fantasia" json:"fantasia"`
	Cnpj         string                  `query:"cnpj" json:"cnpj"`
	He_Ativo     bool                    `query:"he_ativo" json:"he_ativo"`
	CREATED_AT   gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY   gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT    gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY    gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED      gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type PersonStatus struct {
	ID         int64                   `query:"id" json:"id"`
	Pessoas_Id int64                   `query:"pessoas_id" json:"pessoas_id"`
	Descricao  string                  `query:"descricao" json:"descricao"`
	Obs        string                  `query:"obs" json:"obs"`
	He_Ativo   bool                    `query:"he_ativo" json:"he_ativo"`
	CREATED_AT gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT  gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY  gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED    gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type PersonUser struct {
	ID          int64                   `query:"id" json:"id"`
	Pessoas_Id  string                  `query:"pessoas_id" json:"pessoas_id"`
	Email       string                  `query:"email" json:"email"`
	Rede_Social string                  `query:"rede_social" json:"rede_social"`
	Foto        string                  `query:"foto" json:"foto"`
	He_Ativo    bool                    `query:"he_ativo" json:"he_ativo"`
	CREATED_AT  gotoolboxtime.Timestamp `query:"created_at" json:"created_at"`
	CREATED_BY  gotoolboxtime.Timestamp `query:"created_by" json:"created_by"`
	UPDATE_AT   gotoolboxtime.Timestamp `query:"updated_at" json:"updated_at"`
	UPDATE_BY   gotoolboxtime.Timestamp `query:"updated_by" json:"updated_by"`
	DELETED     gotoolboxtime.Timestamp `query:"deleted" json:"deleted"`
}

type PersonWeb struct {
	Id                    string `db:"id" json:"id" `
	Tipo                  string `db:"tipo" json:"tipo" `
	Nome                  string `db:"nome" json:"nome" `
	Data_Nascimento       string `db:"data_nascimento" json:"data_nascimento" `
	Cpf                   string `db:"cpf" json:"cpf" `
	Cnpj                  string `db:"cpf" json:"cnpj" `
	Email                 string `db:"email" json:"email" `
	Senha                 string `db:"senha" json:"senha" `
	Nome_Rede_Social      string `db:"rede_social" json:"nome_rede_social" `
	Uid                   string `db:"uid" json:"uid" `
	Celular               string `db:"celular" json:"celular" `
	Foto                  string `db:"foto" json:"foto" `
	Uf                    string `db:"uf" json:"uf" `
	Cidade                string `db:"cidade" json:"cidade" `
	Numero                string `db:"numero" json:"numero" `
	Bairro                string `db:"bairro" json:"bairro" `
	Cep                   string `db:"cep" json:"cep" `
	Logradouro            string `db:"logradouro" json:"logradouro" `
	Latitude              string `db:"latitude" json:"latitude" `
	Longitude             string `db:"longitude" json:"longitude" `
	Complemento           string `db:"complemento" json:"complemento" `
	He_Principal          string `db:"he_principal" json:"he_principal" `
	He_Principal_Parceiro string `db:"he_principal_parceiro" json:"he_principal_parceiro" `
	Area_Abrangencia      string `db:"area_abrangencia" json:"area_abrangencia" `
	Pessoas_Id            string `query:"pessoas_id" json:"pessoas_id"`
	Contato               string `query:"contato" json:"contato"`
	He_Ativo              bool   `query:"he_ativo" json:"he_ativo"`
}
