package storage

import (
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// ErrFailedCondition when condition expression fails.
var ErrFailedCondition = errors.New("failed condition")

// ISO8601 is the layout used for fields representing dates.
const ISO8601 = "2006-01-02T15:04:05-0700"

const pageTable = "targetblank-pages"

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
// Returns "ErrFailedCondition" when page address is already in use.
func PageCreate(p *Page) error {
	page, err := dynamodbattribute.MarshalMap(p)
	if err != nil {
		return fmt.Errorf("marshal page: %v", err)
	}

	_, err = client.PutItem(&dynamodb.PutItemInput{
		TableName:           aws.String(pageTable),
		ConditionExpression: aws.String("attribute_not_exists(addr)"),
		Item:                page,
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == dynamodb.ErrCodeConditionalCheckFailedException {
				return ErrFailedCondition
			}
		}
		return fmt.Errorf("put item: %v", err)
	}

	return nil
}

// PageRead reads a page from storage.
// A null value pointer for page indicates the address was not found.
func PageRead(addr string) (*Page, error) {
	result, err := client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(pageTable),
		Key: map[string]*dynamodb.AttributeValue{
			"addr": {
				S: aws.String(addr),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("get item: %v", err)
	}
	if len(result.Item) == 0 {
		return nil, nil
	}

	page := &Page{}
	err = dynamodbattribute.UnmarshalMap(result.Item, page)
	if err != nil {
		return nil, fmt.Errorf("unmarshal page: %v", err)
	}

	return page, nil
}

// PageUpdate can update any attribute of a stored page.
func pageUpdate(addr, expr, cond string, values map[string]*dynamodb.AttributeValue) error {
	_, err := client.UpdateItem(&dynamodb.UpdateItemInput{
		TableName:           aws.String(pageTable),
		ConditionExpression: aws.String(cond),
		Key: map[string]*dynamodb.AttributeValue{
			"addr": {
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
	currentTime := time.Now().Format(ISO8601)
	return pageUpdate(addr,
		"SET password = :password, password_last_update = :password_last_update",
		"attribute_exists(addr)",
		map[string]*dynamodb.AttributeValue{
			":password":             {S: aws.String(pass)},
			":password_last_update": {S: aws.String(currentTime)},
		},
	)
}

// PageUpdateDocument updates a stored page's document.
// TODO special error about auth
func PageUpdateDocument(addr string, document string, authTimestamp *time.Time) error {
	err := pageUpdate(addr,
		"SET document = :document",
		"attribute_exists(addr) AND (attribute_not_exists(password_last_update) OR password_last_update < :auth_timestamp)",
		map[string]*dynamodb.AttributeValue{
			":document":       {S: aws.String(document)},
			":auth_timestamp": {S: aws.String(authTimestamp.Format(ISO8601))},
		},
	)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == dynamodb.ErrCodeConditionalCheckFailedException {
				return ErrFailedCondition
			}
		}
	}

	return err
}
