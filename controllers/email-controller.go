package controller

import (
	"net/http"
	"strconv"

	entity "github.com/MrWormHole/go-email/entities"
	service "github.com/MrWormHole/go-email/services"
	"github.com/gin-gonic/gin"
)

// EmailController is for performing routes' actions
type EmailController interface {
	Save(context *gin.Context)
	Update(context *gin.Context)
	FindAll(context *gin.Context)
	ShowAll(context *gin.Context)
	Show(context *gin.Context)
	Delete(context *gin.Context)
}

type emailController struct {
	emailService service.EmailService
	dbService    service.DBService
}

// CreateEmailController creates an email controller
func CreateEmailController(emailService service.EmailService, dbService service.DBService) EmailController {
	return &emailController{emailService: emailService, dbService: dbService}
}

func (controller *emailController) Save(context *gin.Context) {
	var email entity.Email
	context.BindJSON(&email)

	//email = controller.emailService.Save(email)
	controller.dbService.Create(email)

	context.JSON(200, email)
}

func (controller *emailController) Update(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.Atoi(idString)

	var email entity.Email
	context.BindJSON(&email)
	email.ID = (uint)(id)

	controller.dbService.Update(email)
}

func (controller *emailController) FindAll(context *gin.Context) {
	emails := controller.dbService.FindAll()
	context.JSON(200, emails)
}

func (controller *emailController) ShowAll(context *gin.Context) {
	// stil using email service here for testing templates
	emails := controller.emailService.FindAll()
	data := gin.H{
		"title":  "Email Page",
		"emails": emails,
	}
	context.HTML(http.StatusOK, "index.html", data)
}

func (controller *emailController) Show(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.Atoi(idString)

	var email entity.Email
	context.BindJSON(&email)
	email.ID = (uint)(id)

	email = controller.dbService.Retrieve(email.ID)
	context.JSON(200, email)
}

func (controller *emailController) Delete(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.Atoi(idString)

	var email entity.Email
	context.BindJSON(&email)
	email.ID = (uint)(id)

	controller.dbService.Delete(email)
}
