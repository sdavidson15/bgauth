package authorizer

import "github.com/sdavidson15/bgauth/model"

type Authorizer interface {
	CreateScope(string, model.Scope) error
	DeleteScope(string, model.Scope) (model.Scope, error)
	GetAccessToken(string) (*model.AccessToken, error)
	GetAccessTokenBasicAuth(string, string) (*model.AccessToken, error)
	GetScopes(string) (*model.Scopes, error)
	RefreshAccessToken(string) (*model.AccessToken, error)
}
