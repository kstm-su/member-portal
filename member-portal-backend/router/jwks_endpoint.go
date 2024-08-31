package router

import (
	"github.com/kstm-su/Member-Portal/backend/config"
	"github.com/labstack/echo/v4"
)

func JWKsHandler(c echo.Context) error {
	println("access ")
	file := config.Cfg.File.Base + "/key/jwks.json"
	return c.File(file)
}
