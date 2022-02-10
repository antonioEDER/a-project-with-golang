package users

const (
	MessageErrorValidateAccountByCode = "Código incorreto. Tente novamente!"
)

// LoginParams ...
type LoginParams struct {
	ID int64 `db:"id" json:"id"`

	Email string `json:"email" validate:"required" errmsg:"Informe o Email do usuário"`
	Senha string `json:"senha" validate:"required" errmsg:"Informe a Senha do usuário"`
	Tipo  string `json:"tipo" validate:"required" errmsg:"Informe a Tipo do usuário"`
	Token string `db:"token" json:"token" query:"token"`
}

type User struct {
	ID              int64  `db:"id" json:"id"`
	Nome            string `db:"nome" json:"nome" `
	Pessoas_Id      string `db:"pessoas_id" json:"pessoas_id"`
	Email           string `db:"email" json:"email"`
	Tipo            string `db:"tipo" json:"tipo"`
	Senha           string `db:"senha" json:"senha"`
	Rede_Social     string `db:"rede_social" json:"rede_social"`
	Foto            string `db:"foto" json:"foto"`
	Cod_Confirmacao string `db:"cod_confirmacao" json:"cod_confirmacao"`
	He_Ativo        bool   `db:"he_ativo" json:"he_ativo"`
	Uid             string `db:"uid" json:"uid"`
	Token           string `db:"token" json:"token" query:"token"`
	Token_Push_Web  string `db:"token_push_web" json:"token_push_web" query:"token_push_web"`
	Token_Push_App  string `db:"token_push_app" json:"token_push_app" query:"token_push_app"`
}

type PersonWeb struct {
	Id                    string `db:"id" json:"id" `
	Tipo                  string `db:"tipo" json:"tipo" `
	Tipo_Contato          string `db:"tipo_contato" json:"tipo_contato" `
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
	Tipos_Pessoas_Id      string `query:"tipos_pessoas_id" json:"tipos_pessoas_id"`
	He_Ativo              bool   `query:"he_ativo" json:"he_ativo"`
	He_Adm                bool   `query:"he_adm" json:"he_adm"`
}

type Contact struct {
	Name       string `db:"name" json:"name" `
	Phone      string `db:"phone" json:"phone" `
	Email      string `db:"email" json:"email" `
	Text       string `db:"text" json:"text" `
	Subject    string `db:"subject" json:"subject" `
	Partner_ID string `db:"partner_id" json:"partner_id" `
}
