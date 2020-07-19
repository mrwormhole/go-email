package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Name string `json:"name"`
}

type Email struct {
	gorm.Model
	Sender   Person `json:"sender" gorm:"foreignkey:SenderID;not null"`
	SenderID uint
	Receiver Person `json:"receiver" gorm:"foreignkey:ReceiverID;not null"`
	ReceiverID uint
	Message  string `json:"message"`
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

