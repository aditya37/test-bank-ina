package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/adity37/task/model"
	"github.com/adity37/task/utils"
)

func (s *Service) DeleteTaskByID(ctx context.Context, data model.RequestDeleteTask) (model.ResponseDeletTaskById, error) {
	// auth...
	if _, err := s.Authorization(ctx); err != nil {
		return model.ResponseDeletTaskById{}, err
	}

	if _, err := s.db.GetTaskByID(ctx, model.RequestGetTaskByID(data)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ResponseDeletTaskById{}, &utils.CustomError{
				Code:    http.StatusNotFound,
				Message: "task id not found",
			}
		}
		return model.ResponseDeletTaskById{}, &utils.CustomError{
			Code:    500,
			Message: err.Error(),
		}
	}

	// delete task
	if err := s.db.DeleteTaskByID(ctx, data); err != nil {
		return model.ResponseDeletTaskById{}, &utils.CustomError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return model.ResponseDeletTaskById{
		ID: data.TaskID,
	}, nil
}
