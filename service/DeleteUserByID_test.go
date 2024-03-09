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

func TestService_DeleteUserByID(t *testing.T) {

	type args struct {
		ctx    context.Context
		userid int64
	}
	tests := []struct {
		name    string
		mock    func(ctx context.Context, f fields)
		args    args
		want    model.ResponseDeleteUserByID
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
			want:    model.ResponseDeleteUserByID{},
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
				ctx:    context.TODO(),
				userid: 1,
			},
			want:    model.ResponseDeleteUserByID{},
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
				ctx:    context.TODO(),
				userid: 1,
			},
			want:    model.ResponseDeleteUserByID{},
			wantErr: true,
		},
		{
			name: "Negative: error on delete user",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, nil)
				f.db.EXPECT().GetUserByID(ctx, gomock.All()).Return(model.ResponseGetUserByID{}, nil)
				f.db.EXPECT().DeleteUserByID(ctx, gomock.All()).Return(sql.ErrConnDone)
			},
			args: args{
				ctx:    context.TODO(),
				userid: 1,
			},
			want:    model.ResponseDeleteUserByID{},
			wantErr: true,
		},
		{
			name: "Positive: success delete user",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "hamzah@mail"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "hamzah@mail"}, nil)
				f.redis.EXPECT().GetUserSession(ctx, gomock.All()).Return(model.SessionPayload{}, nil)
				f.db.EXPECT().GetUserByID(ctx, gomock.All()).Return(model.ResponseGetUserByID{}, nil)
				f.db.EXPECT().DeleteUserByID(ctx, gomock.All()).Return(nil)
			},
			args: args{
				ctx:    context.TODO(),
				userid: 1,
			},
			want: model.ResponseDeleteUserByID{
				UserID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, s := initTest(t)
			tt.mock(tt.args.ctx, f)
			got, err := s.DeleteUserByID(tt.args.ctx, tt.args.userid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.DeleteUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
