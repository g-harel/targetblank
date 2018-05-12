package pages

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/page"
)

// Create adds a new page item.
func (p *Pages) Create(email, pass string) (addr string, err error) {
	err = database.Validate(email, "email")
	if err != nil {
		return "", err
	}

	page, parseErr := page.NewFromSpec("version 1\n===")
	if parseErr != nil {
		return addr, parseErr
	}

	marshalledPageB, err := json.Marshal(page)
	if err != nil {
		return "", err
	}

	hashedPass, err := database.Hash(pass)
	if err != nil {
		return "", err
	}

	hashedEmail, err := database.Hash(email)
	if err != nil {
		return "", err
	}

	item := &Item{
		Email:     hashedEmail,
		Password:  hashedPass,
		Published: false,
		TempPass:  true,
		Page:      string(marshalledPageB),
	}

	input := &dynamodb.PutItemInput{
		TableName:           aws.String(p.name),
		ConditionExpression: aws.String("attribute_not_exists(addr)"),
	}

	for {
		addr = database.RandString(6)
		item.Key = addr
		input.Item = item.toCreateMap()

		_, err := p.client.PutItem(input)
		if err != nil {
			if awsErr, ok := err.(awserr.Error); ok {
				if awsErr.Code() == dynamodb.ErrCodeConditionalCheckFailedException {
					continue
				}
			}
			return "", err
		}

		break
	}

	return addr, nil
}
