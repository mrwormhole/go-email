package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Person struct {
	Name string `json:"name"`
	gorm.Model
}

type Email struct {
	Sender   Person `json:"sender" gorm:"foreignkey:Name;not null"`
	Receiver Person `json:"receiver" gorm:"foreignkey:Name;not null"`
	Message  string `json:"message"`
	gorm.Model
}

func(email Email) ToString() string {
	return fmt.Sprintf("Email Sender: %s , Email Receiver: %s, Email Message: %s", email.Sender.Name, email.Receiver.Name, email.Message)
}

// this doesn't get saved to db. This is for the JSON req/res api.
type EmailTemplate struct {
	From string `json:"from"`
	To string `json:"to"`
	Subject string `json:"subject"`
	PlainText string `json:"plaintext"`
	HTMLContent string `json:"htmlcontent"`
}

