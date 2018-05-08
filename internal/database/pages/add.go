package pages

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/g-harel/targetblank/internal/database"
	"github.com/g-harel/targetblank/internal/page"
)

// Add creates a new page item.
func (p *Pages) Add(email string) (addr, pass string, err error) {
	err = database.Validate(email, "email")
	if err != nil {
		return "", "", err
	}

	page, parseErr := page.NewFromSpec("version 1\n===")
	if parseErr != nil {
		return addr, pass, parseErr
	}

	marshalledPageB, err := json.Marshal(page)
	if err != nil {
		return "", "", err
	}

	pass = database.RandString(16)

	hashedPass, err := database.Hash(pass)
	if err != nil {
		return "", "", err
	}

	hashedEmail, err := database.Hash(email)
	if err != nil {
		return "", "", err
	}

	item := &Item{
		Email:    hashedEmail,
		Password: hashedPass,
		Public:   false,
		Temp:     false,
		Page:     string(marshalledPageB),
	}

	input := &dynamodb.PutItemInput{
		TableName:           aws.String(p.name),
		ConditionExpression: aws.String("attribute_not_exists(addr)"),
	}

	for {
		addr = database.RandString(6)
		item.Key = addr
		input.Item = item.toAttributeValueMap()

		_, err := p.client.PutItem(input)
		if err != nil {
			if awsErr, ok := err.(awserr.Error); ok {
				if awsErr.Code() == dynamodb.ErrCodeConditionalCheckFailedException {
					continue
				}
			}
			return "", "", err
		}

		break
	}

	return addr, pass, nil
}
