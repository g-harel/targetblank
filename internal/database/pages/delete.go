package pages

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Delete removes an item from the table.
func (p *Pages) Delete(addr string) error {
	_, err := p.client.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(p.name),
		Key: map[string]*dynamodb.AttributeValue{
			"addr": {
				S: aws.String(addr),
			},
		},
	})

	return err
}
