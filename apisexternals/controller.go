package apisexternals

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	xj "github.com/basgys/goxml2json"
)

func DistanceBetweenPointers(pointerOn string, pointerTwo string) (distance int, err error) {
	url := "https://maps.googleapis.com/maps/api/distancematrix/json?key=AIzaSyA-Y_Y8nHX_neFDGhA9VDKUfx0-YQpb9HA&origins=" + pointerOn + "&destinations=" + pointerTwo
	resp, err := http.NewRequest("GET", url, nil)
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
		return
	}
	result := Pointers{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		return
	}

	if result.Rows[0].Elements[0].Status != "OK" {
		err = fmt.Errorf("Errou no google Maps")
		return
	}

	distance = result.Rows[0].Elements[0].Distance.Value / 1000

	return
}

func CalculateServiceCorreios(d DataFromCorreios) (freight string, deadline string, err error) {
	url := "http://ws.correios.com.br/calculador/CalcPrecoPrazo.aspx?"
	url += "nCdEmpresa=08082650&"
	url += "sDsSenha=564321&"
	url += "sCepOrigem=" + d.Cep_Origem + "&"
	url += "sCepDestino=" + d.Cep_Destino + "&"
	url += "nVlPeso=" + strings.Replace(d.Peso, ".", ",", -1) + "&"
	url += "nCdFormato=1&"
	url += "nVlComprimento=" + d.Comprimento + "&"
	url += "nVlAltura=" + d.Altura + "&"
	url += "nVlLargura=" + d.Largura + "&"
	url += "sCdMaoPropria=n&"
	url += "nVlValorDeclarado=0&"
	url += "sCdAvisoRecebimento=n&"
	url += "nCdServico=" + d.Tipo_Entrega + "&"
	url += "StrRetorno=xml&"
	url += "nIndicaCalculo=3"

	resp, err := http.NewRequest("GET", url, nil)
	resp.Header.Set("Content-Type", "application/json")
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

	var result ServicosCorreios
	jsonErr := json.Unmarshal([]byte(jsonBody.String()), &result)
	if jsonErr != nil {
		err = fmt.Errorf("Errou ao iniciar conversão para JSON")
		return
	}

	if result.Servicos.CServico.Erro != "0" {
		err = fmt.Errorf("Errou api correios")
		return
	}

	freight = result.Servicos.CServico.Valor
	deadline = result.Servicos.CServico.PrazoEntrega + " dias úteis "

	return
}

func SenPushNotification(pu Push) (err error) {
	SERVER_API_KEY := "AAAAmyTf3h8:APA91bF2biVXH_LVgtV4OHE208XW8oicq70TKgnIB13Ae6QK5VSQoUTE8kvBixBGwTRd7jHpsacvLWUKBT09gAnE5qSIC136BN8HvzTCEGgPvEtd6sjIzg2_iNVGvJxoJFTFAOdBPI2r"

	p, err := json.Marshal(pu)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", strings.NewReader(string(p)))
	resp.Header.Set("Content-Type", "application/json")
	resp.Header.Set("Authorization", "key="+SERVER_API_KEY)
	if err != nil {
		return
	}

	_, err = http.DefaultClient.Do(resp)
	if err != nil {
		return
	}

	return
}

func SearcheAddressViaCep() {
	return
}

func SearchServiceCorreiosForDescription() {
	return
}

func SearcheServiceCorreiosForService() {
	return
}
func SearchServiceCorreiosAmount() {
	return
}
func SearchServiceCorreios() {
	return
}
