package pages

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/g-harel/targetblank/internal/database"
)

// Fetch gets the page attribute from the item with the specified address.
func (p *Pages) Fetch(addr string) (*Item, error) {
	result, err := p.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(p.name),
		Key: map[string]*dynamodb.AttributeValue{
			"addr": {
				S: aws.String(addr),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(result.Item) == 0 {
		return nil, database.ItemNotFoundError(
			errors.New("page not found for given key"),
		)
	}

	item := &Item{}
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}
