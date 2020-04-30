package controller

import (
	"net/http"

	entity "github.com/MrWormHole/go-email/entities"
	service "github.com/MrWormHole/go-email/services"
	"github.com/gin-gonic/gin"
)

type EmailController interface {
	Save(context *gin.Context)
	FindAll(context *gin.Context)
	ShowAll(context *gin.Context)
}

type emailController struct {
	service service.EmailService
}

func CreateEmailController(service service.EmailService) EmailController {
	return &emailController{service: service}
}

func (controller *emailController) Save(context *gin.Context) {
	var email entity.Email
	context.BindJSON(&email)
	email = controller.service.Save(email)
	context.JSON(200, email)
}

func (controller *emailController) FindAll(context *gin.Context) {
	emails := controller.service.FindAll()
	context.JSON(200, emails)
}

func (controller *emailController) ShowAll(context *gin.Context) {
	emails := controller.service.FindAll()
	data := gin.H{
		"title":  "Email Page",
		"emails": emails,
	}
	context.HTML(http.StatusOK, "index.html", data)
}
