package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
)

// Auth middleware that has been used by view routes
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		os.Getenv("VIEW_USERNAME"): os.Getenv("VIEW_PASSWORD"),
	})
}
