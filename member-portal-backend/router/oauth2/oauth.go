package oauth2

import (
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Group) {
	e.GET("/authorize", AuthorizationGetEndpointHandler)
	e.POST("/authorize", AuthorizationPostEndpointHandler)
	e.POST("/token", TokenEndpointHandler)
	e.POST("/revoke", RevokeTokenEndpointHandler)
	e.POST("/introspect", IntrospectTokenEndpointHandler)
	e.GET("/userinfo", UserInfoEndpointHandler)

}
