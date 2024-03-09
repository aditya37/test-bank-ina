package router

import (
	"github.com/adity37/task/transport"
	"github.com/gin-gonic/gin"
)

func NewRouter(tp *transport.Transport) *gin.Engine {
	r := gin.New()

	// middleware
	r.Use(MappingError())
	r.Use(Authorization())

	r.GET("/health", tp.HealthCheck)

	// user
	user := r.Group("/users")
	// GOOGLE AUTH
	user.GET("/auth", tp.Auth)
	user.GET("/callback", tp.AuthCallback)

	// internal auth....
	user.POST("/", tp.RegisterUser)
	user.GET("/:id", tp.GetUserByID)
	user.PUT("/:id", tp.UpdateUserByID)
	user.DELETE("/:id", tp.DeleteUserByID)

	// task
	task := r.Group("/tasks")
	task.POST("/", tp.CreateTask)
	task.GET("/", tp.FetchTask)
	task.GET("/:id", tp.GetTaskByID)
	task.PUT("/:id", tp.UpdateTask)
	task.DELETE("/:id", tp.DeleteTaskByID)

	return r
}
