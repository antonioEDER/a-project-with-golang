package ordereds

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/api-qop-v2/common"
	"github.com/api-qop-v2/config"
	"github.com/api-qop-v2/partners"
	"github.com/api-qop-v2/products"
	"github.com/api-qop-v2/send_emails"
	"github.com/api-qop-v2/users"
	"github.com/eucatur/go-toolbox/database"
	"github.com/eucatur/go-toolbox/env"
	"github.com/labstack/echo"
)

func SearchOrderToResendProductDigital(o Order, c echo.Context) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	orders, err := SearchOrderToResendProductDigitalTx(o, tx)
	if err != nil || len(orders) == 0 {
		err = fmt.Errorf("Pedido não encontrado.")
		return
	}

	var u users.User
	u.ID = orders[0].Pessoas_Usuarios_Id

	us, err := users.SearchUserFromIdUserTx(u, tx)
	if err != nil {
		return
	}

	var params send_emails.OrderCreate
	params.Status_Id = orders[0].Pedidos_Status_Id
	params.Pedidos_Id = o.Pedidos_Id
	params.Status_Descricao = orders[0].Pedidos_Status_Descricao
	params.Nome = "Cliente"
	params.Email = us[0].Email
	params.Pedido_Data = fmt.Sprint(orders[0].CREATED_AT)

	// go send_emails.SendAlterationStatusOrder(params)

	var itens send_emails.OrderItens

	for _, order := range orders {
		var p send_emails.OrderCreate
		p.Descricao = order.Descricao
		p.ID = order.ID
		p.Imagens_Diretorio = order.Imagens_Diretorio
		p.QrCode = "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"
		itens.Itens = append(itens.Itens, p) // Push

	}

	go send_emails.SendProductDigital(params, itens)

	err = tx.Commit()

	return
}

