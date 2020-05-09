package service

import (
	"errors"

	entity "github.com/MrWormHole/go-email/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DBService interface {
	Create(entity.Email)
	Retrieve(primaryKey uint64) entity.Email
	Update(entity.Email)
	Delete(entity.Email)
	FindAll() []entity.Email
	CloseDB() error
}

type dbService struct {
	connection *gorm.DB
}

func CreateDBService() DBService {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database!")
	}
	db.AutoMigrate(&entity.Email{}, &entity.Person{})

	return &dbService{connection: db}
}

func (service *dbService) Create(email entity.Email) {
	service.connection.Create(&email)
}

func (service *dbService) Retrieve(primaryKey uint64) entity.Email {
	email := entity.Email{}
	service.connection.First(&email, primaryKey)
	return email
}

func (service *dbService) Update(email entity.Email) {
	service.connection.Save(&email)
}

func (service *dbService) Delete(email entity.Email) {
	service.connection.Delete(&email)
}

func (service *dbService) FindAll() []entity.Email {
	//var emails []entity.Email   => both same i wanted to show you for learning purposes
	emails := []entity.Email{}
	service.connection.Find(&emails)
	return emails
}

func (service *dbService) CloseDB() error {
	err := service.connection.Close()
	if err != nil {
		return errors.New("Couldn't close db connection!")
	}
	return nil
}
