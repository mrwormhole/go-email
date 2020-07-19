package middleware

import (
	"github.com/MrWormHole/go-email/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

// Auth middleware that has been used by api routes
func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		jwtService := service.CreateJWTService(os.Getenv("SALT_KEY"))
		tokenString := strings.Split(context.GetHeader("Authorization"), " ")[1]

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
