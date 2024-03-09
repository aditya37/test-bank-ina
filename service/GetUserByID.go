package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/adity37/task/model"
	"github.com/adity37/task/utils"
)

func (s *Service) GetUserByID(ctx context.Context, data model.RequestGetUserByID) (model.ResponseGetUserByID, error) {
	// session
	_, err := s.Authorization(ctx)
	if err != nil {
		return model.ResponseGetUserByID{}, err
	}

	resp, err := s.db.GetUserByID(ctx, data)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ResponseGetUserByID{}, &utils.CustomError{
				Code:    http.StatusNotFound,
				Message: "user id not found",
			}
		}
		return model.ResponseGetUserByID{}, &utils.CustomError{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		}
	}

	return resp, nil
}
