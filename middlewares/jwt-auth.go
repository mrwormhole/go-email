package middleware

import (
	service "github.com/MrWormHole/go-email/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

const SALT_KEY = "HOCUSPOCUS"

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		jwtService := service.CreateJWTService(SALT_KEY)
		tokenString := context.GetHeader("Authorization")
		if len(tokenString) < 1 {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		validated, err := jwtService.Validate(tokenString)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if !validated {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
