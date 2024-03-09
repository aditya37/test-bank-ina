package service

import (
	"testing"

	mockrepo "github.com/adity37/task/repository/interface/mock"
	"github.com/golang/mock/gomock"
)

type fields struct {
	db    *mockrepo.MockDBReaderWriter
	redis *mockrepo.MockRedisReaderWriter
	auth  *mockrepo.MockAuth
}

func initTest(t *testing.T) (fields, *Service) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := mockrepo.NewMockDBReaderWriter(ctrl)
	redis := mockrepo.NewMockRedisReaderWriter(ctrl)
	auth := mockrepo.NewMockAuth(ctrl)
	f := fields{
		db:    db,
		redis: redis,
		auth:  auth,
	}

	svc := &Service{
		db:    db,
		redis: redis,
		auth:  auth,
	}
	s := NewService(svc.db, svc.redis, svc.auth)
	return f, s
}
