package db

import (
	accesstoken "github.com/fabres21s/bookstore_oauth-api/src/domain/access_token"
	"github.com/fabres21s/bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*accesstoken.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*accesstoken.AccessToken, *errors.RestErr) {

	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
