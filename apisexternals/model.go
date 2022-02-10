package apisexternals

import (
	gotoolboxtime "github.com/eucatur/go-toolbox/time"
)

type User struct {
	Id        int64                   `json:"id" db:"sale_id"`
	CreatedAt gotoolboxtime.Timestamp `json:"created_at" db:"sale_created_at"`
}

type Pointers struct {
	DestinationAddresses []string `json:"destination_addresses"`
	OriginAddresses      []string `json:"origin_addresses"`
	Rows                 []struct {
		Elements []struct {
			Distance struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"distance"`
			Duration struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"duration"`
			Status string `json:"status"`
		} `json:"elements"`
	} `json:"rows"`
	Status string `json:"status"`
}

type DataFromCorreios struct {
	Cep_Origem   string `json:"cep_origem"`
	Cep_Destino  string `json:"cep_destino"`
	Trasportador string `json:"trasportador"`
	Tipo_Entrega string `json:"tipo_entrega"`
	Altura       string `json:"altura"`
	Comprimento  string `json:"comprimento"`
	Largura      string `json:"largura"`
	Peso         string `json:"peso"`
	Valor        string `json:"valor"`
}

type ServicosCorreios struct {
	Servicos Servicos `json:"Servicos"`
}

type Servicos struct {
	CServico CServico `json:"cServico"`
}

type CServico struct {
	Codigo                string `json:"Codigo"`
	Valor                 string `json:"Valor"`
	PrazoEntrega          string `json:"PrazoEntrega"`
	ValorSemAdicionais    string `json:"ValorSemAdicionais"`
	ValorMaoPropria       string `json:"ValorMaoPropria"`
	ValorAvisoRecebimento string `json:"ValorAvisoRecebimento"`
	ValorValorDeclarado   string `json:"ValorValorDeclarado"`
	EntregaDomiciliar     string `json:"EntregaDomiciliar"`
	EntregaSabado         string `json:"EntregaSabado"`
	ObsFim                string `json:"obsFim"`
	Erro                  string `json:"Erro"`
	MsgErro               string `json:"MsgErro"`
}

type Push struct {
	To           string       `json:"to"`
	Notification Notification `json:"notification"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
