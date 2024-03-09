package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/adity37/task/model"
	"github.com/adity37/task/utils"
)

func (s *Service) UpdateTask(ctx context.Context, data model.RequestUpdateTask) (model.ResponseUpdateTask, error) {

	_, err := s.Authorization(ctx)
	if err != nil {
		return model.ResponseUpdateTask{}, err
	}

	// validate...
	payloadGetTaskByID := model.RequestGetTaskByID{
		TaskID: data.TaskID,
	}
	if _, err := s.db.GetTaskByID(ctx, payloadGetTaskByID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ResponseUpdateTask{}, &utils.CustomError{
				Code:    http.StatusNotFound,
				Message: "task id not found",
			}
		}
		return model.ResponseUpdateTask{}, &utils.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// update
	if err := s.db.UpdateTask(ctx, data); err != nil {
		return model.ResponseUpdateTask{}, &utils.CustomError{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		}
	}

	return model.ResponseUpdateTask{
		ID: data.TaskID,
	}, nil
}
