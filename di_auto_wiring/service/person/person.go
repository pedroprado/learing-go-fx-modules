package person

import (
	"example.auto.wiring/infra/database"
	"example.auto.wiring/infra/loggerfx"
)

type PersonService struct {
	Db     *database.DynamoDB
	Logger loggerfx.Logger
}

func (ref *PersonService) Create() {
	ref.Logger.Println("Person created")
}

func (ref *PersonService) Get() {
	ref.Logger.Println("Person got")
}
