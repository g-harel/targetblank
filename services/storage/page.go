package storage

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const pageTable = "targetblank-pages"
const pageKey = "addr"
const layoutISO8601 = "2006-01-02T15:04:05-0700"

// Page represents a DynamoDB page item.
type Page struct {
	Addr               string `json:"addr"` // pageKey
	Email              string `json:"email"`
	Password           string `json:"password"`
	PasswordLastUpdate string `json:"password_last_update"`
	Published          bool   `json:"published"`
	Document           string `json:"document"`
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
// TODO check if token was issued before last password update.
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
			pageKey: {
				S: aws.String(addr),
			},
		},
		UpdateExpression:          aws.String(expr),
		ExpressionAttributeValues: values,
	})

	return err
}

// PageUpdatePassword updates a stored page's password hash.
// TODO only allow password updates from email links.
func PageUpdatePassword(addr, pass string) error {
	currentTime := time.Now().Format(layoutISO8601)
	return pageUpdate(addr,
		"SET password = :password, password_last_update = :password_last_update",
		map[string]*dynamodb.AttributeValue{
			":password":             {S: aws.String(pass)},
			":password_last_update": {S: aws.String(currentTime)},
		},
	)
}

// PageUpdateDocument updates a stored page's document.
// TODO check if token was issued before last password update.
func PageUpdateDocument(addr string, document string) error {
	return pageUpdate(addr,
		"SET document = :document",
		map[string]*dynamodb.AttributeValue{
			":document": {S: aws.String(document)},
		},
	)
}

// PageDelete removes a page from storage.
// TODO check if token was issued before last password update.
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
