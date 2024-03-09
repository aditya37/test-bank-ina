package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/adity37/task/model"
	"github.com/adity37/task/utils"
)

func (s *Service) GetTaskByID(ctx context.Context, data model.RequestGetTaskByID) (model.ResponseGetTaskById, error) {

	// auth...
	if _, err := s.Authorization(ctx); err != nil {
		return model.ResponseGetTaskById{}, err
	}

	payloadGetTaskByID := model.RequestGetTaskByID{
		TaskID: data.TaskID,
	}
	resp, err := s.db.GetTaskByID(ctx, payloadGetTaskByID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ResponseGetTaskById{}, &utils.CustomError{
				Code:    http.StatusNotFound,
				Message: "task id not found",
			}
		}
		return model.ResponseGetTaskById{}, &utils.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return resp, nil
}
