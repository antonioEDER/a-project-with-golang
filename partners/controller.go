package partners

import (
	"fmt"
	"io"
	"math/rand"
	"os"

	"github.com/api-qop-v2/address"
	"github.com/api-qop-v2/common"
	"github.com/api-qop-v2/config"
	"github.com/api-qop-v2/send_emails"

	"github.com/eucatur/go-toolbox/database"
	"github.com/eucatur/go-toolbox/env"
	"github.com/labstack/echo"
)

func SearchPartnersFromProximity() {
	return
}

func ExistPartner() {
	return
}

func SearchPartnerTypesActivesPublic() (ranges []RangeActivity, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	ranges, err = SearchPartnerTypesActivesPublicTx(tx)

	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func SearchPartnerForUsuario() {
	return
}

func SearchPartner(p Partners) (list []Partners, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	list, err = SearchPartnerTx(p, tx)

	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func SearchPartnersAll(p Partners) (list []Partners, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	list, err = SearchPartnersAllTx(p, tx)

	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func SearchPlans(p Partners) (listProduct []Plano_Produto, listSales []Plano_Vendas, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	listProduct, listSales, err = SearchPlansTx(p, tx)

	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func SearchPartnersForDescription() {
	return
}

func SearchPartnerTypesAll() {
	return
}

func CreatePartner() {
	return
}

func AlterHoursOfOperation(pessoasId string, hours Partners) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	err = AlterHoursOfOperationTx(pessoasId, hours, tx)

	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func AlterPartner() {
	return
}

func AlterPlanPartner() {
	return
}

func CreatePlanPartner() {
	return
}

func SearchPlansPartners() {
	return
}

func SearchFinancialSummary() {
	return
}

func SearchFinancialPartnerSummary(f Filter) (resumo []ResumoVendas, grafico []ResumoVendas, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
	endPartner := ""
	if f.Parceiros_Id != "" {
		endPartner = " AND pedidos.parceiros_id = " + f.Parceiros_Id + " "
	}

	resumo, err = SummaryPurchaseTx(f, endPartner, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	grafico, err = SummaryPurchaseAllToGraphicTx(f, endPartner, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func SearchADMFinancialSummary(f Filter) (resumo []ResumoVendas, resumosParceiros []ResumoVendas, qtdPedidoPorStatus []ResumoVendas, qtdProdutosPorParceiros []ResumoVendas, grafico []ResumoVendas, graficosParceiros []ResumoVendas, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
	endPartner := ""
	if f.Parceiros_Id != "" {
		endPartner = " AND pedidos.parceiros_id = " + f.Parceiros_Id + " "
	}

	resumo, err = SummaryPurchaseTx(f, endPartner, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	resumosParceiros, err = SummaryPurchaseAllPartnerTx(f, endPartner, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	qtdPedidoPorStatus, err = SummaryPurchaseStatusTx(f, endPartner, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	qtdProdutosPorParceiros, err = SummaryPurchaseFromPartnerTx(f, endPartner, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	grafico, err = SummaryPurchaseAllToGraphicTx(f, endPartner, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	graficosParceiros, err = SummaryPurchasePartnerToGraphicTx(f, endPartner, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func CreateRangeActivity(branch RangeActivity) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	id, err = CreateRangeTx(branch, tx)
	if err != nil {
		//  err = echo.NewHTTPError(422, "Não foi possivel gravar ramo atividade")
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func AlterRangeActivity(branch RangeActivity) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	err = AlterRangeActivityTx(branch, tx)
	if err != nil {
		err = echo.NewHTTPError(422, "Não foi possivel gravar ramo atividade")
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func SearchRangeActivity() (branch []RangeActivity, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	branch, err = SearchRangeActivityTx(tx)

	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func SearchIdUserFromPartnerWithEmployeeData(userIDFromEmployee string) (idUserFromPartner string, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	idUserFromPartner, err = SearchIdUserFromPartnerWithEmployeeDataTx(userIDFromEmployee, tx)

	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func SearchPartnerFromUserId(userID string) (partner []Partners, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	partner, err = SearchPartnerFromUserIdTx(userID, tx)

	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func SearchPartnerFromPersonId(personID string) (partner []Partners, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	partner, err = SearchPartnerFromPersonIdTx(personID, tx)

	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func SearchEmployees(partnerID string) (p []address.PersonWeb, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	p, err = SearchEmployeesTx(partnerID, tx)

	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func SearchPartnersAllForUser(p Partners) (list []Partners, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	list, err = SearchPartnersAllForUserTx(p, tx)

	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func SearchPartnerForUser(p Partners) (list []Partners, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	list, err = SearchPartnerForUserTx(p, tx)

	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func SearchInvoicePartnerADM(f Filter) (fatura []ResumoVendas, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()
	endPartner := ""
	if f.Parceiros_Id != "" {
		endPartner = " AND pedidos.parceiros_id = " + f.Parceiros_Id + " "
	}

	fatura, err = SearchInvoicePartnerADMTx(f, endPartner, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func SearchPlansPartnerForValue(f Filter) (plan []Plano_Vendas, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	plan, err = SearchPlansPartnerForValueTx(f, tx)

	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func SearchPlansPartnerForProduct(f Filter) (plan []Plano_Produto, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	plan, err = SearchPlansPartnerForProductTx(f, tx)

	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func CreateEmailInvoice(c echo.Context) (err error) {
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
	params.TextoDaOferta = c.FormValue("textoDaFatura")
	params.NomeParceiro = "Parceiro qop"

	filename := dir
	err = send_emails.SendInvoicePartner(params, filename)

	return
}

func CreatePotentialPartners(f Filter) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	id, err = CreatePotentialPartnersTx(f, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}
