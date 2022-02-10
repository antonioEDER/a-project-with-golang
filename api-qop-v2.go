package main

import (
	"strings"

	"github.com/api-qop-v2/address"
	"github.com/api-qop-v2/apisexternals"
	"github.com/api-qop-v2/config"
	"github.com/api-qop-v2/employees"
	"github.com/api-qop-v2/images"
	"github.com/api-qop-v2/log"
	"github.com/api-qop-v2/ordereds"
	"github.com/api-qop-v2/partners"
	"github.com/api-qop-v2/payments/pagseguro"
	"github.com/api-qop-v2/payments/picpay"
	"github.com/api-qop-v2/payments/pix"
	"github.com/api-qop-v2/persons"
	"github.com/api-qop-v2/products"
	"github.com/api-qop-v2/tools"
	"github.com/api-qop-v2/transports"
	"github.com/api-qop-v2/users"
	"github.com/eucatur/go-toolbox/api"
	"github.com/eucatur/go-toolbox/check"
	"github.com/eucatur/go-toolbox/env"
)

func main() {

	env.MustSetByJSONFile(config.ENVIRONMENT_FILE)

	// Valida parametros e ambiente
	api_environment := strings.ToUpper(env.MustString("api_environment")) //strings.ToUpper(os.Getenv("api_environment"))
	if check.If(len(api_environment) == 0 || api_environment != "HOMOLOGATION", false, true).(bool) == false &&
		check.If(api_environment != "PRODUCTION", false, true).(bool) == false {
		panic("api_environment was not set up correctly or is missing")
	}

	api.Make()
	api.Use(log.Middleware())
	api.UseCustomHTTPErrorHandler()

	api.ProvideEchoInstance(persons.AddRoutes)
	api.ProvideEchoInstance(users.AddRoutes)
	api.ProvideEchoInstance(address.AddRoutes)
	api.ProvideEchoInstance(partners.AddRoutes)
	api.ProvideEchoInstance(transports.AddRoutes)
	api.ProvideEchoInstance(products.AddRoutes)
	api.ProvideEchoInstance(employees.AddRoutes)
	api.ProvideEchoInstance(apisexternals.AddRoutes)
	api.ProvideEchoInstance(pix.AddRoutes)
	api.ProvideEchoInstance(picpay.AddRoutes)
	api.ProvideEchoInstance(ordereds.AddRoutes)
	api.ProvideEchoInstance(pagseguro.AddRoutes)
	api.ProvideEchoInstance(images.AddRoutes)
	api.ProvideEchoInstance(tools.AddRoutes)

	api.Run()
}
