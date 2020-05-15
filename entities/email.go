package entity

// maybe change entity to model in future

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
