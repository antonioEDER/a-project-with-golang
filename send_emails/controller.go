package send_emails

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"path/filepath"
	"strings"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/api-qop-v2/common"
	"github.com/eucatur/go-toolbox/env"
)

var authSMTP = smtp.PlainAuth("", emailSac, "V4U3k&E#", "email-ssl.com.br")

const emailSac = "sac@qop.net.br"

var title = ""
var filename = ""

func NewRequest(to []string, cc []string, subject, body string, attachments map[string][]byte) *RequestEmail {
	return &RequestEmail{
		To:          to,
		Subject:     subject,
		Body:        body,
		CC:          cc,
		Attachments: attachments,
	}
}

func (r *RequestEmail) ParseTemplate(templateFileName string, data interface{}) (err error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.Body = buf.String()
	return
}

func (m *RequestEmail) ToBytes() []byte {
	buf := bytes.NewBuffer(nil)
	withAttachments := len(m.Attachments) > 0
	buf.WriteString(fmt.Sprintf("From: %s\n", title+"<"+emailSac+">"))
	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))

	buf.WriteString(fmt.Sprintf("to: %s\n", strings.Join(m.To, ",")))
	if len(m.CC) > 0 {
		buf.WriteString(fmt.Sprintf("cc: %s\n", strings.Join(m.CC, ",")))
	}

	if len(m.BCC) > 0 {
		buf.WriteString(fmt.Sprintf("bcc: %s\n", strings.Join(m.BCC, ",")))
	}

	buf.WriteString("MIME-Version: 1.0;\n")
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()

	if withAttachments {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
		buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	} else {
		buf.WriteString("Content-Type: text/html; charset=utf-8\n")
	}

	if withAttachments {
		for k, v := range m.Attachments {
			buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(v)))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))

			b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
			base64.StdEncoding.Encode(b, v)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", boundary))
		}

		buf.WriteString("--")
	}

	buf.WriteString(m.Body)

	return buf.Bytes()
}
func AttachFile(src string) (b []byte, fileName string, err error) {
	b, err = ioutil.ReadFile(src)
	if err != nil {
		return
	}

	_, fileName = filepath.Split(src)

	return
}

func (m *RequestEmail) SendEmail() error {
	host := "email-ssl.com.br:587"
	return smtp.SendMail(host, authSMTP, emailSac, m.To, m.ToBytes())
}

func SendEmailWithProposal() {
	return
}

func SendEmailWithInvoice() {
	return
}

func SendAccountCreate() {
	return
}

func SendAccountConfirmation(params EmailConfirmAccount) (err error) {
	templateData := EmailConfirmAccount{
		Name:            params.Name,
		Email:           params.Email,
		Cod_Confirmacao: params.Cod_Confirmacao,
		Token:           params.Token,
		Url:             params.Url,
	}

	r := NewRequest([]string{params.Email}, []string{}, "Olá "+params.Name, "qop - Conta", make(map[string][]byte))

	err = r.ParseTemplate("./email_template/confirm_account.html", templateData)
	if err == nil {
		title = "Confirmação de conta"
		go r.SendEmail()

	}

	return
}

func SendRecoverPassword(params EmailConfirmAccount) (err error) {
	templateData := EmailConfirmAccount{
		Name:  params.Name,
		Email: params.Email,
		Token: params.Token,
		Url:   params.Url,
	}

	r := NewRequest([]string{params.Email}, []string{}, "Olá "+params.Name, "qop - Conta", make(map[string][]byte))

	err = r.ParseTemplate("./email_template/recover-password.html", templateData)
	if err == nil {
		title = "Recuperar Senha"
		go r.SendEmail()
	} else {
		log.Panic("erro email de recuperar Senha", err)
	}

	return
}

func SendOrdersCreated(params OrderCreate) (err error) {
	templateData := OrderCreate{
		Pedidos_Id:  params.Pedidos_Id,
		Nome:        params.Nome,
		Pedido_Data: params.Pedido_Data,
		Email:       params.Email,
	}

	r := NewRequest([]string{params.Email}, []string{params.EmailParceiro}, "Olá "+params.Nome, "qop - Conta", make(map[string][]byte))

	err = r.ParseTemplate("./email_template/order_create.html", templateData)
	if err == nil {
		title = "Pedido Criado!"
		go r.SendEmail()
	} else {
		fmt.Println("Panic", err)

		log.Panic("erro email de pedido criado", err)
	}

	return
}

func SendContactClient(params Contact) (err error) {
	templateData := Contact{
		Nome:     params.Nome,
		Telefone: params.Telefone,
		Email:    params.Email,
		Assunto:  params.Assunto,
		Texto:    params.Texto,
		Title:    params.Title,
	}

	r := NewRequest([]string{params.EmailParceiro}, []string{}, "Olá "+params.NomeParceiro, "qop - Contato", make(map[string][]byte))

	err = r.ParseTemplate("./email_template/send_contact.html", templateData)
	if err == nil {
		title = "Contato enviado!"
		go r.SendEmail()

	} else {
		log.Panic("erro email de Contato enviado", err)
	}

	return
}

func SendMessageByClient() {
	return
}

