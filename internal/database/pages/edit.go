package pages

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Edit modifies a page item's values.
func (p *Pages) Edit(addr string, i Item) error {
	i.Key = addr

	// TODO hash email, pass

	_, err := p.client.UpdateItem(&dynamodb.UpdateItemInput{
		TableName:           aws.String(p.name),
		ConditionExpression: aws.String("attribute_exists(addr)"),
		AttributeUpdates:    i.toAttributeValueUpdateMap(), // TODO refactor to UpdateExpression
	})

	return err
}
