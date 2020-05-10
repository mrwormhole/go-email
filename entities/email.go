package entity

// maybe change entity to model in future

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Person struct {
	Name string `json:"name"`
	gorm.Model
}

type Email struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Sender    Person    `json:"sender"`
	Receiver  Person    `json:"receiver"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	gorm.Model
}