func SendProductDigital(params OrderCreate, itens OrderItens) (err error) {

	rHTML := NewRequest([]string{params.Email}, []string{}, "Olá "+params.Nome, "qop - Status Pedido", make(map[string][]byte))
	err = rHTML.ParseTemplate("./email_template/enviar_produto_digital.html", itens)
	if err != nil {
		return
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	stringHTML := strings.NewReader(rHTML.Body)
	pdfg.AddPage(wkhtmltopdf.NewPageReader(stringHTML))
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	err = pdfg.Create()
	if err != nil {
		fmt.Println(err)
	}
	PathFileTemp := env.MustString(common.PathFileTemp)
	fileName := fmt.Sprint(params.ID) + fmt.Sprint((rand.Float64() * 8)) + ".pdf"
	err = pdfg.WriteFile(PathFileTemp + fileName)
	if err != nil {
		fmt.Println(err)
	}
	b, f, err := AttachFile(PathFileTemp + fileName)
	if err != nil {
		return
	}
	bytFile := make(map[string][]byte)
	bytFile[f] = b

	r := NewRequest([]string{params.Email}, []string{}, "Seu ingresso chegou", "Seu ingresso chegou", bytFile)

	r.Body = fmt.Sprintf(` Olá %s! Segue em anexo seu(s) ingresso(s). `, params.Nome)
	err = r.SendEmail()
	if err != nil {
		log.Println("Erro SendEmailBudget", err)
	}

	return
}

func SendAlterationStatusOrder(params OrderCreate) (err error) {
	templateData := OrderCreate{
		Status_Id:        params.Status_Id,
		Pedidos_Id:       params.Pedidos_Id,
		Status_Descricao: params.Status_Descricao,
		Nome:             params.Nome,
		Pedido_Data:      params.Pedido_Data,
	}

	r := NewRequest([]string{params.Email}, []string{}, "Olá "+params.Nome, "qop - Status Pedido", make(map[string][]byte))

	err = r.ParseTemplate("./email_template/status_pedido.html", templateData)
	if err == nil {
		title = "O seu pedido mudou de Status"
		go r.SendEmail()
	} else {
		log.Panic("erro email de pedido criado", err)
	}

	return
}

func SendEmailBudget(params SendBudget, filePdfName string) (err error) {
	// templateData := SendBudget{
	// 	EmailCliente:  params.EmailCliente,
	// 	EmailLoja:     params.EmailLoja,
	// 	TextoDaOferta: params.TextoDaOferta,
	// 	NomeParceiro:  params.NomeParceiro,
	// }

	filename = filePdfName
	PathFileTemp := env.MustString(common.PathFileTemp)

	b, f, err := AttachFile(PathFileTemp + filePdfName)
	if err != nil {
		return
	}

	bytFile := make(map[string][]byte)
	bytFile[f] = b

	r := NewRequest([]string{params.EmailCliente}, []string{params.EmailLoja}, "qop-marketplace", "Você recebeu uma proposta", bytFile)

	// err = r.ParseTemplate("./email_template/proposta_pedido.html", templateData)
	// if err != nil {
	// 	log.Println("Erro EmailSendOrdersCreated", err)
	// }

	r.Body = fmt.Sprintf(`
	Você recebeu a resposta do seu orçamento da empresa %s

	%s
	
	*O PDF com a proposta está em anexo
				
	Caso tenha algum problema com seu pedido clique no link abaixo:
	https://qop.net.br/minha-conta`, params.NomeParceiro, params.TextoDaOferta)

	err = r.SendEmail()
	if err != nil {
		log.Println("Erro SendEmailBudget", err)
	}

	return
}

func SendPartnerProposalToUser() {
	return
}

func SendInvoicePartner(params SendBudget, filePdfName string) (err error) {
	// templateData := SendBudget{
	// 	EmailCliente:  params.EmailCliente,
	// 	EmailLoja:     params.EmailLoja,
	// 	TextoDaOferta: params.TextoDaOferta,
	// 	NomeParceiro:  params.NomeParceiro,
	// }

	filename = filePdfName
	PathFileTemp := env.MustString(common.PathFileTemp)

	b, f, err := AttachFile(PathFileTemp + filePdfName)
	if err != nil {
		return
	}

	bytFile := make(map[string][]byte)
	bytFile[f] = b

	r := NewRequest([]string{params.EmailCliente}, []string{params.EmailLoja}, "qop-marketplace", "Sua fatura do mês.", bytFile)

	// err = r.ParseTemplate("./email_template/fatura_parceiro.html", templateData)
	// if err != nil {
	// 	log.Println("Erro SendInvoicePartner", err)
	// }

	r.Body = fmt.Sprintf(`
	Sua fatura mensal do qop acabou de chegar!
	%s
	Forma de pagamento: Boleto Bancário.

	*O PDF com o Boleto está em anexo neste email.
				
	Caso tenha algum dúvida entre em contato %s`, params.TextoDaOferta, params.EmailLoja)

	err = r.SendEmail()
	if err != nil {
		log.Println("Erro SendInvoicePartner", err)
	}

	return
}

func SendAlterationStatusOrderCancelled() {
	return
}

func SendNewLeadsPartner() {
	return
}
