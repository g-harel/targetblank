package pages

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/g-harel/targetblank/internal/database"
)

// Change modifies an item's values.
func (p *Pages) Change(addr string, i Item) error {
	if i.Email != "" {
		err := database.Validate(i.Email, "email")
		if err != nil {
			return err
		}

		hashedEmail, err := database.Hash(i.Email)
		if err != nil {
			return err
		}
		i.Email = hashedEmail
	}

	if i.Password != "" {
		err := database.Validate(i.Password, "password")
		if err != nil {
			return err
		}

		hashedPass, err := database.Hash(i.Password)
		if err != nil {
			return err
		}
		i.Password = hashedPass
	}

	expression, values := i.toUpdateExpression()
	_, err := p.client.UpdateItem(&dynamodb.UpdateItemInput{
		TableName:           aws.String(p.name),
		ConditionExpression: aws.String("attribute_exists(addr)"),
		Key: map[string]*dynamodb.AttributeValue{
			"addr": &dynamodb.AttributeValue{
				S: aws.String(addr),
			},
		},
		ExpressionAttributeValues: values,
		UpdateExpression:          aws.String(expression),
	})

	return err
}
