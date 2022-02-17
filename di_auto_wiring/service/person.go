package service

import "example.auto.wiring/infra/repository"

type PersonService struct {
	personRepo *repository.PersonRepository
}

func newPersonService(personRepo *repository.PersonRepository) *PersonService {
	return &PersonService{
		personRepo: personRepo,
	}
}

func (ref *PersonService) Create() error {
	return ref.personRepo.Create()
}

func (ref *PersonService) Get(id string) error {
	return ref.personRepo.Get(id)
}
