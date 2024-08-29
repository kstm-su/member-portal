package oauth2

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/kstm-su/Member-Portal/backend/config"
	"github.com/kstm-su/Member-Portal/backend/crypto"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
	"time"
)

const EXPIRES_IN = 300

type TokenData struct {
	Token        string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Nonce        string `json:"nonce"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	GrantType    string `form:"grant_type"`
	ClientId     string `form:"client_id"`
	ClientSecret string `form:"client_secret"`
	RefreshToken string `form:"refresh_token"`
	RedirectUri  string `form:"redirect_uri"`
}

type AuthorizeCodeRequest struct {
	Code         string `form:"code"`
	ClientId     string `form:"client_id"`
	ClientSecret string `form:"client_secret"`
	RedirectUri  string `form:"redirect_uri"`
	CodeVerifier string `form:"code_verifier"`
}

type TokenClaims struct {
	ClientId    string
	RedirectUri string
	Nonce       string
	Scope       string
}

func TokenEndpointHandler(c echo.Context) error {
	formParameters, _ := c.FormParams()
	grantType := formParameters.Get("grant_type")

	if grantType == "refresh_token" {
		var r RefreshTokenRequest
		if err := c.Bind(&r); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}
		if r.RefreshToken == "" {
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}
		data, err := getClientData(r.ClientId, c)
		if err != nil {
			return err
		}

		if r.ClientSecret != "" {

			//if data.ClientSecret != clientSecret {
			//	return c.JSON(http.StatusBadRequest, "Invalid client_secret")
			//}

			token, state := issueTokenWithRefreshToken(r.RefreshToken, config.Cfg)
			return c.JSON(http.StatusOK, TokenData{token, "Bearer", EXPIRES_IN, state, r.RefreshToken})
		} else {
			redirectUri := r.RedirectUri
			if redirectUri == "" {
				return c.JSON(http.StatusBadRequest, "Invalid request")
			}

			if data.RedirectUri != redirectUri {
				return c.JSON(http.StatusBadRequest, "Invalid redirect_uri")
			}

			token, state := issueTokenWithRefreshToken(r.RefreshToken, config.Cfg)
			return c.JSON(http.StatusOK, TokenData{token, "Bearer", EXPIRES_IN, state, r.RefreshToken})
		}
	} else if grantType == "authorization_code" {
		var r AuthorizeCodeRequest
		if err := c.Bind(&r); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}

		data, exists := AuthorizedData[r.Code]
		if !exists {
			delete(AuthorizedData, r.Code)
			return c.JSON(http.StatusBadRequest, "Invalid code")
		}
		delete(AuthorizedData, r.Code)

		if r.Code == "" || r.RedirectUri == "" || r.ClientId == "" || r.CodeVerifier == "" {
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}

		//if clientSecret != "" {
		//	clientData := getClientData(clientId).(ConfidentialClientData)
		//	if clientData.ClientSecret != clientSecret {
		//		return c.JSON(http.StatusBadRequest, "Invalid client_secret")
		//	}
		//}

		if data.ClientId != r.ClientId || !strings.HasPrefix(r.RedirectUri, data.RedirectUri) {
			return c.JSON(http.StatusBadRequest, "Invalid client_id or redirect_uri")
		}

		hash := sha256.Sum256([]byte(r.CodeVerifier))
		base64Hash := strings.ReplaceAll(base64.URLEncoding.EncodeToString(hash[:]), "=", "")
		if data.Challenge != base64Hash {
			return c.JSON(http.StatusBadRequest, "Invalid code_verifier")
		}

		claims := TokenClaims{r.ClientId, r.RedirectUri, data.Nonce, data.Scope}
		token := issueToken(claims, config.Cfg)
		refreshToken := issueRefreshToken(claims, config.Cfg)
		return c.JSON(http.StatusOK, TokenData{token, "Bearer", EXPIRES_IN, data.Nonce, refreshToken})
	} else {
		return c.JSON(http.StatusBadRequest, "Invalid grant_type")
	}
}

func getClientData(id string, c echo.Context) (ClientData, error) {
	clientDataFile := config.Cfg.File.Base + "/clients/" + id + ".json"
	//クライアントIDのファイルが存在しない場合
	if _, err := os.Stat(clientDataFile); os.IsNotExist(err) {
		return ClientData{}, c.String(http.StatusBadRequest, "Invalid client")
	}

	//クライアントIDのファイルを読み込む
	fileContent, err := os.ReadFile(clientDataFile)
	if err != nil {
		return ClientData{}, c.String(http.StatusInternalServerError, "Internal server error")
	}

	//クライアントIDのファイルを構造体にバインド
	var clientData ClientData
	if err := json.Unmarshal(fileContent, &clientData); err != nil {
		return ClientData{}, err // Handle error appropriately
	}

	return clientData, nil
}

func issueToken(data TokenClaims, config config.Config) string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":        config.JWT.Issuer,
		"aud":        data.ClientId,
		"nbf":        time.Now().Unix(),
		"exp":        time.Now().Add(time.Second * EXPIRES_IN).Unix(),
		"iat":        time.Now().Unix(),
		"jti":        uuid.New().String(),
		"client_id":  data.ClientId,
		"scope":      data.Scope,
		"nonce":      data.Nonce,
		"token_type": "token",
	})
	tokenString, _ := token.SignedString(crypto.GetKeys(config).PrivateKey)
	return tokenString
}

func issueRefreshToken(data TokenClaims, config config.Config) string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":        config.JWT.Issuer,
		"aud":        data.ClientId,
		"nbf":        time.Now().Unix(),
		"exp":        time.Now().Add(time.Hour * 24 * 30).Unix(),
		"iat":        time.Now().Unix(),
		"jti":        uuid.New().String(),
		"client_id":  data.ClientId,
		"scope":      data.Scope,
		"nonce":      data.Nonce,
		"token_type": "refresh_token",
	})
	tokenString, _ := token.SignedString(crypto.GetKeys(config).PrivateKey)
	return tokenString
}

func issueTokenWithRefreshToken(refreshToken string, config config.Config) (string, string) {
	token, _ := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return crypto.GetKeys(config), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	clientId := claims["client_id"].(string)
	redirectUri := claims["redirect_uri"].(string)
	scope := claims["scope"].(string)
	nonce := claims["nonce"].(string)
	data := TokenClaims{clientId, redirectUri, nonce, scope}
	newToken := issueToken(data, config)
	return newToken, nonce
}
