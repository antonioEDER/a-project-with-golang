package user

import (
	"fmt"
	"github.com/api-qop-v2/common"
	"github.com/eucatur/go-toolbox/env"
	"github.com/eucatur/go-toolbox/jwt"
	"github.com/labstack/echo"
	"strconv"
	"strings"
)

const (
	ClaimIDKey = "id"
)

// Middleware ...
func Middleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		token := strings.TrimSpace(c.Request().Header.Get("Authorization"))
		if token == "" {
			return c.JSON(401, echo.Map{"erro": "Token vazio"})
		}

		claims, err := jwt.VerifyTokenAndGetClaims(token, env.MustString(common.EnvAPISecretKey))
		if err != nil {
			return c.JSON(401, echo.Map{"erro": "Token inválido"})
		}

		userID, err := strconv.ParseInt(fmt.Sprint(claims[ClaimIDKey]), 10, 64)
		if err != nil {
			return c.JSON(401, echo.Map{"erro": "O conteúdo do token inválido"})
		}

		c.Set(ClaimIDKey, userID)
		return next(c)
	}
}
