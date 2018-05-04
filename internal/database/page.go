package database

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Page represents the documents stored in the database.
type Page struct {
	Addr     string `json:"addr"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Public   bool   `json:"public"`
	Page     string `json:"page"`
	Spec     string `json:"spec"`
	Temp     bool   `json:"temporary"`
}

type PageNotFoundError struct{ error }

func GetPage(addr string) (*Page, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("targetblank-pages"),
		Key: map[string]*dynamodb.AttributeValue{
			"addr": {
				S: aws.String(addr),
			},
		},
	}

	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, PageNotFoundError{errors.New("page not found for given key")}
	}

	item := &Page{}
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}
