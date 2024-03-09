package service

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/adity37/task/model"
	"github.com/adity37/task/repository/mysql"
	"github.com/adity37/task/utils"
)

func (s *Service) RegisterUser(ctx context.Context, data model.RequestRegisterUser) (model.ResponseRegisterUser, error) {
	payload := model.RequestRegisterUser{
		Name:     data.Name,
		Email:    data.Email,
		Password: base64.StdEncoding.EncodeToString([]byte(data.Password)),
	}
	id, err := s.db.RegisterUser(ctx, payload)
	if err != nil {
		if errors.Is(err, mysql.ErrorDuplicate) {
			return model.ResponseRegisterUser{}, &utils.CustomError{
				Code:    http.StatusConflict,
				Message: err.Error(),
			}
		}
		return model.ResponseRegisterUser{}, &utils.CustomError{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		}
	}
	return model.ResponseRegisterUser{
		UserID: id,
	}, nil
}
