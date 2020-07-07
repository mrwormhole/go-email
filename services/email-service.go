package service

import (
	"fmt"
	model "github.com/MrWormHole/go-email/models"
	repository "github.com/MrWormHole/go-email/repositories/sqlite"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"net/http"
	"strings"
)

const API_KEY = "XxXxX"

type EmailService interface {
	Send(email model.EmailTemplate) (string, error)
	Save(model.Email)
	Find(id uint) model.Email
	FindAll() []model.Email
	Remove(email model.Email)
}

type emailService struct {
	emailRepository repository.EmailRepository
}

// Creates email services which sends emails with sendgrid
func CreateEmailService(repository repository.EmailRepository) EmailService {
	return &emailService{emailRepository: repository}
}

func (s *emailService) Send(emailTemplate model.EmailTemplate) (string, error) {
	fromName := strings.ToUpper(strings.Split(emailTemplate.From, "@")[0])
	from := mail.NewEmail(fromName, emailTemplate.From)
	toName := strings.ToUpper(strings.Split(emailTemplate.From, "@")[0])
	to := mail.NewEmail(toName, emailTemplate.To)
	subject := emailTemplate.Subject
	plainTextContent := emailTemplate.PlainText
	htmlContent := emailTemplate.HTMLContent

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(API_KEY)
	response, err := client.Send(message)
	if err != nil {
		return "Failed to send email via sendgrid client", err
	}
	return fmt.Sprintf("Sent the email with status code: %s", http.StatusText(response.StatusCode)), nil
}

func (s *emailService) Save(email model.Email) {
	s.emailRepository.Create(email)
}

func (s *emailService) Find(id uint) model.Email{
	return s.emailRepository.Retrieve(id)
}

func (s *emailService) FindAll() []model.Email {
	return s.emailRepository.FindAll()
}

func (s *emailService) Remove(email model.Email) {
	s.emailRepository.Delete(email)
}
