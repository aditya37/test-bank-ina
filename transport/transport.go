package transport

import (
	"net/http"
	"strconv"

	"github.com/adity37/task/model"
	"github.com/adity37/task/service"
	"github.com/gin-gonic/gin"
)

type Transport struct {
	svc *service.Service
}

func NewTransport(svc *service.Service) *Transport {
	return &Transport{
		svc: svc,
	}
}

func (t *Transport) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// user...
func (t *Transport) Auth(ctx *gin.Context) {
	authUrl := t.svc.Auth()
	ctx.Redirect(http.StatusTemporaryRedirect, authUrl)
}

// AuthCallback
func (t *Transport) AuthCallback(ctx *gin.Context) {
	code := ctx.Request.FormValue("code")
	resp, err := t.svc.AuthCallback(ctx.Request.Context(), code)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// RegisterUser
func (t *Transport) RegisterUser(ctx *gin.Context) {
	var payload model.RequestRegisterUser
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := t.svc.RegisterUser(ctx.Request.Context(), payload)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

// GetUserByID
func (t *Transport) GetUserByID(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	req := model.RequestGetUserByID{
		UserID: int64(userId),
	}
	resp, err := t.svc.GetUserByID(ctx.Request.Context(), req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// UpdateUserByID
func (t *Transport) UpdateUserByID(ctx *gin.Context) {
	var request model.RequestUpdateUser
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userid, _ := strconv.Atoi(ctx.Param("id"))
	req := model.RequestUpdateUser{
		UserID:   int64(userid),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	resp, err := t.svc.UpdateUserByID(ctx.Request.Context(), req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// DeleteUserByID
func (t *Transport) DeleteUserByID(ctx *gin.Context) {
	userid, _ := strconv.Atoi(ctx.Param("id"))
	resp, err := t.svc.DeleteUserByID(ctx.Request.Context(), int64(userid))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusAccepted, resp)
}

// task...
func (t *Transport) CreateTask(ctx *gin.Context) {
	var request model.RequestAddTask
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := t.svc.AddTask(ctx.Request.Context(), request)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// FetchTask..
func (t *Transport) FetchTask(ctx *gin.Context) {
	resp, err := t.svc.FetchTask(ctx.Request.Context())
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// GetTaskByID
func (t *Transport) GetTaskByID(ctx *gin.Context) {
	taskId, _ := strconv.Atoi(ctx.Param("id"))
	req := model.RequestGetTaskByID{
		TaskID: int64(taskId),
	}
	resp, err := t.svc.GetTaskByID(ctx.Request.Context(), req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// UpdateTask
func (t *Transport) UpdateTask(ctx *gin.Context) {
	var request model.RequestUpdateTask
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskId, _ := strconv.Atoi(ctx.Param("id"))

	req := model.RequestUpdateTask{
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
		TaskID:      int64(taskId),
	}
	resp, err := t.svc.UpdateTask(ctx.Request.Context(), req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp)

}

// DeleteTaskByID
func (t *Transport) DeleteTaskByID(ctx *gin.Context) {
	taskId, _ := strconv.Atoi(ctx.Param("id"))
	req := model.RequestDeleteTask{
		TaskID: int64(taskId),
	}
	resp, err := t.svc.DeleteTaskByID(ctx.Request.Context(), req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
