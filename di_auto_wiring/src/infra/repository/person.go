package repository

import (
	"context"
	"example.auto.wiring/src/core/domain"
	"example.auto.wiring/src/infra/database"

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

func (ref *PersonRepository) Create(person domain.Person) (*domain.Person, error) {
	id := uuid.New().String()
	dynamoItem := marshalPerson(person, id)

	putItem := &dynamodb.PutItemInput{
		TableName: aws.String(personTable),
		Item:      dynamoItem,
	}

	out, err := ref.db.Client.PutItem(context.TODO(), putItem)
	if err != nil {
		return nil, err
	}

	created, _ := ref.Get(id)
	attributevalue.UnmarshalMap(out.Attributes, &person)

	return created, nil
}

func (ref *PersonRepository) Get(id string) (*domain.Person, error) {

	getItem := &dynamodb.GetItemInput{
		TableName: aws.String(personTable),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: id},
		},
	}

	out, err := ref.db.Client.GetItem(context.TODO(), getItem)
	if err != nil {
		return nil, err
	}

	person := domain.Person{}
	attributevalue.UnmarshalMap(out.Item, &person)
	if person.Id == "" {
		return nil, nil
	}

	return &person, nil
}

func marshalPerson(person domain.Person, id string) map[string]types.AttributeValue {

	return map[string]types.AttributeValue{
		"Id":        &types.AttributeValueMemberS{Value: id},
		"FirstName": &types.AttributeValueMemberS{Value: person.FirstName},
		"LastName":  &types.AttributeValueMemberS{Value: person.LastName},
	}
}
