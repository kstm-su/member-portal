package oauth2

import (
	"github.com/golang-jwt/jwt"
	"github.com/kstm-su/Member-Portal/backend/database"
	"github.com/labstack/echo/v4"
)

// OIDC Standard Claims https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#StandardClaims
type userInfoClaim struct {
	Sub      string `json:"sub" gorm:"column:user_id"`
	Nickname string `json:"nickname" gorm:"column:nickname"`
	Picture  string `json:"picture" gorm:"column:profile_image"`
	Email    string `json:"email" gorm:"column:school_email"`
}

func UserInfoEndpointHandler(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
	if !ok {
		return c.String(500, "user not found")
	}

	claims := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`

	userId := claims["user_id"].(string)

	profile := userInfoClaim{}

	database.DB.Table("users").
		Select("users.user_id, users.nickname, profiles.profile_image, contacts.school_email").
		Joins("JOIN profiles ON users.user_id = profiles.user_id").
		Joins("JOIN contacts ON users.user_id = contacts.user_id").
		Where("users.user_id = ?", userId).
		Scan(&profile)

	return c.JSON(200, profile)
}
