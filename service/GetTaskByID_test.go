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

func TestService_GetTaskByID(t *testing.T) {

	type args struct {
		ctx  context.Context
		data model.RequestGetTaskByID
	}
	tests := []struct {
		name    string
		mock    func(ctx context.Context, f fields)
		args    args
		want    model.ResponseGetTaskById
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
				ctx:  context.Background(),
				data: model.RequestGetTaskByID{},
			},
			want:    model.ResponseGetTaskById{},
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
				data: model.RequestGetTaskByID{
					TaskID: 1,
				},
			},
			want:    model.ResponseGetTaskById{},
			wantErr: true,
		},
		{
			name: "Negative: error sql",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, nil)
				f.db.EXPECT().GetTaskByID(ctx, gomock.All()).Return(model.ResponseGetTaskById{}, sql.ErrConnDone)
			},
			args: args{
				ctx: context.Background(),
				data: model.RequestGetTaskByID{
					TaskID: 1,
				},
			},
			want:    model.ResponseGetTaskById{},
			wantErr: true,
		},
		{
			name: "Positive: valid task",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, nil)
				f.db.EXPECT().GetTaskByID(ctx, gomock.All()).Return(model.ResponseGetTaskById{ID: 1}, nil)
			},
			args: args{
				ctx: context.Background(),
				data: model.RequestGetTaskByID{
					TaskID: 1,
				},
			},
			want: model.ResponseGetTaskById{
				ID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, s := initTest(t)
			tt.mock(tt.args.ctx, f)
			got, err := s.GetTaskByID(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetTaskByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetTaskByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
