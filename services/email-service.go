package service

import entity "github.com/MrWormHole/go-email/entities"

type EmailService interface {
	Save(entity.Email) entity.Email
	FindAll() []entity.Email
}

type emailService struct {
	emails []entity.Email
}

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
