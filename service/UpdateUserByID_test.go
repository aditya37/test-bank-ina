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

func TestService_UpdateUserByID(t *testing.T) {

	type args struct {
		ctx  context.Context
		data model.RequestUpdateUser
	}
	tests := []struct {
		name    string
		mock    func(ctx context.Context, f fields)
		args    args
		want    model.ResponseUpdateUser
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
				ctx:  context.WithValue(context.Background(), "access_token", "yyy"),
				data: model.RequestUpdateUser{},
			},
			want:    model.ResponseUpdateUser{},
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
				data: model.RequestUpdateUser{
					UserID: 1,
				},
			},
			want:    model.ResponseUpdateUser{},
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
				data: model.RequestUpdateUser{
					UserID: 1,
				},
			},
			want:    model.ResponseUpdateUser{},
			wantErr: true,
		},
		{
			name: "Negative: error on update",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, nil)
				f.db.EXPECT().GetUserByID(ctx, gomock.All()).Return(model.ResponseGetUserByID{}, nil)
				f.db.EXPECT().UpdateUserByID(ctx, gomock.All()).Return(sql.ErrConnDone)
			},
			args: args{
				ctx: context.WithValue(context.Background(), "access_token", "yyy"),
				data: model.RequestUpdateUser{
					UserID: 1,
				},
			},
			want:    model.ResponseUpdateUser{},
			wantErr: true,
		},
		{
			name: "Positive: update with empty password",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, nil)
				f.db.EXPECT().GetUserByID(ctx, gomock.All()).Return(model.ResponseGetUserByID{}, nil)
				f.db.EXPECT().UpdateUserByID(ctx, gomock.All()).Return(nil)
			},
			args: args{
				ctx: context.WithValue(context.Background(), "access_token", "yyy"),
				data: model.RequestUpdateUser{
					UserID: 1,
					Name:   "",
				},
			},
			want: model.ResponseUpdateUser{
				UserID: 1,
			},
			wantErr: false,
		},
		{
			name: "Positive: update with filled password",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, nil)
				f.db.EXPECT().GetUserByID(ctx, gomock.All()).Return(model.ResponseGetUserByID{}, nil)
				f.db.EXPECT().UpdateUserByID(ctx, gomock.All()).Return(nil)
			},
			args: args{
				ctx: context.WithValue(context.Background(), "access_token", "yyy"),
				data: model.RequestUpdateUser{
					UserID:   1,
					Name:     "",
					Password: "oke",
				},
			},
			want: model.ResponseUpdateUser{
				UserID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, s := initTest(t)
			tt.mock(tt.args.ctx, f)
			got, err := s.UpdateUserByID(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UpdateUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UpdateUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
