package storage

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var client = dynamodb.New(
	session.New(),
	aws.NewConfig().WithRegion(os.Getenv("AWS_REGION")),
)
