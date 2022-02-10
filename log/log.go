package log

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/eucatur/go-toolbox/card"
	logfile "github.com/eucatur/go-toolbox/log"

	"github.com/jmoiron/sqlx/types"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Warning ...
func Warning(text string) error {
	return logfile.File(time.Now().Format("warnings/2006/01/02/15h.log"), text)
}

// Warningf ...
func Warningf(format string, a ...interface{}) error {
	text := fmt.Sprintf(format, a...)
	return Warning(text)
}

// Middleware ...
func Middleware() echo.MiddlewareFunc {
	return middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		go func() {
			logBody := struct {
				IP            string      `json:"ip"`
				Authorization string      `json:"authorization"`
				Method        string      `json:"method"`
				URL           string      `json:"url"`
				Received      interface{} `json:"received"`
				Status        int         `json:"status"`
				Sent          interface{} `json:"sent"`
			}{
				IP:            c.RealIP(),
				Authorization: c.Request().Header.Get("Authorization"),
				Method:        c.Request().Method,
				URL:           c.Request().RequestURI,
				Received:      types.JSONText(reqBody),
				Status:        c.Response().Status,
				Sent:          types.JSONText(resBody),
			}

			logBytes, err := json.Marshal(logBody)
			if err != nil {
				fmt.Println(err)
				return
			}

			jsonOut, err := hideSensitiveData(string(logBytes))
			if err != nil {
				fmt.Println(err)
				return
			}

			err = logfile.File(time.Now().Format("requests/2006/01/02/15h.log"), jsonOut)
			if err != nil {
				fmt.Println(err)
			}
		}()
	})
}

func hideSensitiveData(jsonIn string) (jsonOut string, err error) {
	jsonOut = regexp.MustCompile(`"number":"[0-9]*"`).ReplaceAllStringFunc(jsonIn, func(substring string) string {
		pieces := strings.Split(substring, ":")
		if len(pieces) == 2 {
			cardMask, err := card.Mask(pieces[1])
			if err == nil {
				return fmt.Sprintf("%s:%s", pieces[0], cardMask)
			}
		}

		return substring
	})

	jsonOut = regexp.MustCompile(`"password":".*"`).ReplaceAllStringFunc(jsonOut, func(substring string) string {
		return `"password":"xxxxxx"`
	})

	jsonOut = regexp.MustCompile(`"cvv":"[0-9]{3}"`).ReplaceAllStringFunc(jsonOut, func(substring string) string {
		return `"cvv":"xxx"`
	})

	jsonOut = regexp.MustCompile(`"security_code":"[0-9]{3}"`).ReplaceAllStringFunc(jsonOut, func(substring string) string {
		return `"security_code":"xxx"`
	})

	return
}
