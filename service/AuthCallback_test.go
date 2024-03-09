package service

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/adity37/task/model"
	"github.com/golang/mock/gomock"
	"golang.org/x/oauth2"
)

func TestService_AuthCallback(t *testing.T) {

	type args struct {
		ctx  context.Context
		code string
	}
	tests := []struct {
		name    string
		mock    func(ctx context.Context, f fields)
		args    args
		want    oauth2.Token
		wantErr bool
	}{
		{
			name: "Negative: error on get oauth2 exhchange",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().OauthExchange(ctx, gomock.All()).Return(&oauth2.Token{}, errors.New("failed get exchange"))
			},
			args: args{
				ctx:  context.TODO(),
				code: "19222-xc-cc",
			},
			want:    oauth2.Token{},
			wantErr: true,
		},
		{
			name: "Negative: error on parse oauth token",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().OauthExchange(ctx, gomock.All()).Return(&oauth2.Token{AccessToken: "ya29.a0Ad52N39ePTynR8sPbGxpl_HUSrMRpkVqdtOHx34xS8HANp1R5e9L1EuGdCdui5CjyuxsPDcyv_pNytVex_nT7-GgjWHtyYUSHqgZlFw90w61qugDX-mGIw_l9NTEhCOk-mCnJLwq-3AEvK1L03jlyUD20yZI-fWib0HVaCgYKATUSARISFQHGX2MiUR45xTzHqhQK6La18GXbhg0171"}, nil)
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{}, errors.New("failed to parse token"))
			},
			args: args{
				ctx:  context.TODO(),
				code: "19222-xc-cc",
			},
			want:    oauth2.Token{},
			wantErr: true,
		},
		{
			name: "Negative: error email not registered",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().OauthExchange(ctx, gomock.All()).Return(&oauth2.Token{AccessToken: "ya29.a0Ad52N39ePTynR8sPbGxpl_HUSrMRpkVqdtOHx34xS8HANp1R5e9L1EuGdCdui5CjyuxsPDcyv_pNytVex_nT7-GgjWHtyYUSHqgZlFw90w61qugDX-mGIw_l9NTEhCOk-mCnJLwq-3AEvK1L03jlyUD20yZI-fWib0HVaCgYKATUSARISFQHGX2MiUR45xTzHqhQK6La18GXbhg0171"}, nil)
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "yuhu@mail.com"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{}, sql.ErrNoRows)
			},
			args: args{
				ctx:  context.TODO(),
				code: "19222-xc-cc",
			},
			want:    oauth2.Token{},
			wantErr: true,
		},
		{
			name: "Negative: another error get detail user by email",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().OauthExchange(ctx, gomock.All()).Return(&oauth2.Token{AccessToken: "ya29.a0Ad52N39ePTynR8sPbGxpl_HUSrMRpkVqdtOHx34xS8HANp1R5e9L1EuGdCdui5CjyuxsPDcyv_pNytVex_nT7-GgjWHtyYUSHqgZlFw90w61qugDX-mGIw_l9NTEhCOk-mCnJLwq-3AEvK1L03jlyUD20yZI-fWib0HVaCgYKATUSARISFQHGX2MiUR45xTzHqhQK6La18GXbhg0171"}, nil)
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "yuhu@mail.com"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{}, sql.ErrConnDone)
			},
			args: args{
				ctx:  context.TODO(),
				code: "19222-xc-cc",
			},
			want:    oauth2.Token{},
			wantErr: true,
		},
		{
			name: "Negative: error set to redis",
			mock: func(ctx context.Context, f fields) {
				f.auth.EXPECT().OauthExchange(ctx, gomock.All()).Return(&oauth2.Token{AccessToken: "ya29.a0Ad52N39ePTynR8sPbGxpl_HUSrMRpkVqdtOHx34xS8HANp1R5e9L1EuGdCdui5CjyuxsPDcyv_pNytVex_nT7-GgjWHtyYUSHqgZlFw90w61qugDX-mGIw_l9NTEhCOk-mCnJLwq-3AEvK1L03jlyUD20yZI-fWib0HVaCgYKATUSARISFQHGX2MiUR45xTzHqhQK6La18GXbhg0171"}, nil)
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "yuhu@mail.com"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "yuhu@mail.com"}, nil)
				f.redis.EXPECT().Set(ctx, gomock.All(), gomock.All(), gomock.All()).Return(errors.New("faield to set session"))
			},
			args: args{
				ctx:  context.TODO(),
				code: "19222-xc-cc",
			},
			want:    oauth2.Token{},
			wantErr: true,
		},
		{
			name: "Positve: auth success",
			mock: func(ctx context.Context, f fields) {
				respExchange := &oauth2.Token{
					TokenType:    "Bearer",
					RefreshToken: "xxxx",
					Expiry:       time.Date(2024, 1, 1, 1, 1, 1, 0, &time.Location{}),
					AccessToken:  "ya29.a0Ad52N39ePTynR8sPbGxpl_HUSrMRpkVqdtOHx34xS8HANp1R5e9L1EuGdCdui5CjyuxsPDcyv_pNytVex_nT7-GgjWHtyYUSHqgZlFw90w61qugDX-mGIw_l9NTEhCOk-mCnJLwq-3AEvK1L03jlyUD20yZI-fWib0HVaCgYKATUSARISFQHGX2MiUR45xTzHqhQK6La18GXbhg0171",
				}
				f.auth.EXPECT().OauthExchange(ctx, gomock.All()).Return(respExchange, nil)
				f.auth.EXPECT().ParseTokenDetail(gomock.All()).Return(model.ResponseParseToken{Email: "yuhu@mail.com"}, nil)
				f.db.EXPECT().GetUserByEmail(ctx, gomock.All()).Return(model.ResponseGetUserByID{UserID: 1, Email: "yuhu@mail.com"}, nil)
				f.redis.EXPECT().Set(ctx, gomock.All(), gomock.All(), gomock.All()).Return(nil)
			},
			args: args{
				ctx:  context.TODO(),
				code: "19222-xc-cc",
			},
			want: oauth2.Token{
				TokenType:    "Bearer",
				RefreshToken: "xxxx",
				Expiry:       time.Date(2024, 1, 1, 1, 1, 1, 0, &time.Location{}),
				AccessToken:  "ya29.a0Ad52N39ePTynR8sPbGxpl_HUSrMRpkVqdtOHx34xS8HANp1R5e9L1EuGdCdui5CjyuxsPDcyv_pNytVex_nT7-GgjWHtyYUSHqgZlFw90w61qugDX-mGIw_l9NTEhCOk-mCnJLwq-3AEvK1L03jlyUD20yZI-fWib0HVaCgYKATUSARISFQHGX2MiUR45xTzHqhQK6La18GXbhg0171",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, s := initTest(t)
			tt.mock(tt.args.ctx, f)
			got, err := s.AuthCallback(tt.args.ctx, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AuthCallback() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AuthCallback() = %v, want %v", got, tt.want)
			}
		})
	}
}
