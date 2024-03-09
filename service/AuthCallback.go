package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/adity37/task/model"
	"github.com/adity37/task/utils"
	"golang.org/x/oauth2"
)

func (s *Service) AuthCallback(ctx context.Context, code string) (oauth2.Token, error) {
	token, err := s.auth.OauthExchange(ctx, code)
	if err != nil {
		return oauth2.Token{}, &utils.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// parse token
	userInfo, err := s.auth.ParseTokenDetail(token.AccessToken)
	if err != nil {
		return oauth2.Token{}, &utils.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// validate user by email
	reqDetailUser := model.RequestRegisterUser{
		Email: userInfo.Email,
	}
	userDetail, err := s.db.GetUserByEmail(ctx, reqDetailUser)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return oauth2.Token{}, &utils.CustomError{
				Code:    http.StatusNotFound,
				Message: "email not registered",
			}
		}
		return oauth2.Token{}, &utils.CustomError{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		}
	}

	// set session
	exp := token.Expiry.Sub(time.Now()).Seconds()
	ttl := time.Duration(exp) * time.Second
	sessionKey := fmt.Sprintf("session:%d", userDetail.UserID)

	sessionPayload := model.SessionPayload{
		Id:    userDetail.UserID,
		Email: userInfo.Email,
	}
	buf, _ := json.Marshal(sessionPayload)

	if err := s.redis.Set(ctx, sessionKey, buf, ttl); err != nil {
		return oauth2.Token{}, &utils.CustomError{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		}
	}

	return oauth2.Token{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}, nil
}
