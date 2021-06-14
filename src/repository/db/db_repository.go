package db

import (
	"github.com/murilloarturo/bookstore_oauth_api/src/clients/cassandra"
	"github.com/murilloarturo/bookstore_oauth_api/src/domain/access_token"
	"github.com/murilloarturo/bookstore_oauth_api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (repo *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
