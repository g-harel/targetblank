package database

import "github.com/aws/aws-sdk-go/service/dynamodb"

var _ Client = &dynamodb.DynamoDB{}
