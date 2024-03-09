package service

import "github.com/google/uuid"

func (s *Service) Auth() string {
	authState := uuid.New()
	return s.auth.AuthCodeURL(authState.String())
}
