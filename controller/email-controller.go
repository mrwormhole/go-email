package controller

import (
	"github.com/MrWormHole/go-email/entity"
	"github.com/MrWormHole/go-email/service"
	"github.com/gin-gonic/gin"
)

type EmailController interface {
	Save(context *gin.Context) entity.Email
	FindAll() []entity.Email
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
