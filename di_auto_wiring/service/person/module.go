package person

import (
	"example.auto.wiring/infra/database"
	"example.auto.wiring/infra/loggerfx"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(newPersonService),
)

func newPersonService(db *database.DynamoDB, logger loggerfx.Logger) *PersonService {
	return &PersonService{
		Db:     db,
		Logger: logger,
	}
}
