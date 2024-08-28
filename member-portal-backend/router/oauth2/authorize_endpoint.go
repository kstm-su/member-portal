package oauth2

import (
	"encoding/json"
	"github.com/kstm-su/Member-Portal/backend/crypto"
	"github.com/kstm-su/Member-Portal/backend/database"
	"github.com/kstm-su/Member-Portal/backend/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kstm-su/Member-Portal/backend/config"
	"github.com/labstack/echo/v4"
)

type ClientData struct {
	ClientId        string `json:"clientId"`
	ClientName      string `json:"clientName"`
	RedirectUri     string `json:"redirectUri"`
	ApplicationName string `json:"applicationName"`
}

type AuthorizeGetRequest struct {
	ResponseType        string `query:"response_type"`
	ClientId            string `query:"client_id"`
	RedirectUri         string `query:"redirect_uri"`
	Scope               string `query:"scope"`
	State               string `query:"state"`
	CodeChallenge       string `query:"code_challenge"`
	CodeChallengeMethod string `query:"code_challenge_method"`
	Nonce               string `query:"nonce"`
}

type AuthorizePostRequest struct {
	ResponseType  string `form:"response_type"`  //code
	UserId        string `form:"userid"`         //ユーザーID
	Password      string `form:"password"`       //パスワード
	ClientId      string `form:"client_id"`      //クライアントID
	RedirectUri   string `form:"redirect_uri"`   //リダイレクトURI
	Scope         string `form:"scope"`          //スコープ
	State         string `form:"state"`          //CSRF対策
	CodeChallenge string `form:"code_challenge"` //PKCE
	Nonce         string `form:"nonce"`          //Nonce
}

type Record struct {
	Code        string    //認可コード
	ClientId    string    //public clientのチェックのため
	RedirectUri string    //public clientのチェックのため
	Scope       string    //scopeのチェックのため(tokenに入れる)
	Challenge   string    //pkceのチェックのため(hash化されたもの あとで元の値と比較)
	Nonce       string    //nonceのチェックのため(tokenに入れる) RFC 7636
	User        string    //ユーザーID
	NotAfter    time.Time //有効期限
}

func AuthorizationGetEndpointHandler(c echo.Context) error {
	var r AuthorizeGetRequest
	//構造体にリクエストの値をバインド
	if err := c.Bind(&r); err != nil {
		return err
	}
	//クエリパラメータの値をチェック
	if r.ClientId == "" || r.RedirectUri == "" || r.Scope == "" {
		return c.String(http.StatusBadRequest, "Invalid request(ClientId, RedirectUri, Scope is required)")
	}

	if r.ResponseType != "code" {
		return c.String(http.StatusBadRequest, "Response type is not supported(Only code is supported)")
	}

	if r.State == "" || r.Nonce == "" {
		return c.String(http.StatusBadRequest, "Invalid request(State, Nonce is required for security)")
	}

	if r.CodeChallenge != "" && r.CodeChallengeMethod != "S256" {
		return c.String(http.StatusBadRequest, "This server only supports S256 code challenge method")
	}
	//設定ファイルの読み込み
	conf := config.Cfg

	//クライアントIDのファイルパス
	clientDataFile := conf.File.Base + "/clients/" + r.ClientId + ".json"
	//クライアントIDのファイルが存在しない場合
	if _, err := os.Stat(clientDataFile); os.IsNotExist(err) {
		return c.String(http.StatusBadRequest, "Invalid client")
	}

	//クライアントIDのファイルを読み込む
	fileContent, err := os.ReadFile(clientDataFile)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	//クライアントIDのファイルを構造体にバインド
	var clientData ClientData
	if err := json.Unmarshal(fileContent, &clientData); err != nil {
		return err // Handle error appropriately
	}

	//リダイレクトURIが一致しない場合
	if strings.HasPrefix(r.RedirectUri, clientData.RedirectUri) {
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	//モデルの作成
	model := map[string]interface{}{
		"clientId":            clientData.ClientId,
		"clientName":          clientData.ClientName,
		"redirectUri":         r.RedirectUri,
		"responseType":        "code",
		"state":               r.State,
		"scope":               r.Scope,
		"codeChallenge":       r.CodeChallenge,
		"codeChallengeMethod": r.CodeChallengeMethod,
		"nonce":               r.Nonce,
		"applicationName":     clientData.ApplicationName,
	}

	//テンプレートをレンダリング 認証画面を表示
	return c.Render(http.StatusOK, "authorize", model)
}

func AuthorizationPostEndpointHandler(c echo.Context) error {
	var r AuthorizePostRequest
	//構造体にリクエストの値をバインド
	if err := c.Bind(&r); err != nil {
		return err
	}

	//クエリパラメータの値をチェック
	if r.UserId == "" || r.Password == "" || r.ResponseType != "code" || r.ClientId == "" || r.RedirectUri == "" || r.Scope == "" || r.State == "" || r.Nonce == "" || r.CodeChallenge == "" {
		return redirectWithError(c, r.RedirectUri, "invalid_request", "It does not have the required parameters", r.State)
	}

	userId := r.UserId
	if !isUserRegistered(userId) {
		return redirectWithError(c, r.RedirectUri, "access_denied", "This player is not registered", r.State)
	}

	hashedPassword := getUserHashedPassword(userId)
	if !crypto.VerifyPassword(hashedPassword, r.Password) {
		return redirectWithError(c, r.RedirectUri, "access_denied", "Password is incorrect", r.State)
	}

	code := crypto.GenerateRandomString(10)
	storeAuthorizedData(code, r.ClientId, r.RedirectUri, r.Scope, r.CodeChallenge, r.Nonce, userId)

	return redirectWithCode(c, r.RedirectUri, code, r.State)
}

func storeAuthorizedData(code string, ClientId string, uri string, scope string, challenge string, nonce string, user string) {
	current := time.Now()
	record := Record{
		Code:        code,
		ClientId:    ClientId,
		RedirectUri: uri,
		Scope:       scope,
		Challenge:   challenge,
		Nonce:       nonce,
		User:        user,
		NotAfter:    current.Add(time.Minute * 10),
	}

	AuthorizedData[code] = record

	//認可コードの有効期限が切れた場合削除
	go func() {
		time.Sleep(time.Minute * 10)
		delete(AuthorizedData, code)
	}()
}

var AuthorizedData = make(map[string]Record)

func getUserHashedPassword(id string) string {
	var auth models.Auth
	database.DB.Select(models.Auth{}).Where("id = ?", id).First(&auth)
	return auth.HashedPassword
}

func isUserRegistered(id string) bool {
	var user models.User
	result := database.DB.Where("id = ?", id).First(&user)
	return result.RowsAffected > 0
}

func redirectWithError(c echo.Context, redirectUri, error, errorDescription, state string) error {
	uri := strings.Builder{}
	uri.WriteString(redirectUri)
	uri.WriteString("?error=")
	uri.WriteString(error)
	uri.WriteString("&error_description=")
	uri.WriteString(errorDescription)
	uri.WriteString("&state=")
	uri.WriteString(state)
	return c.Redirect(http.StatusFound, uri.String())
}

func redirectWithCode(c echo.Context, redirectUri, code, state string) error {
	uri := strings.Builder{}
	uri.WriteString(redirectUri)
	uri.WriteString("?code=")
	uri.WriteString(code)
	uri.WriteString("&state=")
	uri.WriteString(state)
	return c.Redirect(http.StatusFound, uri.String())
}
