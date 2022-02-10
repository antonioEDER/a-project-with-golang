package pagseguro

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	xj "github.com/basgys/goxml2json"

	"github.com/api-qop-v2/common"
	"github.com/api-qop-v2/config"
	"github.com/api-qop-v2/ordereds"
	"github.com/api-qop-v2/partners"
	"github.com/api-qop-v2/products"

	"github.com/api-qop-v2/users"
	"github.com/eucatur/go-toolbox/database"
	"github.com/eucatur/go-toolbox/env"
)

func PaymentPagSeguro() {
	return
}

func GenerateSessionPagSeguro(o ordereds.Order, p partners.Partners) (idSession string, err error) {

	UrlApiPagSeguro := env.MustString(common.UrlApiPagSeguro)

	params := "?email=" + p.Email_Pag_Seguro + "&token=" + p.Token_Pag_Seguro

	resp, err := http.NewRequest("POST", UrlApiPagSeguro+"/sessions"+params, nil)
	resp.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return
	}

	respCli, err := http.DefaultClient.Do(resp)
	if err != nil {
		return
	}

	jsonBody, err := xj.Convert(respCli.Body)
	if err != nil {
		err = fmt.Errorf("Errou ao iniciar Body")
		return
	}

	var result ResultSession
	jsonErr := json.Unmarshal([]byte(jsonBody.String()), &result)
	if jsonErr != nil {
		err = fmt.Errorf("Errou ao iniciar conversão para JSON")
		return
	}

	if result.ID == "" {
		err = fmt.Errorf("Erro ao gerar Sessão")
		return
	}

	idSession = result.ID

	return
}

func CheckoutPagSeguro(o ordereds.Order, p partners.Partners) (codeCheckout string, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	var user users.User
	user.Pessoas_Id = o.Pessoas_Id
	us, err := users.SearchUserAllDataFromIdPersonTx(user, tx)
	if err != nil {
		return
	}

	UrlApiPagSeguro := env.MustString(common.UrlApiPagSeguro)
	params := "?email=" + p.Email_Pag_Seguro + "&token=" + p.Token_Pag_Seguro

	data := url.Values{}
	data.Set("currency", "BRL")
	data.Set("reference", "PD-"+fmt.Sprint(o.Pedidos_Id))
	data.Set("senderName", o.NomeCartao)
	data.Set("senderEmail", us[0].Email)
	data.Set("senderHash", o.HashCard)
	data.Set("shippingAddressRequired", "false")
	data.Set("billingAddressStreet", us[0].Logradouro)
	data.Set("billingAddressNumber", us[0].Numero)
	data.Set("billingAddressComplement", us[0].Complemento)
	data.Set("billingAddressDistrict", us[0].Bairro)
	data.Set("billingAddressPostalCode", us[0].Cep)
	data.Set("billingAddressCity", us[0].Cidade)
	data.Set("billingAddressState", us[0].Uf)
	data.Set("billingAddressCountry", "BRA")
	data.Set("creditCardToken", o.TokenCard)
	data.Set("installmentQuantity", fmt.Sprint(o.QtdParcelas))
	data.Set("installmentValue", o.ValorParcelas)
	data.Set("paymentMode", "default")
	data.Set("paymentMethod", "creditCard")
	data.Set("receiverEmail", p.Email_Pag_Seguro)
	data.Set("notificationURL", "https://qop.net.br/")
	data.Set("creditCardHolderName", o.NomeCartao)
	data.Set("creditCardHolderCPF", o.CpfCartao)
	data.Set("creditCardHolderBirthDate", o.NascimentoCartao)

	reg := regexp.MustCompile("[^0-9]")
	telefoneComprador := reg.ReplaceAllString(o.TelefoneComprador, "")
	data.Set("creditCardHolderAreaCode", telefoneComprador[0:2])

	telefoneCompradorNove := telefoneComprador[3:len(telefoneComprador)]
	if len(telefoneCompradorNove) < 9 {
		telefoneCompradorNove = "9" + telefoneCompradorNove
	}
	data.Set("creditCardHolderPhone", telefoneCompradorNove)

	telefoneUser := reg.ReplaceAllString(us[0].Celular, "")
	data.Set("senderAreaCode", telefoneUser[0:2])
	data.Set("senderPhone", telefoneUser[3:len(telefoneUser)])

	cpf := reg.ReplaceAllString(us[0].Cpf, "")
	data.Set("senderCPF", cpf)

	if o.Taxa_Entrega != 0.00 {
		contar := fmt.Sprint(len(o.Produtos) + 1)
		data.Set("itemId"+contar, "TX-"+fmt.Sprint(o.Pedidos_Id))
		data.Set("itemDescription"+contar, "Taxa de entrega")
		data.Set("itemAmount"+contar, fmt.Sprint(o.Taxa_Entrega))
		data.Set("itemQuantity"+contar, "1")
	}

	for i, produto := range o.Produtos {
		var paramProduct products.Product
		paramProduct.Parceiros_Id = o.Parceiros_Id
		paramProduct.ID = produto.ID
		product, _ := products.SearchProductsForUSerFromIdTx(paramProduct, tx)
		if len(product) == 0 {
			if err != nil {
				err = fmt.Errorf("Produto não encontrado")
				tx.Rollback()
				return
			}
		}
		contar := fmt.Sprint(i + 1)
		data.Set("itemId"+contar, "PD-"+fmt.Sprint(o.Pedidos_Id))
		data.Set("itemDescription"+contar, product[0].Descricao)
		if product[0].He_Promocao == 1 {
			data.Set("itemAmount"+contar, product[0].Valor_Promocao)
		} else {
			data.Set("itemAmount"+contar, product[0].Valor)
		}
		data.Set("itemQuantity"+contar, "1")
	}

	resp, err := http.NewRequest("POST", UrlApiPagSeguro+"/transactions"+params, strings.NewReader(data.Encode()))
	resp.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return
	}

	respCli, err := http.DefaultClient.Do(resp)
	if err != nil {
		return
	}

	userID, err := strconv.ParseInt(us[0].Id, 10, 64)
	if err != nil {
		return
	}
	o.Pessoas_Usuarios_Id = userID

	if respCli.StatusCode == 400 || respCli.StatusCode == 401 {
		err = fmt.Errorf("Problema com PagSeguro", respCli.StatusCode)
		return
	}

	jsonBody, err := xj.Convert(respCli.Body)
	if err != nil {
		err = fmt.Errorf("Errou ao iniciar Body")
		return
	}

	var result ResultTransaction
	_ = json.Unmarshal([]byte(jsonBody.String()), &result)

	if result.Code == "" {
		err = fmt.Errorf("Erro ao gerar Sessão")
		return
	}

	codeCheckout = result.Code

	return
}

func AlterStatusPagSeguro() {
	return
}

func SearchPaymentPagSeguro() {
	return
}
