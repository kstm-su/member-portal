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

func AuthorizationGetEndpointHandler(c echo.Context) error {
	responseType := c.QueryParam("response_type")
	clientId := c.QueryParam("client_id")
	redirectUri := c.QueryParam("redirect_uri")
	scope := c.QueryParam("scope")
	state := c.QueryParam("state")
	codeChallenge := c.QueryParam("code_challenge")
	codeChallengeMethod := c.QueryParam("code_challenge_method")
	nonce := c.QueryParam("nonce")

	if clientId == "" || redirectUri == "" || scope == "" || responseType != "code" || state == "" || nonce == "" {
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	if codeChallenge != "" && codeChallengeMethod != "S256" {
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	conf := config.Cfg

	clientDataFile := conf.File.Base + "/clients/" + clientId + ".json"
	if _, err := os.Stat(clientDataFile); os.IsNotExist(err) {
		return c.String(http.StatusBadRequest, "Invalid client")
	}

	fileContent, err := os.ReadFile(clientDataFile)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	var clientData ClientData
	if err := json.Unmarshal(fileContent, &clientData); err != nil {
		return err // Handle error appropriately
	}

	if strings.HasPrefix(redirectUri, clientData.RedirectUri) {
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	model := map[string]interface{}{
		"clientId":            clientData.ClientId,
		"clientName":          clientData.ClientName,
		"redirectUri":         redirectUri,
		"responseType":        "code",
		"state":               state,
		"scope":               scope,
		"codeChallenge":       codeChallenge,
		"codeChallengeMethod": codeChallengeMethod,
		"nonce":               nonce,
		"applicationName":     clientData.ApplicationName,
	}

	return c.Render(http.StatusOK, "authorize", model)
}

func AuthorizationPostEndpointHandler(c echo.Context) error {

	// TODO implement this function
	return nil
}
