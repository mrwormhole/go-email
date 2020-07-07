package model

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	Name string `json:"name"`
	gorm.Model
}

type Email struct {
	Sender   Person `json:"sender"`
	Receiver Person `json:"receiver"`
	Message  string `json:"message"`
	gorm.Model
}

// this doesn't get saved to db. This is for the JSON api.
type EmailTemplate struct {
	From string
	To string
	Subject string
	PlainText string
	HTMLContent string
}

