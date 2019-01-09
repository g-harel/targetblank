package storage

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const pageTable = "targetblank-pages"
const pageKey = "addr"

// Page represents a DynamoDB page item.
type Page struct {
	Addr      string `json:"addr"` // pageKey
	Email     string `json:"email"`
	Password  string `json:"password"`
	Published bool   `json:"published"`
	Document  string `json:"document"`
}

// PageCreate writes the page to storage.
// Conflict flag will be set if the address is already taken.
func PageCreate(p *Page) (conflict bool, err error) {
	page, err := dynamodbattribute.MarshalMap(p)
	if err != nil {
		return false, err
	}

	_, err = client.PutItem(&dynamodb.PutItemInput{
		TableName:           aws.String(pageTable),
		ConditionExpression: aws.String(fmt.Sprintf("attribute_not_exists(%v)", pageKey)),
		Item:                page,
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			conflict = awsErr.Code() == dynamodb.ErrCodeConditionalCheckFailedException
		}
		return conflict, err
	}

	return false, nil
}

// PageRead reads a page from storage.
// A null value pointer for page indicates the address was not found.
func PageRead(addr string) (*Page, error) {
	result, err := client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(pageTable),
		Key: map[string]*dynamodb.AttributeValue{
			pageKey: {
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

	page := &Page{}
	err = dynamodbattribute.UnmarshalMap(result.Item, page)
	if err != nil {
		return nil, err
	}

	return page, nil
}

// PageUpdate can update any attribute of a stored page.
func pageUpdate(addr, expr string, values map[string]*dynamodb.AttributeValue) error {
	_, err := client.UpdateItem(&dynamodb.UpdateItemInput{
		TableName:           aws.String(pageTable),
		ConditionExpression: aws.String(fmt.Sprintf("attribute_exists(%v)", pageKey)),
		Key: map[string]*dynamodb.AttributeValue{
			pageKey: &dynamodb.AttributeValue{
				S: aws.String(addr),
			},
		},
		UpdateExpression:          aws.String(expr),
		ExpressionAttributeValues: values,
	})

	return err
}

// PageUpdatePassword updates a stored page's password hash.
func PageUpdatePassword(addr, pass string) error {
	return pageUpdate(addr,
		"SET password = :password",
		map[string]*dynamodb.AttributeValue{
			":password": {S: aws.String(pass)},
		},
	)
}

// PageUpdatePublished updates a stored page's published status.
func PageUpdatePublished(addr string, published bool) error {
	return pageUpdate(addr,
		"SET published = :published",
		map[string]*dynamodb.AttributeValue{
			":published": {BOOL: aws.Bool(published)},
		},
	)
}

// PageUpdateDocument updates a stored page's document.
func PageUpdateDocument(addr string, document string) error {
	return pageUpdate(addr,
		"SET document = :document",
		map[string]*dynamodb.AttributeValue{
			":document": {S: aws.String(document)},
		},
	)
}

// PageDelete removes a page from storage.
func PageDelete(addr string) error {
	_, err := client.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(pageTable),
		Key: map[string]*dynamodb.AttributeValue{
			pageKey: {
				S: aws.String(addr),
			},
		},
	})

	return err
}
