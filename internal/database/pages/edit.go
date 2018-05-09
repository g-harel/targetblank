package pages

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/g-harel/targetblank/internal/database"
)

// Edit modifies a page item's values.
func (p *Pages) Edit(addr string, i Item) error {
	if i.Email != "" {
		err := database.Validate(i.Email, "email")
		if err != nil {
			return err
		}
	}

	if i.Password != "" {
		err := database.Validate(i.Password, "password")
		if err != nil {
			return err
		}
	}

	i.Key = addr

	hashedPass, err := database.Hash(i.Password)
	if err != nil {
		return err
	}
	i.Password = hashedPass

	hashedEmail, err := database.Hash(i.Email)
	if err != nil {
		return err
	}
	i.Email = hashedEmail

	_, err = p.client.UpdateItem(&dynamodb.UpdateItemInput{
		TableName:                 aws.String(p.name),
		ConditionExpression:       aws.String("attribute_exists(addr)"),
		ExpressionAttributeValues: i.toUpdateMap(),
	})

	return err
}
