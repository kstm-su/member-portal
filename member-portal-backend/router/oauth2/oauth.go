package oauth2

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/kstm-su/Member-Portal/backend/config"
	"github.com/kstm-su/Member-Portal/backend/crypto"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

var key interface{}

func Setup(e *echo.Group) {
	e.GET("/authorize", AuthorizationGetEndpointHandler)
	e.POST("/authorize", AuthorizationPostEndpointHandler)
	e.POST("/token", TokenEndpointHandler)
	e.POST("/revoke", RevokeTokenEndpointHandler)
	e.POST("/introspect", IntrospectTokenEndpointHandler)

	key = crypto.GetKeys(config.Cfg).PublicKey

	e.GET("/userinfo", UserInfoEndpointHandler, oauth2Middleware)
}

func oauth2Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    key,
		SigningMethod: jwt.SigningMethodRS256.Name,
		ErrorHandler: func(c echo.Context, err error) error {
			var ve *jwt.ValidationError
			if errors.As(err, &ve) {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					slog.Error("token is malformed")
					return c.JSON(http.StatusBadRequest, "token is malformed")
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					slog.Warn("token is expired")
					return c.JSON(http.StatusUnauthorized, "token is expired or not valid yet")
				} else {
					slog.Warn("token is invalid")
					return c.JSON(http.StatusBadRequest, "token is invalid")
				}
			}
			println(err.Error())
			return err
		},
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			return jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
				return key, nil
			})
		},
	})(next)
}
