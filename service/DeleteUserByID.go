package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/adity37/task/model"
	"github.com/adity37/task/utils"
)

func (s *Service) DeleteUserByID(ctx context.Context, userid int64) (model.ResponseDeleteUserByID, error) {

	// autho...
	if _, err := s.Authorization(ctx); err != nil {
		return model.ResponseDeleteUserByID{}, err
	}

	// validate user id
	fieldGetUserByID := model.RequestGetUserByID{
		UserID: userid,
	}
	_, err := s.db.GetUserByID(ctx, fieldGetUserByID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ResponseDeleteUserByID{}, &utils.CustomError{
				Code:    http.StatusNotFound,
				Message: "user id not found",
			}
		}
		return model.ResponseDeleteUserByID{}, &utils.CustomError{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		}
	}

	// delete user
	if err := s.db.DeleteUserByID(ctx, userid); err != nil {
		return model.ResponseDeleteUserByID{}, &utils.CustomError{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		}
	}

	return model.ResponseDeleteUserByID{
		UserID: userid,
	}, nil
}
