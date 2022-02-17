package database

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type DynamoDB struct {
	Client *dynamodb.Client
}
