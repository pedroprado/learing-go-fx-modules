package database

import (
	"context"

	"example.auto.wiring/infra/configs"
	"example.auto.wiring/infra/loggerfx"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewDynamoDB),
	fx.Invoke(initTables),
)

func NewDynamoDB(logger loggerfx.Logger,
	databaseConfig *configs.DatabaseConfig,
	awsConfig *configs.AwsConfig) *DynamoDB {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsConfig.AwsRegion),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: awsConfig.AwsEndpoint}, nil
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: awsConfig.AwsAccessKey, SecretAccessKey: awsConfig.AwsSecretKey, SessionToken: awsConfig.AwsSessionToken,
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		panic(err)
	}

	return &DynamoDB{
		Client: dynamodb.NewFromConfig(cfg),
	}
}