func SearchOrderSpecific(o Order) (orders []Order, userPerson []users.PersonWeb, partner []partners.Partners, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	var u users.User
	u.Pessoas_Id = o.Pessoas_Id
	us, err := users.SearchUserFromIdPersonTx(u, tx)
	if err != nil {
		return
	}

	var lg users.LoginParams
	lg.Tipo = "USUARIO"
	lg.Email = us[0].Email

	userPerson, err = users.LoginDBTx(lg, false, tx)
	if err != nil {
		return
	}

	o.Pessoas_Usuarios_Id = us[0].ID
	orders, err = SearchOrderSpecificActiveTx(o, tx)
	if err != nil || len(orders) == 0 {
		err = fmt.Errorf("Pedido não encontrado.")
		return
	}

	var p partners.Partners
	p.Parceiros_Id = orders[0].Parceiros_Id
	partner, err = partners.SearchPartnerForUserTx(p, tx)
	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func SearchOrderSpecificForParner(o Order) (orders []Order, userPerson []users.PersonWeb, partner []partners.Partners, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	orders, err = SearchOrderForPartnerTx(o, tx)
	if err != nil || len(orders) == 0 {
		err = fmt.Errorf("Pedido não encontrado.")
		return
	}

	var u users.User
	u.Pessoas_Id = orders[0].Pessoas_Id
	us, err := users.SearchUserFromIdPersonTx(u, tx)
	if err != nil {
		return
	}

	var lg users.LoginParams
	lg.Tipo = "USUARIO"
	lg.Email = us[0].Email
	userPerson, err = users.LoginDBTx(lg, false, tx)
	if err != nil {
		return
	}

	var p partners.Partners
	p.Parceiros_Id = orders[0].Parceiros_Id
	partner, err = partners.SearchPartnerForUserTx(p, tx)
	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func SearchOrder() {
	return
}

func SearchOrderDetailed(o ParamsProductComposite) (orders []ParamsProductComposite, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	orders, err = SearchProductCompositeByIdTx(o, tx)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	err = tx.Commit()

	return
}

func SearchOrderForPartner() {
	return
}

func SearchOrdersConfirmed() {
	return
}

func AlterOrdersAccepted(o Order) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
	pt, err := partners.SearchPartnerFromPersonIdTx(fmt.Sprint(o.Pessoas_Id), tx)
	if err != nil {
		return
	}

	parceirosId, err := strconv.ParseInt(fmt.Sprint(pt[0].Id), 10, 64)
	if err != nil {
		return
	}
	o.Parceiros_Id = parceirosId

	err = AlterOrdersAcceptedTx(o, tx)
	if err != nil {
		return
	}

	err = tx.Commit()

	return

}
func SearchOrdersForCSV(f Filter) (csv string, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
	pt, err := partners.SearchPartnerFromPersonIdTx(fmt.Sprint(f.Pessoas_Id), tx)
	if err != nil {
		return
	}
	f.Parceiros_Id = pt[0].Id

	filters := ` WHERE pedidos.parceiros_id = $2 AND enderecos.he_principal = 1  `

	if f.Tipos_Status_Ids != "" {
		filters += ` AND pedidos_status.id IN (` + f.Tipos_Status_Ids + `) `
	}
	if f.Tipos_Pedidos_Ids != "" {
		filters += ` AND pedidos_categorias.id IN (` + f.Tipos_Pedidos_Ids + `) `
	}

	if f.Data_Final == "" {
		filters += ` AND 
		CAST((SELECT pedidos.created_at at time zone 'UTC' at time zone $1) AS DATE) 
		=  
		CAST('` + f.Data_Inicio + `' AS DATE) `
	}

	if f.Data_Final != "" {
		filters += ` AND 
		CAST((SELECT pedidos.created_at at time zone 'UTC' at time zone $1) AS DATE) 
		BETWEEN  
		CAST('` + f.Data_Inicio + `' AS DATE) 
		AND 
		CAST('` + f.Data_Final + `' AS DATE) `
	}

	if f.Numero > 0 {
		filters += ` AND pedidos.id =` + fmt.Sprint(f.Numero) + ` `
	}

	visualizado := 0
	if f.Visualizado {
		visualizado = 1
	}
	if f.Todos == false {
		filters += ` AND pedidos.visualizado =` + fmt.Sprint(visualizado) + ` `
	}

	args := []interface{}{
		f.Time_Zone,
		f.Parceiros_Id,
	}
	orders, err := SearchOrdersForFiltersTx(filters, args, tx)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	csv = ""
	csv += "id,"
	csv += "parceiros_id,"
	csv += "pessoas_usuarios_id,"
	csv += "pedidos_categorias_id,"
	csv += "pedidos_status_id,"
	csv += "enderecos_id,"
	csv += "valor_total,"
	csv += "forma_pagamento,"
	csv += "troco,"
	csv += "servico_pagamento,"
	csv += "taxa_entrega,"
	csv += "tipo_veiculo,"
	csv += "placa_veiculo,"
	csv += "cor_veiculo,"
	csv += "entregador,"
	csv += "visualizado,"
	csv += "he_ativo,"
	csv += "prazo_entrega,"
	csv += "created_at,"
	csv += "pedidos_status_descricao,"
	csv += "pedidos_status_id,"
	csv += "pedidos_status_cor,"
	csv += "pedidos_categorias_descricao,"
	csv += "pedidos_categorias_id,"
	csv += "pedidos_status_ids,"
	csv += "pessoas_id,"
	csv += "pessoas_usuarios_id,"
	csv += "email,"
	csv += "pessoas_id,"
	csv += "nome,"
	csv += "cpf,"
	csv += "data_nascimento,"
	csv += "celular,"
	csv += "cep,"
	csv += "logradouro,"
	csv += "bairro,"
	csv += "complemento,"
	csv += "numero,"
	csv += "cidade,"
	csv += "uf,"
	csv += "\n"

	for _, pt := range orders {
		csv += fmt.Sprintf("%#v,", pt.ID)
		csv += fmt.Sprintf("%#v,", pt.Parceiros_Id)
		csv += fmt.Sprintf("%#v,", pt.Pessoas_Usuarios_Id)
		csv += fmt.Sprintf("%#v,", pt.Pedidos_Categorias_Id)
		csv += fmt.Sprintf("%#v,", pt.Pedidos_Status_Id)
		csv += fmt.Sprintf("%#v,", pt.Enderecos_Id)
		csv += fmt.Sprintf("%#v,", pt.Valor_Total)
		csv += fmt.Sprintf("%#v,", pt.Forma_Pagamento)
		csv += fmt.Sprintf("%#v,", pt.Troco)
		csv += fmt.Sprintf("%#v,", pt.Servico_Pagamento)
		csv += fmt.Sprintf("%#v,", pt.Taxa_Entrega)
		csv += fmt.Sprintf("%#v,", pt.Tipo_Veiculo)
		csv += fmt.Sprintf("%#v,", pt.Placa_Veiculo)
		csv += fmt.Sprintf("%#v,", pt.Cor_Veiculo)
		csv += fmt.Sprintf("%#v,", pt.Entregador)
		csv += fmt.Sprintf("%#v,", pt.Visualizado)
		csv += fmt.Sprintf("%#v,", pt.He_Ativo)
		csv += fmt.Sprintf("%#v,", pt.Prazo_Entrega)

		csv += fmt.Sprintf("%#v,", pt.CREATED_AT.String())

		csv += fmt.Sprintf("%#v,", pt.Pedidos_Status_Descricao)
		csv += fmt.Sprintf("%#v,", pt.Pedidos_Status_Id)
		csv += fmt.Sprintf("%#v,", pt.Pedidos_Status_Cor)
		csv += fmt.Sprintf("%#v,", pt.Pedidos_Categorias_Descricao)
		csv += fmt.Sprintf("%#v,", pt.Pedidos_Categorias_Id)
		csv += fmt.Sprintf("%#v,", pt.Pedidos_Status_Id)
		csv += fmt.Sprintf("%#v,", pt.Pessoas_Id)
		csv += fmt.Sprintf("%#v,", pt.Pessoas_Usuarios_Id)
		csv += fmt.Sprintf("%#v,", pt.Email)
		csv += fmt.Sprintf("%#v,", pt.Pessoas_Id)
		csv += fmt.Sprintf("%#v,", pt.Nome)
		csv += fmt.Sprintf("%#v,", pt.Cpf)
		csv += fmt.Sprintf("%#v,", pt.Data_Nascimento)
		csv += fmt.Sprintf("%#v,", pt.Celular)
		csv += fmt.Sprintf("%#v,", pt.Cep)
		csv += fmt.Sprintf("%#v,", pt.Logradouro)
		csv += fmt.Sprintf("%#v,", pt.Bairro)
		csv += fmt.Sprintf("%#v,", pt.Complemento)
		csv += fmt.Sprintf("%#v,", pt.Numero)
		csv += fmt.Sprintf("%#v,", pt.Cidade)
		csv += fmt.Sprintf("%#v,", pt.Uf)
		csv += "\n"
	}

	err = tx.Commit()

	return
}

func SearchOrdersForFilters(f Filter) (orders []Order, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
	pt, err := partners.SearchPartnerFromPersonIdTx(fmt.Sprint(f.Pessoas_Id), tx)
	if err != nil {
		return
	}
	f.Parceiros_Id = pt[0].Id

	filters := ` WHERE pedidos.parceiros_id = $2 AND enderecos.he_principal = 1  `

	if f.Tipos_Status_Ids != "" {
		filters += ` AND pedidos_status.id IN (` + f.Tipos_Status_Ids + `) `
	}
	if f.Tipos_Pedidos_Ids != "" {
		filters += ` AND pedidos_categorias.id IN (` + f.Tipos_Pedidos_Ids + `) `
	}

	if f.Data_Final == "" {
		filters += ` AND 
		CAST((SELECT pedidos.created_at at time zone 'UTC' at time zone $1) AS DATE) 
		=  
		CAST('` + f.Data_Inicio + `' AS DATE) `
	}

	if f.Data_Final != "" {
		filters += ` AND 
		CAST((SELECT pedidos.created_at at time zone 'UTC' at time zone $1) AS DATE) 
		BETWEEN  
		CAST('` + f.Data_Inicio + `' AS DATE) 
		AND 
		CAST('` + f.Data_Final + `' AS DATE) `
	}

	if f.Numero > 0 {
		filters += ` AND pedidos.id =` + fmt.Sprint(f.Numero) + ` `
	}

	visualizado := 0
	if f.Visualizado {
		visualizado = 1
	}
	if f.Todos == false {
		filters += ` AND pedidos.visualizado =` + fmt.Sprint(visualizado) + ` `
	}

	args := []interface{}{
		f.Time_Zone,
		f.Parceiros_Id,
	}
	orders, err = SearchOrdersForFiltersTx(filters, args, tx)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	err = tx.Commit()

	return
}

func CreateOrderForPayCash() {
	return
}

func CreateOrder(o Order, c echo.Context) (idOrder int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	var user users.User
	user.Pessoas_Id = o.Pessoas_Id
	us, err := users.SearchUserFromIdPersonTx(user, tx)
	if err != nil {
		return
	}
	o.Pessoas_Usuarios_Id = us[0].ID

	var pt partners.Partners
	pt.Parceiros_Id = o.Parceiros_Id
	partner, err := partners.SearchPartnerForUserTx(pt, tx)
	if err != nil {
		return
	}

	o.He_Servico_Pagamento = partner[0].He_Servico_Pagamento

	o.Pedidos_Status_Id = 1
	if o.Pedidos_Categorias_Id != 3 && o.Forma_Pagamento != "money" && o.Servico_Pagamento != "" && o.He_Servico_Pagamento == 1 {
		o.Pedidos_Status_Id = 11
	}

	idOrder, err = CreateOrderTx(o, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	for _, p := range o.Produtos {

		var paramProduct products.Product
		paramProduct.Parceiros_Id = o.Parceiros_Id
		paramProduct.ID = p.ID
		product, _ := products.SearchProductsForUSerFromIdTx(paramProduct, tx)
		if len(product) == 0 {
			if err != nil {
				err = fmt.Errorf("Produto não encontrado")
				tx.Rollback()
				return
			}
		}

		var paramItens OrderItens
		paramItens.Pedidos_Id = idOrder
		paramItens.Produtos_Id = p.ID
		paramItens.Vinculo_Para_Produto_Composto = p.Vinculo_Para_Produto_Composto

		amount := product[0].Valor
		if product[0].Valor_Promocao != "" && product[0].He_Promocao == 1 {
			amount = product[0].Valor_Promocao
		}

		paramItens.Valor = amount

		_, err = CreateListProductsFromOrderTx(idOrder, paramItens, tx)
		if err != nil {
			tx.Rollback()
			return
		}

		if p.Vinculo_Para_Produto_Composto != "" {
			for _, pc := range p.ProdutosCompostoAdd {
				_, err = CreateListProductsCompoundFromOrderTx(idOrder, o, p, pc, tx)
				if err != nil {
					tx.Rollback()
					return
				}
			}
		}
	}

	var lg users.LoginParams
	lg.Tipo = "USUARIO"
	lg.Email = us[0].Email

	person, err := users.LoginDBTx(lg, false, tx)
	if err != nil {
		return
	}
	now := time.Now()
	timeZoneDescription := strings.TrimSpace(c.Request().Header.Get("QOP-User-Time-Zone"))
	loc, _ := time.LoadLocation(timeZoneDescription)

	var params send_emails.OrderCreate
	params.Pedidos_Id = idOrder
	params.Email = us[0].Email
	params.Nome = person[0].Nome
	params.Pedido_Data = fmt.Sprint(now.In(loc).Format("02-01-2006 15:04:05"))

	if partner[0].Receber_Pedido_Por_Email == 1 {
		params.EmailParceiro = partner[0].Email
	}

	go send_emails.SendOrdersCreated(params)

	err = tx.Commit()

	return

}

func SearchOrderForUser(o Order) (orders []Order, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	var u users.User
	u.Pessoas_Id = o.Pessoas_Id
	us, err := users.SearchUserFromIdPersonTx(u, tx)
	if err != nil {
		return
	}

	fmt.Println("01 ===")

	var lg users.LoginParams
	lg.Tipo = "USUARIO"
	lg.Email = us[0].Email
	o.Pessoas_Usuarios_Id = us[0].ID

	orders, err = SearchOrderForUserTx(o, tx)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	fmt.Println("02 ===")

	err = tx.Commit()

	return
}

func SearchOrderTypes() {
	return
}

func SearchOrderStatus() {
	return
}

func SearchOrdersNews() {
	return
}

func AlterOrderTrackingCode() {
	return
}

func AlterOrderStatus(o Order) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
	pt, err := partners.SearchPartnerFromPersonIdTx(fmt.Sprint(o.Pessoas_Id), tx)
	if err != nil {
		return
	}

	err = AlterOrderStatusTx(o, pt[0], tx)
	if err != nil {
		return
	}

	orders, err := SearchOrderForPartnerTx(o, tx)
	if err != nil || len(orders) == 0 {
		err = fmt.Errorf("Pedido não encontrado.")
		return
	}

	err = tx.Commit()

	var params send_emails.OrderCreate
	params.Status_Id = orders[0].Pedidos_Status_Id
	params.Pedidos_Id = o.Pedidos_Id
	params.Status_Descricao = orders[0].Pedidos_Status_Descricao
	params.Nome = orders[0].Nome
	params.Email = orders[0].Email
	params.Pedido_Data = fmt.Sprint(orders[0].CREATED_AT)

	go send_emails.SendAlterationStatusOrder(params)

	return

}

func CancelOrder(order Order) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
	err = CancelOrderTx(order, tx)
	if err != nil {
		return
	}
	err = tx.Commit()

	return
}

func GenerateCsvOrders() {
	return
}

func CreateEmailBudget(o Order, c echo.Context) (err error) {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	// Destination
	PathFileTemp := env.MustString(common.PathFileTemp)
	dir := file.Filename + fmt.Sprint((rand.Float64() * 8))

	dst, err := os.Create(PathFileTemp + dir)
	if err != nil {
		return err
	}

	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	var params send_emails.SendBudget
	params.EmailCliente = c.FormValue("emailCliente")
	params.EmailLoja = c.FormValue("emailLoja")
	params.TextoDaOferta = c.FormValue("textoDaOferta")
	params.NomeParceiro = c.FormValue("nomeParceiro")

	filename := dir
	err = send_emails.SendEmailBudget(params, filename)

	return
}

func SearchCategoriesOrder() (orders []OrderStatus, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	orders, err = SearchCategoriesOrderTx(tx)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	err = tx.Commit()

	return
}

func SearchStatusOrder() (orders []OrderCategory, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	orders, err = SearchStatusOrderTx(tx)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	err = tx.Commit()

	return
}

func SearchOrdersForId(o Order) (orders []Order, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	orders, err = SearchOrdersForIConditionsTx(o, tx)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	err = tx.Commit()

	return
}
