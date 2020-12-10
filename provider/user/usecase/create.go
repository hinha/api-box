package usecase

import (
	"context"
	"errors"
	"github.com/hinha/api-box/entity"
	"github.com/hinha/api-box/provider"
	"net/http"
)

// type Create
type Create struct{}

func (c *Create) Perform(ctx context.Context, userInsertable entity.GoogleUser, db provider.DB) (int, *entity.ApplicationError) {

	result, err := db.ExecContext(ctx, "", "")
	if err != nil {
		return 0, &entity.ApplicationError{
			Err:        []error{errors.New("service unavailable")},
			HTTPStatus: http.StatusServiceUnavailable,
		}
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, &entity.ApplicationError{
			Err:        []error{errors.New("internal server error when acquiring last inserted id or user")},
			HTTPStatus: http.StatusServiceUnavailable,
		}
	}
	return int(ID), nil
}
