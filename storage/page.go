package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const PAGE_TABLE = "targetblank-tables"

type Page struct {
	Key       string `json:"addr"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Published bool   `json:"published"`
	Data      string `json:"data"`
}

func PageCreate(p *Page) (conflict bool, err error) {
	item, err := dynamodbattribute.MarshalMap(p)
	if err != nil {
		return false, err
	}

	_, err = client.PutItem(&dynamodb.PutItemInput{
		TableName:           aws.String(PAGE_TABLE),
		ConditionExpression: aws.String("attribute_not_exists(addr)"),
		Item:                item,
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			conflict = awsErr.Code() == dynamodb.ErrCodeConditionalCheckFailedException
		}
		return conflict, err
	}

	return false, nil
}

func PageRead(addr string) (*Page, error) {
	result, err := client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(PAGE_TABLE),
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

	item := &Page{}
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func pageUpdate(addr, expr string, values map[string]*dynamodb.AttributeValue) error {
	_, err := client.UpdateItem(&dynamodb.UpdateItemInput{
		TableName:           aws.String(PAGE_TABLE),
		ConditionExpression: aws.String("attribute_exists(addr)"),
		Key: map[string]*dynamodb.AttributeValue{
			"addr": &dynamodb.AttributeValue{
				S: aws.String(addr),
			},
		},
		UpdateExpression:          aws.String(expr),
		ExpressionAttributeValues: values,
	})

	return err
}

func PageUpdatePassword(addr, pass string) error {
	return pageUpdate(addr,
		"SET password = :password",
		map[string]*dynamodb.AttributeValue{
			":password": {S: aws.String(pass)},
		},
	)
}

func PageUpdatePublished(addr string, published bool) error {
	return pageUpdate(addr,
		"SET published = :published",
		map[string]*dynamodb.AttributeValue{
			":published": {BOOL: aws.Bool(published)},
		},
	)
}

func PageUpdateData(addr string, data string) error {
	return pageUpdate(addr,
		"SET data = :data",
		map[string]*dynamodb.AttributeValue{
			":data": {S: aws.String(data)},
		},
	)
}

func PageDelete(addr string) error {
	_, err := client.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(PAGE_TABLE),
		Key: map[string]*dynamodb.AttributeValue{
			"addr": {
				S: aws.String(addr),
			},
		},
	})

	return err
}
