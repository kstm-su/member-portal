package router

import (
	"github.com/kstm-su/Member-Portal/backend/config"
	"github.com/kstm-su/Member-Portal/backend/router/oauth2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"strconv"
	"text/template"
)

type TemplateRegistry struct {
	templates *template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Execute(c *config.Config) {
	e := echo.New()

	//テンプレートエンジンの設定
	reg := &TemplateRegistry{
		templates: template.Must(template.ParseGlob("public/view/*.html")),
	}
	e.Renderer = reg

	//ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//ルーティングの設定
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	//OAuth2のエンドポイントの設定
	oauth2Router := e.Group("/oauth2")
	oauth2.Setup(oauth2Router)

	//e.GET("/.well-known/openid-configuration", OpenIDConfigurationHandler)
	e.GET("/.well-known/jwks.json", JWKsHandler)

	e.GET("/assets/*", func(c echo.Context) error {
		return c.File("public/assets/" + c.Param("*"))
	})

	//サーバーの起動
	var port = c.Server.Port
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))

}
