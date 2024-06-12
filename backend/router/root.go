package router

import (
	"github.com/kstm-su/Member-Portal/backend/router/oauth2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Execute() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	oauth2Router := e.Group("/oauth2")
	oauth2.Setup(oauth2Router)

	e.Logger.Fatal(e.Start(":8080"))

}
