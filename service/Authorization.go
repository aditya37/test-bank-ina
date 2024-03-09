package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adity37/task/model"
	"github.com/adity37/task/utils"
)

func (s *Service) Authorization(ctx context.Context) (model.SessionPayload, error) {
	token := fmt.Sprintf("%s", ctx.Value("access_token"))
	parseToken, err := s.auth.ParseTokenDetail(token)
	if err != nil {
		return model.SessionPayload{}, &utils.CustomError{
			Code:    498,
			Message: err.Error(),
		}
	}

	userDetail, err := s.db.GetUserByEmail(ctx,
		model.RequestRegisterUser{
			Email: parseToken.Email,
		},
	)
	if err != nil {
		return model.SessionPayload{}, &utils.CustomError{
			Code:    http.StatusUnauthorized,
			Message: "not authorized",
		}
	}
	key := fmt.Sprintf("session:%d", userDetail.UserID)
	session, err := s.redis.GetUserSession(ctx, key)
	if err != nil {
		return model.SessionPayload{}, &utils.CustomError{
			Code:    498,
			Message: "session expired",
		}
	}

	return session, nil
}
