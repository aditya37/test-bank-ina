package service

import (
	"context"
	"strings"

	"github.com/adity37/task/model"
)

func (s *Service) AddTask(ctx context.Context, data model.RequestAddTask) (model.ResponseAddTask, error) {

	// session
	session, err := s.Authorization(ctx)
	if err != nil {
		return model.ResponseAddTask{}, err
	}

	// default value for status
	if data.Status == "" {
		data.Status = "PENDING"
	}

	fieldAddTask := model.RequestAddTask{
		UserId:      session.Id,
		Title:       data.Title,
		Description: data.Description,
		Status:      strings.ToUpper(data.Status),
	}
	resp, err := s.db.AddTask(ctx, fieldAddTask)
	if err != nil {
		return model.ResponseAddTask{}, err
	}
	return resp, nil
}
