package controller

import (
	"net/http"

	entity "github.com/MrWormHole/go-email/entities"
	service "github.com/MrWormHole/go-email/services"
	"github.com/gin-gonic/gin"
)

type EmailController interface {
	Save(context *gin.Context) entity.Email
	FindAll() []entity.Email
	ShowAll(context *gin.Context)
}

type emailController struct {
	service service.EmailService
}

func CreateEmailController(service service.EmailService) EmailController {
	return &emailController{service: service}
}

func (controller *emailController) Save(context *gin.Context) entity.Email {
	var email entity.Email
	context.BindJSON(&email)
	return controller.service.Save(email)
}

func (controller *emailController) FindAll() []entity.Email {
	return controller.service.FindAll()
}

func (controller *emailController) ShowAll(context *gin.Context) {
	emails := controller.service.FindAll()
	data := gin.H{
		"title":  "Email Page",
		"emails": emails,
	}
	context.HTML(http.StatusOK, "index.html", data)
}
