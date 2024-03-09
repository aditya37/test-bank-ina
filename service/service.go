package service

import (
	_interface "github.com/adity37/task/repository/interface"
)

type Service struct {
	db    _interface.DBReaderWriter
	redis _interface.RedisReaderWriter
	auth  _interface.Auth
}

func NewService(db _interface.DBReaderWriter, redis _interface.RedisReaderWriter, auth _interface.Auth) *Service {
	return &Service{
		db:    db,
		redis: redis,
		auth:  auth,
	}
}
