package repository

import (
	"errors"
	model "github.com/MrWormHole/go-email/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type EmailRepository interface {
	Create(model.Email)
	Retrieve(primaryKey uint) model.Email
	Update(model.Email)
	Delete(model.Email)
	FindAll() []model.Email
	Close() error
}

type sqliteRepository struct {
	database *gorm.DB
}

// Creates email repository which uses GORM with sqlite3
func NewSqliteRepository() (EmailRepository, error) {
	db, err := gorm.Open("sqlite3","test.db")
	if err != nil {
		return nil,errors.New("Failed to connect database!")
	}
	db.AutoMigrate(&model.Email{}, &model.Person{})

	return &sqliteRepository{database: db}, nil
}

func (r *sqliteRepository) Create(email model.Email) {
	r.database.Create(&email)
}

func (r *sqliteRepository) Retrieve(id uint) model.Email {
	email := model.Email{}
	r.database.First(email, id)
	return email
}

func (r *sqliteRepository) Update(email model.Email) {
	r.database.Save(&email)
}

func (r *sqliteRepository) Delete(email model.Email) {
	r.database.Delete(&email)
}

func (r *sqliteRepository) FindAll() []model.Email  {
	emails := []model.Email{}
	r.database.Find(&emails)
	return emails
}

func (r *sqliteRepository) Close() error {
	err := r.database.Close()
	if err != nil {
		return errors.New("Failed to close database!")
	}
	return nil
}


