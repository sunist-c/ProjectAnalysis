package oauth

import "CeylonPlatform/middleware/authentication"

type TokenOptions struct {
	UserID   string
	ClientID string
	Token    string
	Scope    authentication.ScopeType
}

type ClientOptions struct {
	Name        string
	RedirectURL string
	Scope       authentication.ScopeType
	Method      authentication.AuthType
}

type UserOptions struct {
	Name     string
	Password string
	Scope    authentication.ScopeType
}

type AccessTokenOptions struct {
	UserID   string
	ClientID string
	Scope    authentication.ScopeType
}

type RefreshTokenOptions struct {
	UserID   string
	ClientID string
	Scope    authentication.ScopeType
}

type AuthorizationCodeOptions struct {
	UserID      string
	ClientID    string
	Scope       authentication.ScopeType
	RedirectURL string
}
