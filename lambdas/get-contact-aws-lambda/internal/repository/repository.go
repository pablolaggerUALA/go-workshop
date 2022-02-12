package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"uala/go-workshop/pkg/dto"
)

const (
	TableName = "Contacts_Lagger"
	Get       = "GET"
)

type Repository interface {
	GetItem(request dto.Request) (dto.Contact, error)
}

type LambdaRepository struct {
	TableName string
	svc       *dynamodb.DynamoDB
}

type GetItemOutput struct {
}

func New() Repository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &LambdaRepository{
		TableName: TableName,
		svc:       dynamodb.New(sess),
	}
}

func (r *LambdaRepository) GetItem(request dto.Request) (dto.Contact, error) {
	inputItem := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Uuid": {
				S: aws.String(request.Uuid),
			},
		},
		TableName: aws.String(r.TableName),
	}

	// Get item from dynamodb table
	dbContact, err := r.svc.GetItem(inputItem)
	if err != nil {
		return dto.Contact{}, &dto.DynamoDbError{
			Op:  Get,
			Err: dto.GetItemError,
		}
	}

	contact := dto.Contact{}
	if err := dynamodbattribute.UnmarshalMap(dbContact.Item, &contact); err != nil {
		return contact, &dto.DynamoDbError{
			Op:  Get,
			Err: dto.GetItemError,
		}
	}

	log.Printf("EVENT: %s", &dbContact.Item)

	return contact, nil
}
