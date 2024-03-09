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

func TestService_AddTask(t *testing.T) {

	type args struct {
		ctx  context.Context
		data model.RequestAddTask
	}
	tests := []struct {
		name    string
		mock    func(ctx context.Context, f fields)
		args    args
		want    model.ResponseAddTask
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
				ctx: context.WithValue(context.Background(), "access_token", "yyy"),
				data: model.RequestAddTask{
					Title: "",
				},
			},
			want:    model.ResponseAddTask{},
			wantErr: true,
		},
		{
			name: "Negative: error on insert task",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{Id: 1}, nil)
				f.db.EXPECT().AddTask(ctx, gomock.All()).Return(model.ResponseAddTask{}, sql.ErrTxDone)
			},
			args: args{
				ctx: context.WithValue(context.Background(), "access_token", "yyy"),
				data: model.RequestAddTask{
					Title: "",
				},
			},
			want:    model.ResponseAddTask{},
			wantErr: true,
		},
		{
			name: "Positive: success insert task",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{Id: 1}, nil)
				f.db.EXPECT().AddTask(ctx, gomock.All()).Return(model.ResponseAddTask{ID: 1}, nil)
			},
			args: args{
				ctx: context.WithValue(context.Background(), "access_token", "yyy"),
				data: model.RequestAddTask{
					Title: "",
				},
			},
			want: model.ResponseAddTask{
				ID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, s := initTest(t)
			tt.mock(tt.args.ctx, f)
			got, err := s.AddTask(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
