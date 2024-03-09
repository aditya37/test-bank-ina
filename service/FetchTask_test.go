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

func TestService_FetchTask(t *testing.T) {

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		mock    func(ctx context.Context, f fields)
		args    args
		want    model.ResponseFetchTask
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
			want:    model.ResponseFetchTask{},
			wantErr: true,
		},
		{
			name: "Negative: error on fetch task",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, nil)
				f.db.EXPECT().FetchTask(ctx).Return(model.ResponseFetchTask{}, sql.ErrTxDone)
			},
			args: args{
				ctx: context.Background(),
			},
			want:    model.ResponseFetchTask{},
			wantErr: true,
		},
		{
			name: "Positive: success on fetch task",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, nil)
				f.db.EXPECT().FetchTask(ctx).Return(model.ResponseFetchTask{
					Datas: []model.TaskItem{},
				}, nil)
			},
			args: args{
				ctx: context.Background(),
			},
			want: model.ResponseFetchTask{
				Total: 0,
				Datas: []model.TaskItem{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, s := initTest(t)
			tt.mock(tt.args.ctx, f)
			got, err := s.FetchTask(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.FetchTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.FetchTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
