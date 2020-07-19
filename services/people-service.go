package service

import (
	model "github.com/MrWormHole/go-email/models"
	repository "github.com/MrWormHole/go-email/repositories/sqlite"
)

type PeopleService interface {
	Find(id uint) model.Person
	FindAll() []model.Person
	Remove(person model.Person)
}

type peopleService struct {
	emailRepository repository.EmailRepository
}

// Creates a people service which queries
func CreatePeopleService(repository repository.EmailRepository) PeopleService {
	return &peopleService{emailRepository: repository}
}

func (s *peopleService) Find(id uint) model.Person {
	return s.emailRepository.GetPerson(id)
}

func (s *peopleService) FindAll() []model.Person {
	return s.emailRepository.GetPeople()
}

func (s *peopleService) Remove(person model.Person) {
	s.emailRepository.DeletePerson(person)
}
