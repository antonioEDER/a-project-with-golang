package example

import (
	"github.com/eucatur/go-toolbox/handler"
	"github.com/labstack/echo"
)

func FindByFilterHandler(c echo.Context) (err error) {
	p := *c.Get(handler.PARAMETERS).(*Filter)

	sectionals, err := FindByFilter(p)
	if err != nil {
		return
	}

	return c.JSON(200, sectionals)
}
