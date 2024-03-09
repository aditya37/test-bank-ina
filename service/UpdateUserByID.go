package service

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/adity37/task/model"
	"github.com/adity37/task/utils"
)

func (s *Service) UpdateUserByID(ctx context.Context, data model.RequestUpdateUser) (model.ResponseUpdateUser, error) {

	// auth...
	if _, err := s.Authorization(ctx); err != nil {
		return model.ResponseUpdateUser{}, err
	}

	_, err := s.db.GetUserByID(ctx,
		model.RequestGetUserByID{
			UserID: data.UserID,
		},
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ResponseUpdateUser{}, &utils.CustomError{
				Code:    http.StatusNotFound,
				Message: "user id not found",
			}
		}
		return model.ResponseUpdateUser{}, &utils.CustomError{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		}
	}

	// update
	reqUpdate := model.RequestUpdateUser{
		UserID: data.UserID,
		Name:   data.Name,
		Email:  data.Email,
	}

	if data.Password != "" {
		reqUpdate.Password = base64.StdEncoding.EncodeToString([]byte(data.Password))
	} else {
		reqUpdate.Password = ""
	}

	if err := s.db.UpdateUserByID(ctx, reqUpdate); err != nil {
		return model.ResponseUpdateUser{}, &utils.CustomError{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		}
	}
	return model.ResponseUpdateUser{
		UserID: data.UserID,
	}, nil
}
