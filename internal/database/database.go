package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Client interface abstracts away the DynamoDB type.
type Client interface {
	GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}

// New creates a new database connection which implements Client.
func New() *dynamodb.DynamoDB {
	return dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))
}
