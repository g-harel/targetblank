package pages

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/g-harel/targetblank/internal/database"
)

// Create adds a new page item.
// Final key is added to the referenced Item object.
func (p *Pages) Create(item *Item) error {
	input := &dynamodb.PutItemInput{
		TableName:           aws.String(p.name),
		ConditionExpression: aws.String("attribute_not_exists(addr)"),
	}

	for {
		item.Key = database.RandString(6)
		input.Item = item.toCreateMap()

		_, err := p.client.PutItem(input)
		if err != nil {
			if awsErr, ok := err.(awserr.Error); ok {
				if awsErr.Code() == dynamodb.ErrCodeConditionalCheckFailedException {
					continue
				}
			}
			return err
		}

		break
	}

	return nil
}
