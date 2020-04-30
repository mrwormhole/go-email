package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Fool() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("Hey you fool! %s - [%s] %s %s  %d %s \n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC822),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency)
	})
}
