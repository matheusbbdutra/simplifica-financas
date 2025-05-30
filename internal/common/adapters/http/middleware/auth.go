package middleware

import (
	"simplificafinancas/pkg/utils"
	"strings"

	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTAuth() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: utils.GetPublicKey(),
		SigningMethod: "RS256",
		Skipper: func(c echo.Context) bool {
			path := c.Path()
 			if path == "/api/user/login" || path == "/api/user/register"  || strings.HasPrefix(path, "/public") {
                return true
            }
			return false
		},
	})
}