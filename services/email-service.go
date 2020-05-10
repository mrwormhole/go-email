package service

import entity "github.com/MrWormHole/go-email/entities"

// EmailService is for keeping data in temproary situations.
// When program restarts you will lose all the data here
type EmailService interface {
	Save(entity.Email) entity.Email
	FindAll() []entity.Email
}

type emailService struct {
	emails []entity.Email
}

// CreateEmailService creates an email service
func CreateEmailService() EmailService {
	return &emailService{}
}

func (service *emailService) Save(email entity.Email) entity.Email {
	service.emails = append(service.emails, email)
	return email
}

func (service *emailService) FindAll() []entity.Email {
	return service.emails
}
