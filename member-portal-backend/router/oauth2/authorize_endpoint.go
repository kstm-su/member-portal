package oauth2

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/kstm-su/Member-Portal/backend/config"
	"github.com/labstack/echo/v4"
)

type ClientData struct {
	ClientId        string `json:"clientId"`
	ClientName      string `json:"clientName"`
	RedirectUri     string `json:"redirectUri"`
	ApplicationName string `json:"applicationName"`
}

type AuthorizeRequest struct {
	ResponseType        string `query:"response_type"`
	ClientId            string `query:"client_id"`
	RedirectUri         string `query:"redirect_uri"`
	Scope               string `query:"scope"`
	State               string `query:"state"`
	CodeChallenge       string `query:"code_challenge"`
	CodeChallengeMethod string `query:"code_challenge_method"`
	Nonce               string `query:"nonce"`
}

func AuthorizationGetEndpointHandler(c echo.Context) error {
	var r AuthorizeRequest
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

	// TODO implement this function
	return nil
}
