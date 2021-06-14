package db

import (
	"github.com/gocql/gocql"
	"github.com/murilloarturo/bookstore_oauth_api/src/clients/cassandra"
	"github.com/murilloarturo/bookstore_oauth_api/src/domain/access_token"
	"github.com/murilloarturo/bookstore_oauth_api/src/utils/errors"
)

const (
	getAccessTokenQuery    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	createAccessTokenQuery = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	updateExpiresQuery     = "UPDATE access_tokens SET expires =? WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (repo *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	var result access_token.AccessToken
	if err := session.Query(getAccessTokenQuery, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("No access token found for given id")
		} else {
			return nil, errors.NewInternalServerError(err.Error())
		}
	}

	return &result, nil
}

func (repo *dbRepository) Create(token access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	if err := session.Query(createAccessTokenQuery, token.AccessToken, token.UserId, token.ClientId, token.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (repo *dbRepository) UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	if err := session.Query(updateExpiresQuery, token.Expires, token.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
