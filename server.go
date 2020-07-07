package main

import (
	repository "github.com/MrWormHole/go-email/repositories/sqlite"
	"io"
	"os"

	"github.com/MrWormHole/go-email/controllers"
	"github.com/MrWormHole/go-email/middlewares"
	"github.com/MrWormHole/go-email/services"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

func init() {
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func main() {
	emailRepository, err := repository.NewSqliteRepository()
	emailService := service.CreateEmailService(emailRepository)
	emailController := controller.CreateEmailController(emailService)

	if err != nil {
		panic(err)
	}

	server := gin.New()

	server.Static("views/css", "./templates/css")
	server.Static("views/js", "./templates/js")
	//server.LoadHTMLGlob("./templates/*.html")
	server.LoadHTMLFiles("./templates/index.html")

	server.Use(gin.Recovery(), middleware.BasicAuth(), gindump.Dump())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("sendEmail", emailController.Send)
		apiRoutes.GET("/emails/:id", emailController.Find)
		apiRoutes.GET("/emails", emailController.FindAll)
		apiRoutes.DELETE("/emails/:id", emailController.Remove)
	}

	viewRoutes := server.Group("/views")
	{
		viewRoutes.GET("/emails", emailController.ShowAll)
	}

	err = server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
