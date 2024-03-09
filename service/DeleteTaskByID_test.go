package service

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/adity37/task/model"
	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
)

func TestService_DeleteTaskByID(t *testing.T) {

	type args struct {
		ctx  context.Context
		data model.RequestDeleteTask
	}
	tests := []struct {
		name    string
		mock    func(ctx context.Context, f fields)
		args    args
		want    model.ResponseDeletTaskById
		wantErr bool
	}{
		{
			name: "Negative: not authorized",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, redis.Nil)
			},
			args: args{
				ctx: context.Background(),
			},
			want:    model.ResponseDeletTaskById{},
			wantErr: true,
		},
		{
			name: "Negative: task id not found",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, nil)
				f.db.EXPECT().GetTaskByID(ctx, gomock.All()).Return(model.ResponseGetTaskById{}, sql.ErrNoRows)
			},
			args: args{
				ctx: context.Background(),
				data: model.RequestDeleteTask{
					TaskID: 1,
				},
			},
			want:    model.ResponseDeletTaskById{},
			wantErr: true,
		},
		{
			name: "Negative: error on get detail task by id",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, nil)
				f.db.EXPECT().GetTaskByID(ctx, gomock.All()).Return(model.ResponseGetTaskById{}, sql.ErrConnDone)
			},
			args: args{
				ctx: context.Background(),
				data: model.RequestDeleteTask{
					TaskID: 1,
				},
			},
			want:    model.ResponseDeletTaskById{},
			wantErr: true,
		},
		{
			name: "Negative: error on delete task",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, nil)
				f.db.EXPECT().GetTaskByID(ctx, gomock.All()).Return(model.ResponseGetTaskById{}, nil)
				f.db.EXPECT().DeleteTaskByID(ctx, gomock.All()).Return(sql.ErrConnDone)
			},
			args: args{
				ctx: context.Background(),
				data: model.RequestDeleteTask{
					TaskID: 1,
				},
			},
			want:    model.ResponseDeletTaskById{},
			wantErr: true,
		},
		{
			name: "Positive: success delete task",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, nil)
				f.db.EXPECT().GetTaskByID(ctx, gomock.All()).Return(model.ResponseGetTaskById{}, nil)
				f.db.EXPECT().DeleteTaskByID(ctx, gomock.All()).Return(nil)
			},
			args: args{
				ctx: context.Background(),
				data: model.RequestDeleteTask{
					TaskID: 1,
				},
			},
			want: model.ResponseDeletTaskById{
				ID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, s := initTest(t)
			tt.mock(tt.args.ctx, f)
			got, err := s.DeleteTaskByID(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteTaskByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.DeleteTaskByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
