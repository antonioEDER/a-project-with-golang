package transports

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/api-qop-v2/apisexternals"
	"github.com/api-qop-v2/partners"
	"github.com/api-qop-v2/persons"
	"github.com/api-qop-v2/products"
	"github.com/api-qop-v2/tools"

	"github.com/api-qop-v2/address"
	"github.com/api-qop-v2/config"
	"github.com/labstack/echo"

	"github.com/eucatur/go-toolbox/database"
)

func CreateTransport(transport Transport) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	var p address.PersonWeb
	p.Email = transport.Email
	p.Nome = transport.Nome
	p.Tipo = transport.Tipo
	p.Celular = transport.Contato
	p.Cpf = transport.Cpf
	p.Data_Nascimento = transport.Data_Nascimento
	p.Razao_Social = transport.Razao_Social
	p.Fantasia = transport.Fantasia
	p.Cnpj = transport.Cnpj
	p.Senha = ""
	p.Uid = ""
	p.Nome_Rede_Social = ""

	p.Cep = transport.Cep
	p.Logradouro = transport.Logradouro
	p.Bairro = transport.Bairro
	p.Complemento = transport.Complemento
	p.Numero = transport.Numero
	p.Latitude = transport.Latitude
	p.Longitude = transport.Longitude
	p.Area_Abrangencia = transport.Area_Abrangencia
	p.Cidade = transport.Cidade
	p.Uf = transport.Uf

	p.He_Principal_Parceiro = transport.He_Principal_Parceiro
	p.He_Principal = transport.He_Principal
	p.He_Entrega_Propria = transport.He_Entrega_Propria

	err = persons.SearchPersonExists(p)
	if err != nil {
		return
	}

	IdPerson, err := persons.CreatePersonDBTx(p, 1, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	p.Pessoas_Id = IdPerson
	transport.Pessoas_Id = IdPerson
	p.Parceiros_Id = transport.Parceiros_Id

	_, err = persons.CreateContactPersonTx(IdPerson, p, 1, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	if transport.Cpf != "" {
		_, err = persons.CreatePhysicalPersonTx(IdPerson, p, 1, tx)
		if err != nil {
			tx.Rollback()
			return
		}

	} else {
		_, err = persons.CreateLegalPersonTx(IdPerson, p, tx)
		if err != nil {
			tx.Rollback()
			return
		}

	}

	plaintext := []byte(p.Senha)
	passEncryp, err := tools.Encrypt(plaintext)
	if err != nil {
		tx.Rollback()
		return
	}
	p.Senha = hex.EncodeToString(passEncryp)

	idUser, err := persons.CreatePersonUserDBTx(IdPerson, "", p, tx)
	if err != nil {
		fmt.Println("CreatePersonUserDBTx(): Pessoa Fisica: " + err.Error())
		tx.Rollback()
		return
	}
	_, err = persons.CreatePersonMigrationDBTx(idUser, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = persons.CreatePersonAddressDBTx(IdPerson, p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	id, err = CreateTransportTx(transport, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func SearchTransports() {
	return
}

func SearchTransportsActive(transport Transport) (result []Transport, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	result, err = SearchTransportsActiveTx(transport, tx)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	err = tx.Commit()

	return
}

func SearchTransportsByActiveForUser(transport Transport) (result []Transport, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	listTransports, err := SearchTransportsByActiveForUserTx(transport, tx)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	addressUser, err := address.SearchAddresTx(fmt.Sprint(transport.Pessoas_Id), tx)
	if err != nil {
		return
	}

	var pt partners.Partners
	pt.Parceiros_Id = transport.Parceiros_Id
	partner, err := partners.SearchPartnerForUserTx(pt, tx)
	if err != nil {
		return
	}

	addressPartner, err := address.SearchAddresTx(fmt.Sprint(partner[0].Pessoas_Id), tx)
	if err != nil {
		return
	}

	for _, list := range listTransports {

		if list.Fantasia == "Correios" {
			var paramTransport TransportService
			paramTransport.Transportadoras_Id = list.Transportadoras_Id

			var p products.Product
			p.Parceiros_Id = transport.Parceiros_Id

			var deadline string
			var freight float64

			idsProducts := strings.Split(transport.Produtos_Ids, ",")
			for _, productId := range idsProducts {
				p.ID, _ = strconv.ParseInt(productId, 10, 64)
				listProducts, _ := products.SearchProductsForUSerFromIdTx(p, tx)

				for _, pp := range listProducts {
					cubage, _ := products.CalculateCubageProduct(pp)
					var dataCorreios apisexternals.DataFromCorreios
					dataCorreios.Cep_Origem = addressPartner[0].Cep
					dataCorreios.Cep_Destino = addressUser[0].Cep
					dataCorreios.Trasportador = "Correios"
					dataCorreios.Tipo_Entrega = list.Codigo
					dataCorreios.Altura = cubage.Dimensao_Altura
					dataCorreios.Comprimento = cubage.Dimensao_Comprimento
					dataCorreios.Largura = cubage.Dimensao_Largura
					dataCorreios.Peso = cubage.Dimensao_Peso
					dataCorreios.Valor = transport.Valor_Pedido

					freightString, deadlineCopy, _ := apisexternals.CalculateServiceCorreios(dataCorreios)
					deadline = deadlineCopy
					amount, _ := strconv.ParseFloat(strings.Replace(freightString, ",", ".", -1), 64)
					freight += amount
				}
			}
			if err == nil {
				list.Frete = freight
				list.Prazo = deadline
				result = append(result, list)
			}

		} else {

			addressTransport, _ := address.SearchAddresTx(fmt.Sprint(list.Pessoas_Id), tx)

			geoOrigin, geoDestiny, _ := address.SearchCoordenatesFromTwoUsersTx(addressUser[0].Id, addressTransport[0].Id, tx)

			km, _ := apisexternals.DistanceBetweenPointers(geoOrigin, geoDestiny)

			var t Transport
			t.Transportadoras_Servicos_Id = list.Transportadoras_Servicos_Id
			t.Valor_Pedido = transport.Valor_Pedido

			services, err := SearchServicesTransportsByActiveForUserTx(t, km, tx)

			if err == nil && len(services) > 0 {
				list.Frete = services[0].Frete
				result = append(result, list)
			}
		}
	}

	err = tx.Commit()

	return
}

func AlterTransport(transport Transport) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	var p address.PersonWeb
	p.Email = transport.Email
	p.Nome = transport.Nome
	p.Tipo = transport.Tipo
	p.Celular = transport.Contato
	p.Contato = transport.Contato
	p.Cpf = transport.Cpf
	p.Data_Nascimento = transport.Data_Nascimento
	p.Razao_Social = transport.Razao_Social
	p.Fantasia = transport.Fantasia
	p.Cnpj = transport.Cnpj
	p.Senha = ""
	p.Uid = ""
	p.Nome_Rede_Social = ""
	p.Pessoas_Contatos_Id = transport.Pessoas_Contatos_Id

	p.Cep = transport.Cep
	p.Logradouro = transport.Logradouro
	p.Bairro = transport.Bairro
	p.Complemento = transport.Complemento
	p.Numero = transport.Numero
	p.Latitude = transport.Latitude
	p.Longitude = transport.Longitude
	p.Area_Abrangencia = transport.Area_Abrangencia
	p.Cidade = transport.Cidade
	p.Uf = transport.Uf

	p.He_Principal_Parceiro = transport.He_Principal_Parceiro
	p.He_Principal = transport.He_Principal
	p.He_Entrega_Propria = transport.He_Entrega_Propria
	p.He_Ativo = transport.He_Ativo

	IdPerson := transport.Pessoas_Id

	p.Pessoas_Id = IdPerson
	p.Parceiros_Id = transport.Parceiros_Id

	err = persons.AlterContactPersonTx(transport.Pessoas_Id, p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	if transport.Cpf != "" {
		err = persons.AlterPhysicalPersonTx(transport.Pessoas_Id, p, tx)
		if err != nil {
			tx.Rollback()
			return
		}

	} else {
		err = persons.AlterLegalPersonTx(transport.Pessoas_Id, p, tx)
		if err != nil {
			tx.Rollback()
			return
		}

	}

	plaintext := []byte(p.Senha)
	passEncryp, err := tools.Encrypt(plaintext)
	if err != nil {
		tx.Rollback()
		return
	}
	p.Senha = hex.EncodeToString(passEncryp)

	err = persons.AlterPersonUserDBTx(transport.Pessoas_Id, p, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	err = persons.AlterPersonAddressDBTx(transport.Pessoas_Id, p, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = AlterTransportTx(transport, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func CreatePartnersToTransport() {
	return
}

func CreateTransports() {
	return
}

func CreateTransportLinkPartner() {
	return
}

func AlterActiveTransport() {
	return
}

func SearchTransportsByActiveForPartner(transport Transport) (resultAll []Transport, resultActive []Transport, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	resultAll, resultActive, err = SearchTransportsByActiveForPartnerTx(transport, tx)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	err = tx.Commit()

	return
}

func CreateServiceToKM(transport TransportServiceKM) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	id, err = CreateServiceToKMTx(transport, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func CreateServiceToMoney(transport TransportServiceAmount) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	id, err = CreateServiceToMoneyTx(transport, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func SearchServiceByIdToServiceToKM() {
	return
}

func SearchServiceByKmAndIdServiceTransporter() {
	return
}

func SearchServiceToMoneyById(service TransportServiceAmount) (result []TransportServiceAmount, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	result, err = SearchServiceToMoneyByIdTx(service, tx)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	err = tx.Commit()

	return
}

func SearchServiceToKmById(service TransportServiceKM) (result []TransportServiceKM, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	result, err = SearchServiceToKmByIdTx(service, tx)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	err = tx.Commit()

	return
}

func SearchServiceByIdToServiceToMoney() {
	return
}

func SearchServiceByMoneyAndIdServiceTransporter() {
	return
}

func CreateService(transport TransportService) (id int64, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	id, err = CreateServiceTx(transport, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func AlterService(transport TransportService) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	err = AlterServiceTx(transport, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func SearchServices(transport TransportService) (result []TransportService, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	result, err = SearchServicesTx(transport, tx)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	err = tx.Commit()

	return
}
func SearchServicesByDescription() {
	return
}

func SearchServicesByIdPartner(transport TransportService) (result []TransportService, err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	result, err = SearchServicesByIdPartnerTx(transport, tx)
	if err != nil {
		err = echo.NewHTTPError(401, err)
		return
	}

	err = tx.Commit()

	return
}

func SearchServiceByMoneyTransporter() {
	return
}

func AlterServiceToKM(service TransportServiceKM) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	err = AlterServiceToKMTx(service, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}

func AlterServiceToMoney(service TransportServiceAmount) (err error) {
	tx := database.MustGetByFile(config.POSTGRES_MASTER_FILE).MustBegin()

	err = AlterServiceToMoneyTx(service, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	return
}
