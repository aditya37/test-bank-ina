package router

import (
	"context"
	"net/http"
	"strings"

	"github.com/adity37/task/utils"
	"github.com/gin-gonic/gin"
)

func MappingError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		// get last error
		if err := ctx.Errors.Last(); err != nil {
			// cast error struct
			switch e := err.Err.(type) {
			case *utils.CustomError:
				ctx.JSON(e.Code, e)
			default:
				errPayload := utils.CustomError{
					Code:    500,
					Message: e.Error(),
				}
				ctx.JSON(500, errPayload)
			}
			ctx.Abort()
		}
	}
}

// authorization
func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// fillter path without token
		if withoutToken := pathWithoutToken(ctx); withoutToken {
			ctx.Next()
			return
		}

		// get auth header
		authHeader := ctx.Request.Header.Get("authorization")
		if authHeader == "" {
			errPayload := utils.CustomError{
				Code:    http.StatusBadRequest,
				Message: "auth header not found",
			}
			ctx.JSON(http.StatusBadRequest, errPayload)
			ctx.Abort()
			return
		}

		// validate token type
		// Check if the header starts with "Bearer"
		if !strings.HasPrefix(authHeader, "Bearer ") {
			errPayload := utils.CustomError{
				Code:    http.StatusBadRequest,
				Message: "invalid header format",
			}
			ctx.JSON(http.StatusBadRequest, errPayload)
			ctx.Abort()
			return
		}

		// Extract the token part after "Bearer "
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// assign token to context
		cc := context.WithValue(ctx.Request.Context(), "access_token", token)
		ctx.Request = ctx.Request.WithContext(cc)
		ctx.Next()
	}
}
func pathWithoutToken(ctx *gin.Context) bool {
	switch ctx.FullPath() {
	case "/users/auth":
		return true
	case "/users/callback":
		return true
	case "/users/":
		return true
	case "/health":
		return true
	}
	return false
}
