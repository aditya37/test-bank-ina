package _interface

import (
	"context"
	"io"
	"time"

	"github.com/adity37/task/model"
	"golang.org/x/oauth2"
)

type DBReaderWriter interface {
	io.Closer
	AddTask(ctx context.Context, data model.RequestAddTask) (model.ResponseAddTask, error)
	FetchTask(ctx context.Context) (model.ResponseFetchTask, error)
	GetTaskByID(ctx context.Context, data model.RequestGetTaskByID) (model.ResponseGetTaskById, error)
	DeleteTaskByID(ctx context.Context, data model.RequestDeleteTask) error
	UpdateTask(ctx context.Context, data model.RequestUpdateTask) error

	// USER...
	GetUserByID(ctx context.Context, data model.RequestGetUserByID) (model.ResponseGetUserByID, error)
	DeleteUserByID(ctx context.Context, userid int64) error
	UpdateUserByID(ctx context.Context, data model.RequestUpdateUser) error
	RegisterUser(ctx context.Context, data model.RequestRegisterUser) (int64, error)
	GetUserByEmail(ctx context.Context, data model.RequestRegisterUser) (model.ResponseGetUserByID, error)
}

type RedisReaderWriter interface {
	io.Closer
	Set(ctx context.Context, key string, payload interface{}, ttl time.Duration) error
	GetUserSession(ctx context.Context, key string) (model.SessionPayload, error)
}

type Auth interface {
	AuthCodeURL(state string) string
	OauthExchange(ctx context.Context, code string) (*oauth2.Token, error)
	ParseTokenDetail(token string) (model.ResponseParseToken, error)
}
