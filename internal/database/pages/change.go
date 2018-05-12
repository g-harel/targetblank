package pages

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Change modifies an item's values.
func (p *Pages) Change(addr string, i *Item) error {
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
