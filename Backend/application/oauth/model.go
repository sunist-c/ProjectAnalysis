package oauth

import "time"

type AuthType string

// Scope the structure of scope in oauth
type Scope struct {
	ID            string    `json:"id" xorm:"pk varchar(32) unique notnull index"`
	Name          string    `json:"name" xorm:"varchar(32) notnull"`
	ParentScopeID string    `json:"parent_scope_id" xorm:"varchar(32)"`
	Description   string    `json:"description" xorm:"notnull varchar(255)"`
	CreateAt      time.Time `json:"create_at" xorm:"notnull"`
	UpdateAt      time.Time `json:"update_at" xorm:"notnull"`
}

// User the structure of user in oauth
type User struct {
	ID       string    `json:"id" xorm:"pk varchar(32) unique notnull index"`
	Name     string    `json:"name" xorm:"varchar(32) notnull"`
	Password string    `json:"password" xorm:"varchar(32) notnull"`
	Scopes   []Scope   `json:"scopes" xorm:"notnull"`
	CreateAt time.Time `json:"create_at" xorm:"notnull"`
	UpdateAt time.Time `json:"update_at" xorm:"notnull"`
}

// Client the structure of client in oauth
type Client struct {
	ID             string    `json:"id" xorm:"notnull varchar(32) unique index pk"`
	Name           string    `json:"name" xorm:"notnull varchar(32)"`
	Key            string    `json:"key" xorm:"notnull varchar(32)"`
	Secret         string    `json:"secret" xorm:"notnull varchar(32)"`
	RedirectDomain string    `json:"redirect_domain" xorm:"varchar(255)"`
	Scopes         []Scope   `json:"scopes" xorm:"notnull"`
	Method         AuthType  `json:"method" xorm:"varchar(32) notnull default('client')"`
	CreateAt       time.Time `json:"create_at" xorm:"notnull"`
	UpdateAt       time.Time `json:"update_at" xorm:"notnull"`
}

// AccessToken the structure of access-token in oauth
type AccessToken struct {
	ID       string    `json:"id" xorm:"notnull varchar(32) unique index pk"`
	Token    string    `json:"token" xorm:"notnull varchar(32)"`
	UserID   string    `json:"user_id" xorm:"notnull varchar(32)"`
	ClientID string    `json:"client_id" xorm:"notnull varchar(32)"`
	Scopes   []Scope   `json:"scopes" xorm:"notnull"`
	ExpireAt time.Time `json:"expire_at" xorm:"notnull"`
	CreateAt time.Time `json:"create_at" xorm:"notnull"`
}

// RefreshToken the structure of refresh-token in oauth
type RefreshToken struct {
	ID       string    `json:"id" xorm:"notnull varchar(32) unique index pk"`
	Token    string    `json:"token" xorm:"notnull varchar(32)"`
	UserID   string    `json:"user_id" xorm:"notnull varchar(32)"`
	ClientID string    `json:"client_id" xorm:"notnull varchar(32)"`
	Scopes   []Scope   `json:"scopes" xorm:"notnull"`
	ExpireAt time.Time `json:"expire_at" xorm:"notnull"`
	CreateAt time.Time `json:"create_at" xorm:"notnull"`
}

// AuthorizationCode the structure of authorization code in oauth
type AuthorizationCode struct {
	ID          string    `json:"id" xorm:"notnull varchar(32) unique index pk"`
	Code        string    `json:"code" xorm:"notnull varchar(32)"`
	UserID      string    `json:"user_id" xorm:"notnull varchar(32)"`
	ClientID    string    `json:"client_id" xorm:"notnull varchar(32)"`
	Scopes      []Scope   `json:"scopes" xorm:"notnull"`
	ExpireAt    time.Time `json:"expire_at" xorm:"notnull"`
	CreateAt    time.Time `json:"create_at" xorm:"notnull"`
	RedirectUrl string    `json:"redirect_url" xorm:"notnull varchar(255)"`
}
