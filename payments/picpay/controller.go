package picpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/api-qop-v2/common"
	"github.com/api-qop-v2/config"
	"github.com/api-qop-v2/ordereds"
	"github.com/api-qop-v2/partners"
	"github.com/eucatur/go-toolbox/database"
	"github.com/eucatur/go-toolbox/env"
)

func CheckoutPicPay(o ordereds.Order, p partners.Partners) (result Response, err error) {
	urlFront := env.MustString(common.UrlFront)

	nome := strings.Split(o.Nome, " ")
	firstname := nome[0]
	lastname := nome[1]

	var pay = Payment{}
	pay.Referenceid = "PD-" + fmt.Sprint(o.Pedidos_Id)
	pay.Callbackurl = urlFront
	pay.Returnurl = "https://qop.net.br/sucesso/" + fmt.Sprint(o.Pedidos_Id) + "?picpay=true"
	pay.Value = o.Valor_Total
	pay.Channel = "qop"
	pay.Purchasemode = "in-store"

	t := time.Now()
	newT := t.Add(time.Hour * 10).Format(time.RFC3339)
	pay.Expiresat = fmt.Sprint(newT)

	pay.Buyer.Firstname = firstname
	pay.Buyer.Lastname = lastname
	pay.Buyer.Document = o.Cpf
	pay.Buyer.Email = o.Email
	pay.Buyer.Phone = o.Telefone

	jsonData, err := json.Marshal(pay)
	if err != nil {
		return
	}

	UrlApiPicpay := env.MustString(common.UrlApiPicpay)

	resp, err := http.NewRequest("POST", UrlApiPicpay+"/payments", bytes.NewBuffer(jsonData))
	resp.Header.Set("Content-Type", "application/json")
	resp.Header.Set("x-picpay-token", p.Token_Pic_Pay)
	if err != nil {
		return
	}

	respCli, err := http.DefaultClient.Do(resp)
	if err != nil {
		return
	}

	body, readErr := ioutil.ReadAll(respCli.Body)
	if readErr != nil {
		err = fmt.Errorf("Errou ao iniciar Body")
		return
	}

	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		err = fmt.Errorf("Errou ao iniciar conversão para JSON")
		return
	}

	if result.Message != "" {
		err = fmt.Errorf("Erro sem mensage " + result.Message)
		return
	}

	return
}

func CancelPaymentPicPay(o ordereds.Order, p partners.Partners) (err error) {

	UrlApiPicpay := env.MustString(common.UrlApiPicpay)

	resp, err := http.NewRequest("POST", UrlApiPicpay+"/payments/PD-"+fmt.Sprint(o.Pedidos_Id)+"/cancellations", nil)
	resp.Header.Set("Content-Type", "application/json")
	resp.Header.Set("x-picpay-token", p.Token_Pic_Pay)
	if err != nil {
		return
	}

	respCli, err := http.DefaultClient.Do(resp)
	if err != nil {
		return
	}

	body, readErr := ioutil.ReadAll(respCli.Body)
	if readErr != nil {
		err = fmt.Errorf("Errou ao iniciar Body")
		return
	}

	var result Response
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		err = fmt.Errorf("Errou ao iniciar conversão para JSON")
		return
	}

	if result.Message != "" {
		err = fmt.Errorf("Erro sem mensage " + result.Message)
		return
	}

	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	err = CancelPaymentTx(o, tx)
	if err != nil {
		return
	}

	err = ordereds.CancelOrderTx(o, tx)
	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func SearchPaymentPicPay() {
	return
}
