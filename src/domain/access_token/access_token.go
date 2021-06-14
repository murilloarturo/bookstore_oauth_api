package access_token

import (
	"strings"
	"time"

	"github.com/murilloarturo/bookstore_oauth_api/src/utils/errors"
)

const (
	expirationTime = 24
	grantTypePassword = "password"
	grantTypeCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (token *AccessTokenRequest) Validate() *errors.RestErr {
	switch token.GrantType {
	case grantTypePassword:
		break
	case grantTypeCredentials:
		break
	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}

	return nil
}

func (token *AccessToken) Validate() *errors.RestErr {
	token.AccessToken = strings.TrimSpace(token.AccessToken)
	if token.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if token.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if token.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if token.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
	// now := time.Now().UTC()
	// expirationTime := time.Unix(at.Expires, 0)

	// return expirationTime.Before(now)
}
