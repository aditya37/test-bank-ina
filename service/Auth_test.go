package service

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestService_Auth(t *testing.T) {

	tests := []struct {
		name string
		mock func(f fields)
		want string
	}{
		{
			name: "Auth",
			mock: func(f fields) {
				f.auth.EXPECT().AuthCodeURL(gomock.All()).Return("oke")
			},
			want: "oke",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, s := initTest(t)
			tt.mock(f)
			if got := s.Auth(); got != tt.want {
				t.Errorf("Service.Auth() = %v, want %v", got, tt.want)
			}
		})
	}
}
