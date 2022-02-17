package database

import (
	"example.auto.wiring/infra/configs"
	"example.auto.wiring/infra/loggerfx"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewDynamoDB),
	fx.Invoke(connectDB),
)

func NewDynamoDB() *DynamoDB {
	return &DynamoDB{}
}

func connectDB(lifecycle fx.Lifecycle,
	logger loggerfx.Logger,
	dynamo *DynamoDB,
	databaseConfig *configs.DatabaseConfig,
	awsConfig *configs.AwsConfig) {

}
