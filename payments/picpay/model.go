package picpay

import (
	gotoolboxtime "github.com/eucatur/go-toolbox/time"
)

type User struct {
	Id int64 `json:"id" db:"sale_id"`

	CreatedAt gotoolboxtime.Timestamp `json:"created_at" db:"sale_created_at"`
}

type Payment struct {
	Referenceid  string  `json:"referenceId"`
	Callbackurl  string  `json:"callbackUrl"`
	Returnurl    string  `json:"returnUrl"`
	Value        float64 `json:"value"`
	Expiresat    string  `json:"expiresAt"`
	Channel      string  `json:"channel"`
	Purchasemode string  `json:"purchaseMode"`
	Buyer        Pagador `json:"buyer"`
}

type Pagador struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Document  string `json:"document"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type Response struct {
	Referenceid string `json:"referenceId"`
	Paymenturl  string `json:"paymentUrl"`
	Expiresat   string `json:"expiresAt"`
	Qrcode      struct {
		Content string `json:"content"`
		Base64  string `json:"base64"`
	} `json:"qrcode"`
	Message string `json:"message"`
	Errors  []struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	} `json:"errors"`
}
