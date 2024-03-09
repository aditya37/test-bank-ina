package model

type (
	RequestGetUserByID struct {
		UserID int64
	}
	ResponseGetUserByID struct {
		UserID int64  `json:"user_id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
	}
)

type (
	ResponseDeleteUserByID struct {
		UserID int64 `json:"user_id"`
	}
)

type (
	RequestUpdateUser struct {
		UserID   int64  `json:"user_id,omitempty"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	ResponseUpdateUser struct {
		UserID int64 `json:"user_id"`
	}
)

type (
	RequestRegisterUser struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	ResponseRegisterUser struct {
		UserID int64 `json:"user_id"`
	}
)
