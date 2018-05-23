package tables

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/g-harel/targetblank/api/internal/rand"
)

// IPage represents the actions that can be done on a page table.
type IPage interface {
	Create(*PageItem) error
	Change(string, *PageItem) error
	Delete(string) error
	Fetch(string) (*PageItem, error)
}

// Page represents the table of page items.
type Page struct {
	name   string
	client *dynamodb.DynamoDB
}

// NewPage creates a new Pages object.
func NewPage() IPage {
	return &Page{
		name:   "targetblank-pages",
		client: dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1")),
	}
}

// Create adds a new page item.
// Final key is added to the referenced Item object.
func (p *Page) Create(item *PageItem) error {
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
func (p *Page) Change(addr string, i *PageItem) error {
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
func (p *Page) Delete(addr string) error {
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
func (p *Page) Fetch(addr string) (*PageItem, error) {
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
		return nil, nil
	}

	item := &PageItem{}
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}
