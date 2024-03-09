package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adity37/task/model"
	_interface "github.com/adity37/task/repository/interface"
	getenv "github.com/aditya37/get-env"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Auth struct {
	oauthCf *oauth2.Config
}

func NewOauth(redirectUrl string, scope []string) _interface.Auth {
	cf := &oauth2.Config{
		ClientID:     getenv.GetString("GOOGLE_CLIENT_ID", ""),
		ClientSecret: getenv.GetString("GOOGLE_CLIENT_SECRET", ""),
		RedirectURL:  redirectUrl,
		Endpoint:     google.Endpoint,
		Scopes:       scope,
	}
	return &Auth{
		oauthCf: cf,
	}
}

func (a *Auth) AuthCodeURL(state string) string {
	return a.oauthCf.AuthCodeURL(state)
}
func (a *Auth) OauthExchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return a.oauthCf.Exchange(ctx, code)
}
func (a *Auth) ParseTokenDetail(token string) (model.ResponseParseToken, error) {
	urlUserDetail := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", token)
	resp, err := http.Get(urlUserDetail)
	if err != nil {
		return model.ResponseParseToken{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.ResponseParseToken{}, err
	}

	// unmarshall response
	var result model.ResponseParseToken
	if err := json.Unmarshal(body, &result); err != nil {
		return model.ResponseParseToken{}, err
	}
	return result, nil
}
