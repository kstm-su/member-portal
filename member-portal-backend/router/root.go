package router

import (
	"github.com/kstm-su/Member-Portal/backend/config"
	"github.com/kstm-su/Member-Portal/backend/router/oauth2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

func Execute(c *config.Config) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	oauth2Router := e.Group("/oauth2")
	oauth2.Setup(oauth2Router)

	var port = c.Server.Port
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))

}
