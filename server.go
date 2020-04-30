package main

import (
	"io"
	"os"

	controller "github.com/MrWormHole/go-email/controllers"
	middleware "github.com/MrWormHole/go-email/middlewares"
	service "github.com/MrWormHole/go-email/services"
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

	server.Static("views/css", "./templates/css")
	server.Static("views/js", "./templates/js")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middleware.Fool(), middleware.BasicAuth(), gindump.Dump())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/emails", emailController.FindAll)
		apiRoutes.POST("/emails", emailController.Save)
	}

	viewRoutes := server.Group("/views")
	{
		viewRoutes.GET("/emails", emailController.ShowAll)
	}

	server.Run(":8080")
}
