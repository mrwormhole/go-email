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
	dbService       service.DBService          = service.CreateDBService()
	emailService    service.EmailService       = service.CreateEmailService()
	emailController controller.EmailController = controller.CreateEmailController(emailService, dbService)
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
		// TODO: test db service after finishing up refactoring
		// TODO: gracefully handle db closing
		// TODO: duplicate ids in db? something to do with api changes
		apiRoutes.GET("/emails", emailController.FindAll)
		apiRoutes.GET("/emails/:id", emailController.Show)
		apiRoutes.POST("/emails", emailController.Save)
		apiRoutes.PUT("/emails/:id", emailController.Update)
		apiRoutes.DELETE("/emails/:id", emailController.Delete)
	}

	viewRoutes := server.Group("/views")
	{
		viewRoutes.GET("/emails", emailController.ShowAll)
	}

	server.Run(":8080")
}
