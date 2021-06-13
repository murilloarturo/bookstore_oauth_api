package access_token

import (
	"strings"

	"github.com/murilloarturo/bookstore_oauth_api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("Invalid access token id")
	}
	accesToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accesToken, nil
}
