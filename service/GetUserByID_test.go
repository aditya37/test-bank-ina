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

func TestService_GetUserByID(t *testing.T) {

	type args struct {
		ctx  context.Context
		data model.RequestGetUserByID
	}
	tests := []struct {
		name    string
		mock    func(ctx context.Context, f fields)
		args    args
		want    model.ResponseGetUserByID
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
				data: model.RequestGetUserByID{
					UserID: 1,
				},
			},
			want:    model.ResponseGetUserByID{},
			wantErr: true,
		},
		{
			name: "Negative: user not found",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, nil)
				f.db.EXPECT().GetUserByID(ctx, gomock.All()).Return(model.ResponseGetUserByID{}, sql.ErrNoRows)
			},
			args: args{
				ctx: context.WithValue(context.Background(), "access_token", "yyy"),
				data: model.RequestGetUserByID{
					UserID: 1,
				},
			},
			want:    model.ResponseGetUserByID{},
			wantErr: true,
		},
		{
			name: "Negative: error get user by id",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, nil)
				f.db.EXPECT().GetUserByID(ctx, gomock.All()).Return(model.ResponseGetUserByID{}, sql.ErrConnDone)
			},
			args: args{
				ctx: context.WithValue(context.Background(), "access_token", "yyy"),
				data: model.RequestGetUserByID{
					UserID: 1,
				},
			},
			want:    model.ResponseGetUserByID{},
			wantErr: true,
		},
		{
			name: "Positive: get user by id",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, nil)
				f.db.EXPECT().GetUserByID(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1}, nil)
			},
			args: args{
				ctx: context.WithValue(context.Background(), "access_token", "yyy"),
				data: model.RequestGetUserByID{
					UserID: 1,
				},
			},
			want: model.ResponseGetUserByID{
				UserID: 1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, s := initTest(t)
			tt.mock(tt.args.ctx, f)
			got, err := s.GetUserByID(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
