package service

import (
	"context"

	"github.com/adity37/task/model"
)

func (s *Service) FetchTask(ctx context.Context) (model.ResponseFetchTask, error) {
	// autj..
	if _, err := s.Authorization(ctx); err != nil {
		return model.ResponseFetchTask{}, err
	}

	resp, err := s.db.FetchTask(ctx)
	if err != nil {
		return model.ResponseFetchTask{}, err
	}
	// get total
	total := len(resp.Datas)
	return model.ResponseFetchTask{
		Datas: resp.Datas,
		Total: total,
	}, nil
}
