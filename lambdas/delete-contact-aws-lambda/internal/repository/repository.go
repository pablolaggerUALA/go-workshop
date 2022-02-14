package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"uala/go-workshop/pkg/dto"
)

const (
	TableName = "Contacts_Lagger"
	Delete    = "DELETE"
)

type Repository interface {
	DeleteItem(request dto.Request) (dto.Response, error)
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

func (r *LambdaRepository) DeleteItem(request dto.Request) (dto.Response, error) {
	inputItem := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Uuid": {
				S: aws.String(request.Uuid),
			},
		},
		TableName: aws.String(r.TableName),
	}

	// Delete item from dynamodb table
	_, err := r.svc.DeleteItem(inputItem)
	if err != nil {
		return dto.Response{}, &dto.DynamoDbError{
			Op:  Delete,
			Err: dto.DeleteItemError,
		}
	}

	log.Printf("Deleted item with UUID: %s", &request.Uuid)

	return dto.Response{Uuid: request.Uuid, Message: "Item deleted"}, nil
}
