package configs

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(newDatabaseConfig),
	fx.Provide(newAwsConfig),
)

func newDatabaseConfig() *DatabaseConfig {
	database := &DatabaseConfig{}
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	database.Host = os.Getenv("DB_HOST")
	database.Password = os.Getenv("DB_PORT")
	database.User = os.Getenv("DB_USER")
	database.Password = os.Getenv("DB_PASS")
	database.DBName = os.Getenv("DB_NAME")
	database.SSLMode = os.Getenv("DB_SSL")
	return database
}

func newAwsConfig() *AwsConfig {
	awsConfig := &AwsConfig{}
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	awsConfig.AwsRegion = os.Getenv("AWS_REGION")
	awsConfig.AwsEndpoint = os.Getenv("AWS_ENDPOINT")
	awsConfig.AwsAccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	awsConfig.AwsSecretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsConfig.AwsSessionToken = os.Getenv("AWS_SESSION_TOKEN")
	return awsConfig
}
