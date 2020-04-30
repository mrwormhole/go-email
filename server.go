package main

import (
	"io"
	"os"

	"github.com/MrWormHole/go-email/controller"
	"github.com/MrWormHole/go-email/middleware"
	"github.com/MrWormHole/go-email/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	emailService    service.EmailService       = service.CreateEmailService()
	emailController controller.EmailController = controller.CreateEmailController(emailService)
)

func init() {
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func main() {
	server := gin.New()
	server.Use(gin.Recovery(), middleware.Fool(), middleware.BasicAuth(), gindump.Dump())

	server.GET("/emails", func(context *gin.Context) {
		context.JSON(200, emailController.FindAll())
	})

	server.POST("/emails", func(context *gin.Context) {
		context.JSON(200, emailController.Save(context))
	})

	server.Run(":8080")
}
