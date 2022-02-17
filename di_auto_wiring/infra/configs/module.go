package configs

import (
	"os"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(newDatabaseConfig),
	fx.Provide(newAwsConfig),
	fx.Invoke(loadDatabaseConfig),
	fx.Invoke(loadAwsConfig),
)

func newDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{}
}

func loadDatabaseConfig(database *DatabaseConfig) {
	database.Host = os.Getenv("DB_HOST")
	database.Password = os.Getenv("DB_PORT")
	database.User = os.Getenv("DB_USER")
	database.Password = os.Getenv("DB_PASS")
	database.DBName = os.Getenv("DB_NAME")
	database.SSLMode = os.Getenv("DB_SSL")
}

func newAwsConfig() *AwsConfig {
	return &AwsConfig{}
}

func loadAwsConfig(awsConfig *AwsConfig) {
	awsConfig.AwsRegion = os.Getenv("AWS_REGION")
	awsConfig.AwsEndpoint = os.Getenv("AWS_ENDPOINT")
}
