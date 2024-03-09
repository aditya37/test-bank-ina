package model

import "time"

// RequestAddTask
type (
	RequestAddTask struct {
		UserId      int64  `json:"user_id,omitempty"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}
	ResponseAddTask struct {
		ID int64 `json:"id"`
	}
)

// ResponseFetchTask
type (
	TaskItem struct {
		UserId      int64  `json:"user_id,omitempty"`
		TaskID      int64  `json:"task_id"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}
	ResponseFetchTask struct {
		Total int        `json:"total"`
		Datas []TaskItem `json:"datas"`
	}
)

// RequestGetTaskByID
type (
	RequestGetTaskByID struct {
		TaskID int64
	}
	UserTask struct {
		UserID int64  `json:"user_id"`
		Email  string `json:"email"`
	}
	ResponseGetTaskById struct {
		ID          int64     `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		ModifiedAt  time.Time `json:"modified_at"`
		CreatedAt   time.Time `json:"created_at"`
		UserTask    UserTask  `json:"user"`
	}
)

// RequestDeleteTask
type (
	RequestDeleteTask struct {
		TaskID int64
	}
	ResponseDeletTaskById struct {
		ID int64 `json:"id"`
	}
)

type (
	RequestUpdateTask struct {
		TaskID      int64  `json:"task_id"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}
	ResponseUpdateTask struct {
		ID int64 `json:"id"`
	}
)
