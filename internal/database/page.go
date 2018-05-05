package database

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// PageItem represents a document stored in the page table.
type PageItem struct {
	Addr     string `json:"addr"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Public   bool   `json:"public"`
	Temp     bool   `json:"temporary"`
	Page     string `json:"page"`
}

// PageNotFoundError communicates the error type to the caller.
type PageNotFoundError struct{ error }

// GetPage fetches the page attribute from the item with the specified address.
func GetPage(addr string) (*PageItem, error) {
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

	item := &PageItem{}
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}
