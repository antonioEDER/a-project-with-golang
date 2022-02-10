package transports

import (
	"fmt"
	"strconv"

	"github.com/eucatur/go-toolbox/handler"
	"github.com/labstack/echo"
)

func CreateTransportHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchTransportsHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchTransportsActiveHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Transport)

	transports, err := SearchTransportsActive(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"transports": transports})
}

func SearchTransportsByActiveForUserHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Transport)

	idPerson := fmt.Sprint(c.Get("PessoasId"))
	if err != nil {
		return
	}

	id, err := strconv.ParseInt(idPerson, 0, 64)
	if err != nil {
		return
	}
	p.Pessoas_Id = id

	transports, err := SearchTransportsByActiveForUser(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"servicos": transports})
}

func AlterTransportHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Transport)

	transports, err := AlterTransport(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"transports": transports})
}

func CreatePartnersToTransportHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func CreateTransportsHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Transport)

	id, err := CreateTransport(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"id": id})
}

func CreateTransportLinkPartnerHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func AlterActiveTransportHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchTransportsByActiveForPartnerHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Transport)

	t, a, err := SearchTransportsByActiveForPartner(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{
		"transportAll":    t,
		"transportActive": a,
	})
}

func CreateServiceToKMHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*TransportServiceKM)

	id, err := CreateServiceToKM(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"id": id})
}

func CreateServiceToMoneyHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*TransportServiceAmount)

	id, err := CreateServiceToMoney(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"id": id})
}

func AlterServiceToKMHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*TransportServiceKM)

	err = AlterServiceToKM(p)
	if err != nil {
		return
	}

	return c.JSON(201, err)
}

func AlterServiceToMoneyHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*TransportServiceAmount)

	err = AlterServiceToMoney(p)
	if err != nil {
		return
	}

	return c.JSON(201, err)
}

func SearchServiceToKmByIdHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*TransportServiceKM)

	service, err := SearchServiceToKmById(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"service": service})
}

func SearchServiceToMoneyByIdHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*TransportServiceAmount)

	service, err := SearchServiceToMoneyById(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"service": service})
}

func SearchServiceByIdToServiceToKMHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchServiceByKmAndIdServiceTransporterHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchServiceByIdToServiceToMoneyHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchServiceByMoneyAndIdServiceTransporterHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func CreateServiceHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*TransportService)

	id, err := CreateService(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"id": id})
}

func AlterServiceHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*TransportService)

	err = AlterService(p)
	if err != nil {
		return
	}

	return c.JSON(201, "sucesso")
}

func SearchServicesHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*TransportService)

	services, err := SearchServices(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"services": services})
}

func SearchServicesByIdPartnerHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*TransportService)

	services, err := SearchServicesByIdPartner(p)
	if err != nil {
		return
	}

	return c.JSON(201, echo.Map{"services": services})
}

func SearchServicesByDescriptionHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}

func SearchServiceByMoneyTransporterHandler(c echo.Context) (err error) {
	return c.JSON(200, 0)
}
