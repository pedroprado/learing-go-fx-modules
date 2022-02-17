package repository

import (
	"context"
	"fmt"

	"example.auto.wiring/infra/database"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
)

var personTable = "person-table"

type PersonRepository struct {
	db *database.DynamoDB
}

func newPersonRepository(db *database.DynamoDB) *PersonRepository {
	return &PersonRepository{
		db: db,
	}
}

func (ref *PersonRepository) Create() error {
	person := Person{
		FirstName: "Paulo",
		LastName:  "Simoes",
	}

	dynamoItem := marshalPerson(person)

	putItem := &dynamodb.PutItemInput{
		TableName: aws.String(personTable),
		Item:      dynamoItem,
	}

	out, err := ref.db.Client.PutItem(context.TODO(), putItem)
	if err != nil {
		return err
	}

	fmt.Println("Person created")
	fmt.Printf("%+v", out.ResultMetadata)
	return nil
}

func (ref *PersonRepository) Get(id string) error {

	getItem := &dynamodb.GetItemInput{
		TableName: aws.String(personTable),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: id},
		},
	}

	out, err := ref.db.Client.GetItem(context.TODO(), getItem)
	if err != nil {
		return err
	}

	fmt.Println("Person got")
	person := Person{}
	attributevalue.UnmarshalMap(out.Item, &person)
	fmt.Printf("%+v", person)
	return nil
}

func marshalPerson(person Person) map[string]types.AttributeValue {
	id := uuid.New().String()
	fmt.Println("new person id: ", id)
	return map[string]types.AttributeValue{
		"Id":        &types.AttributeValueMemberS{Value: id},
		"FirstName": &types.AttributeValueMemberS{Value: person.FirstName},
		"LastName":  &types.AttributeValueMemberS{Value: person.LastName},
	}
}
