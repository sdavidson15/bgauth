package authorizer

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/sdavidson15/bgauth/model"
)

type clientAuthorizer struct {
	baseUrl string
}

func NewClientAuthorizer(baseUrl string) Authorizer {
	return &clientAuthorizer{baseUrl: baseUrl}
}

func (c *clientAuthorizer) CreateScope(accessToken string, scope model.Scope) error {
	headers := map[string]string{
		`ContentType`:   `application/json`,
		`Authorization`: fmt.Sprintf("Bearer %s", accessToken),
	}
	scopes := model.Scopes{Scopes: []model.Scope{scope}}
	body, err := json.Marshal(scopes)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/scope", c.baseUrl)
	_, err = httpPost(url, string(body), &headers)
	return err
}

func (c *clientAuthorizer) DeleteScope(accessToken string, scope model.Scope) (model.Scope, error) {
	// TODO: stub
	return ``, fmt.Errorf(`call to unimplemented function`)
}

func (c *clientAuthorizer) GetAccessToken(authCode string) (*model.AccessToken, error) {
	url := fmt.Sprintf("%s/auth?code=%s", c.baseUrl, authCode)
	resp, err := httpGet(url, nil)
	if err != nil {
		return nil, err
	}
	accessToken := &model.AccessToken{}
	if err := json.Unmarshal([]byte(resp), accessToken); err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (c *clientAuthorizer) GetAccessTokenBasicAuth(
	username string,
	password string,
) (*model.AccessToken, error) {
	authCode := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	headers := map[string]string{
		`Authorization`: fmt.Sprintf(`Basic %s`, authCode),
	}
	url := fmt.Sprintf("%s/auth", c.baseUrl)
	resp, err := httpGet(url, &headers)
	if err != nil {
		return nil, err
	}
	accessToken := &model.AccessToken{}
	if err := json.Unmarshal([]byte(resp), accessToken); err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (c *clientAuthorizer) GetScopes(accessToken string) (*model.Scopes, error) {
	headers := map[string]string{
		`Authorization`: fmt.Sprintf("Bearer %s", accessToken),
	}
	url := fmt.Sprintf("%s/scope", c.baseUrl)
	resp, err := httpGet(url, &headers)
	if err != nil {
		return nil, err
	}
	scopes := &model.Scopes{}
	if err := json.Unmarshal([]byte(resp), scopes); err != nil {
		return nil, err
	}
	return scopes, nil
}

func (c *clientAuthorizer) RefreshAccessToken(refreshToken string) (*model.AccessToken, error) {
	url := fmt.Sprintf("%s/refresh?token=%s", c.baseUrl, refreshToken)
	resp, err := httpGet(url, nil)
	if err != nil {
		return nil, err
	}
	accessToken := &model.AccessToken{}
	if err := json.Unmarshal([]byte(resp), accessToken); err != nil {
		return nil, err
	}
	return accessToken, nil
}
