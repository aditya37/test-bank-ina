package service

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/adity37/task/model"
	"github.com/adity37/task/repository/mysql"
	"github.com/golang/mock/gomock"
)

func TestService_RegisterUser(t *testing.T) {

	type args struct {
		ctx  context.Context
		data model.RequestRegisterUser
	}
	tests := []struct {
		name    string
		mock    func(ctx context.Context, f fields)
		args    args
		want    model.ResponseRegisterUser
		wantErr bool
	}{
		{
			name: "Negative: error duplicate",
			mock: func(ctx context.Context, f fields) {
				f.db.EXPECT().RegisterUser(ctx, gomock.All()).Return(int64(0), mysql.ErrorDuplicate)
			},
			args: args{
				ctx: context.TODO(),
				data: model.RequestRegisterUser{
					Name:     "agus",
					Email:    "agus@mail",
					Password: "6666",
				},
			},
			want:    model.ResponseRegisterUser{},
			wantErr: true,
		},
		{
			name: "Negative: error sql on insert",
			mock: func(ctx context.Context, f fields) {
				f.db.EXPECT().RegisterUser(ctx, gomock.All()).Return(int64(0), sql.ErrConnDone)
			},
			args: args{
				ctx: context.TODO(),
				data: model.RequestRegisterUser{
					Name:     "agus",
					Email:    "agus@mail",
					Password: "6666",
				},
			},
			want:    model.ResponseRegisterUser{},
			wantErr: true,
		},
		{
			name: "POsitive: success insert user",
			mock: func(ctx context.Context, f fields) {
				f.db.EXPECT().RegisterUser(ctx, gomock.All()).Return(int64(1), nil)
			},
			args: args{
				ctx: context.TODO(),
				data: model.RequestRegisterUser{
					Name:     "agus",
					Email:    "agus@mail",
					Password: "6666",
				},
			},
			want: model.ResponseRegisterUser{
				UserID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, s := initTest(t)
			tt.mock(tt.args.ctx, f)
			got, err := s.RegisterUser(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.RegisterUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
