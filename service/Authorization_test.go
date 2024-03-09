package service

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/adity37/task/model"
	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
)

func TestService_Authorization(t *testing.T) {

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		mock    func(ctx context.Context, f fields)
		args    args
		want    model.SessionPayload
		wantErr bool
	}{
		{
			name: "Negative: failed parse token",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{}, errors.New("faield get token detail"))
			},
			args: args{
				ctx: context.WithValue(context.Background(), "access_token", "yyy"),
			},
			wantErr: true,
		},
		{
			name: "Negative: error on get user detail by email",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{}, sql.ErrNoRows)
			},
			args: args{
				ctx: context.WithValue(context.Background(), "access_token", "yyy"),
			},
			wantErr: true,
		},
		{
			name: "Negative: error on get user user session",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, redis.Nil)
			},
			args: args{
				ctx: context.WithValue(context.Background(), "access_token", "yyy"),
			},
			wantErr: true,
		},
		{
			name: "Positive: user authorized",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{Id: 1, Email: "hamzah@mail"}, nil)
			},
			args: args{
				ctx: context.WithValue(context.Background(), "access_token", "yyy"),
			},
			wantErr: false,
			want: model.SessionPayload{
				Id:    1,
				Email: "hamzah@mail",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, s := initTest(t)
			tt.mock(tt.args.ctx, f)
			got, err := s.Authorization(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Authorization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Authorization() = %v, want %v", got, tt.want)
			}
		})
	}
}
