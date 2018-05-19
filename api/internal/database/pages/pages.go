package pages

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/g-harel/targetblank/api/internal/database"
	"github.com/g-harel/targetblank/api/internal/rand"
)

// Pages represents the table of page items.
type Pages struct {
	name   string
	client database.Client
}

// New creates a new Pages object.
func New(c database.Client) *Pages {
	return &Pages{
		name:   "targetblank-pages",
		client: c,
	}
}

// Create adds a new page item.
// Final key is added to the referenced Item object.
func (p *Pages) Create(item *Item) error {
	input := &dynamodb.PutItemInput{
		TableName:           aws.String(p.name),
		ConditionExpression: aws.String("attribute_not_exists(addr)"),
	}

	for {
		item.Key = rand.String(6)
		input.Item = item.toCreateMap()

		_, err := p.client.PutItem(input)
		if err == nil {
			break
		}

		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == dynamodb.ErrCodeConditionalCheckFailedException {
				continue
			}
		}
		return err
	}

	return nil
}

// Change modifies an item's values.
func (p *Pages) Change(addr string, i *Item) error {
	expression, values := i.toUpdateExpression()

	_, err := p.client.UpdateItem(&dynamodb.UpdateItemInput{
		TableName:           aws.String(p.name),
		ConditionExpression: aws.String("attribute_exists(addr)"),
		Key: map[string]*dynamodb.AttributeValue{
			"addr": &dynamodb.AttributeValue{
				S: aws.String(addr),
			},
		},
		ExpressionAttributeValues: values,
		UpdateExpression:          aws.String(expression),
	})

	return err
}

// Delete removes an item from the table.
func (p *Pages) Delete(addr string) error {
	_, err := p.client.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(p.name),
		Key: map[string]*dynamodb.AttributeValue{
			"addr": {
				S: aws.String(addr),
			},
		},
	})

	return err
}

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
