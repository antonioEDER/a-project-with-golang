package pix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/api-qop-v2/common"
	"github.com/api-qop-v2/config"
	"github.com/api-qop-v2/ordereds"
	"github.com/api-qop-v2/partners"
	"github.com/api-qop-v2/users"
	"github.com/eucatur/go-toolbox/database"
	"github.com/eucatur/go-toolbox/env"
)

func GenerateQrCode(o ordereds.Order, p partners.Partners, user users.PersonWeb) (result Response, err error) {
	UrlApiPixManual := env.MustString(common.UrlApiPixManual)

	var pay = Payment{}
	pay.Dados.Chave_Pix = p.Chave_Pix
	pay.Dados.Cidade = user.Cidade
	pay.Dados.Fantasia = p.Fantasia
	pay.Dados.Pedidos_Id = o.Pedidos_Id
	pay.Dados.Valor_Total = o.Valor_Total

	jsonData, err := json.Marshal(pay)
	if err != nil {
		return
	}

	resp, err := http.NewRequest("GET", UrlApiPixManual, bytes.NewBuffer(jsonData))
	resp.Header.Set("Content-Type", "application/json")
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
	fmt.Println("body", jsonErr)

	if jsonErr != nil {
		err = fmt.Errorf("Errou ao iniciar conversão para JSON")
		return
	}

	return
}

func CancelPaymentPix(o ordereds.Order) (err error) {

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
