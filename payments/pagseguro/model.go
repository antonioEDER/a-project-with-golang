package pagseguro

import (
	gotoolboxtime "github.com/eucatur/go-toolbox/time"
)

type User struct {
	Id int64 `json:"id" db:"sale_id"`

	CreatedAt gotoolboxtime.Timestamp `json:"created_at" db:"sale_created_at"`
}

type ResultSession struct {
	Session `json:"session" xml:"session"`
}

type Session struct {
	ID string `json:"id" xml:"id" xml:",chardata"`
}

type ResultTransaction struct {
	Transaction `json:"transaction" xml:"transaction"`
}
type Transaction struct {
	Date          string `xml:"date" json:"date"`
	Code          string `xml:"code" json:"code"`
	Reference     string `xml:"reference" json:"reference"`
	Type          string `xml:"type" json:"type"`
	Status        string `xml:"status" json:"status"`
	LastEventDate string `xml:"lastEventDate" json:"lastEventDate"`

	PaymentMethod struct {
		Type string `xml:"type" json:"type"`
		Code string `xml:"code" json:"code"`
	} `xml:"paymentMethod" json:"paymentMethod"`

	GrossAmount      string `xml:"grossAmount" json:"grossAmount"`
	DiscountAmount   string `xml:"discountAmount" json:"discountAmount"`
	FeeAmount        string `xml:"feeAmount" json:"feeAmount"`
	NetAmount        string `xml:"netAmount" json:"netAmount"`
	ExtraAmount      string `xml:"extraAmount" json:"extraAmount"`
	InstallmentCount string `xml:"installmentCount" json:"installmentCount"`
	ItemCount        string `xml:"itemCount" json:"itemCount"`

	Sender struct {
		Name  string `xml:"name" json:"name"`
		Email string `xml:"email" json:"email"`

		Phone struct {
			AreaCode string `xml:"areaCode" json:"areaCode"`
			Number   string `xml:"number" json:"number"`
		} `xml:"phone" json:"phone"`

		Documents struct {
			Document struct {
				Text  string `xml:",chardata" json:"text"`
				Type  string `xml:"type" json:"type"`
				Value string `xml:"value" json:"value"`
			} `xml:"document" json:"document"`
		} `xml:"documents" json:"documents"`
	} `xml:"sender" json:"sender"`

	Shipping struct {
		Address struct {
			Text       string `xml:",chardata" json:"text"`
			Street     string `xml:"street" json:"street"`
			Number     string `xml:"number" json:"number"`
			Complement string `xml:"complement" json:"complement"`
			District   string `xml:"district" json:"district"`
			City       string `xml:"city" json:"city"`
			State      string `xml:"state" json:"state"`
			Country    string `xml:"country" json:"country"`
			PostalCode string `xml:"postalCode" json:"postalCode"`
		} `xml:"address" json:"address"`
		Type string `xml:"type" json:"type"`
		Cost string `xml:"cost" json:"cost"`
	} `xml:"shipping" json:"shipping"`
	GatewaySystem struct {
		Type    string `xml:"type" json:"type"`
		RawCode struct {
			Nil string `xml:"nil,attr" json:"nil"`
			Xsi string `xml:"xsi,attr" json:"xsi"`
		} `xml:"rawCode" json:"rawCode"`
		RawMessage struct {
			Nil string `xml:"nil,attr" json:"nil"`
			Xsi string `xml:"xsi,attr" json:"xsi"`
		} `xml:"rawMessage" json:"rawMessage"`
		NormalizedCode struct {
			Nil string `xml:"nil,attr" json:"nil"`
			Xsi string `xml:"xsi,attr" json:"xsi"`
		} `xml:"normalizedCode" json:"normalizedCode"`
		NormalizedMessage struct {
			Nil string `xml:"nil,attr" json:"nil"`
			Xsi string `xml:"xsi,attr" json:"xsi"`
		} `xml:"normalizedMessage" json:"normalizedMessage"`
		AuthorizationCode string `xml:"authorizationCode" json:"authorizationCode"`
		Nsu               string `xml:"nsu" json:"nsu"`
		Tid               string `xml:"tid" json:"tid"`
		EstablishmentCode string `xml:"establishmentCode" json:"establishmentCode"`
		AcquirerName      string `xml:"acquirerName" json:"acquirerName"`
	} `xml:"gatewaySystem" json:"gatewaySystem"`
}
