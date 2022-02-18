package service

import (
	"example.auto.wiring/src/core/domain"
	"example.auto.wiring/src/infra/repository"
)

type PersonService struct {
	personRepo *repository.PersonRepository
}

func newPersonService(personRepo *repository.PersonRepository) *PersonService {
	return &PersonService{
		personRepo: personRepo,
	}
}

func (ref *PersonService) Create(person domain.Person) (*domain.Person, error) {
	return ref.personRepo.Create(person)
}

func (ref *PersonService) Get(id string) (*domain.Person, error) {
	return ref.personRepo.Get(id)
}
