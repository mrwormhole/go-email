package controller

import (
	"errors"
	"net/http"
	"strconv"

	model "github.com/MrWormHole/go-email/models"
	service "github.com/MrWormHole/go-email/services"
	"github.com/gin-gonic/gin"
)

// EmailController is for performing routes' actions
type EmailController interface {
	Send(context *gin.Context)
	Find(context *gin.Context)
	FindAll(context *gin.Context)
	Remove(context *gin.Context)
	ShowAll(context *gin.Context)
}

type emailController struct {
	emailService service.EmailService
}

// CreateEmailController creates an email controller
func CreateEmailController(emailService service.EmailService) EmailController {
	return &emailController{emailService: emailService}
}

// This has been used by api
func (c *emailController) Send(context *gin.Context) {
	emailTemplate := model.EmailTemplate{}
	err := context.BindJSON(&emailTemplate)
	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{ "status:" : "Failed to bind JSON to the email template model"})
		return
	}

	_, err = c.emailService.Send(emailTemplate)
	// i feel like we should log the status of this
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "status:" : "Failed to send email"})
		return
	}

	email := model.Email{}
	email.Sender.Name = emailTemplate.From
	email.Receiver.Name = emailTemplate.To
	email.Message = emailTemplate.PlainText
	// we might need to save html later on
	c.emailService.Save(email)
	context.JSON(http.StatusOK, emailTemplate)
}

// This has been used by api
func (c *emailController) Find(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "status:" : "Failed to parse the id from URL"})
		return
	}

	email := model.Email{}
	err = context.BindJSON(&email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "status:" : "Failed to bind JSON to the email model"})
	}

	email.ID = uint(id)
	email = c.emailService.Find(email.ID)
	context.JSON(http.StatusOK, email)
}

// This has been used by api
func (c *emailController) FindAll(context *gin.Context) {
	emails := c.emailService.FindAll()
	context.JSON(http.StatusOK, emails)
}

// This has been used by api
func (c *emailController) Remove(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "status:" : "Failed to parse the id from URL"})
		return
	}

	email := model.Email{}
	err = context.BindJSON(&email)
	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{ "status:" : "Failed to bind JSON to the email model"})
		return
	}

	email.ID = (uint)(id)
	c.emailService.Remove(email)
	context.JSON(http.StatusOK, email)
}

// This has been used by index view
func (c *emailController) ShowAll(context *gin.Context) {
	emails := c.emailService.FindAll()
	data := gin.H {
		"title":  "Email Page",
		"emails": emails,
	}
	context.HTML(http.StatusOK, "index.html", data)
}
