package model

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	IDToken      string `json:"id_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`

	ExpiresIn int    `json:"expires_in,omitempty"`
	TokenType string `json:"token_type,omitempty"`
}

type Scope string

type Scopes struct {
	Scopes []Scope `json:"scopes"`
}
