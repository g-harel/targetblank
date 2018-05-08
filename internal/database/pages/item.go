package pages

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Item represents a document stored in the page table.
type Item struct {
	Key      string `json:"addr"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Public   bool   `json:"public"`
	Temp     bool   `json:"temporary"`
	Page     string `json:"page"`

	// Flags used when using Item object as a source of updates.
	PublicHasBeenSet bool `json:"-"`
	TempHasBeenSet   bool `json:"-"`
}

func (i *Item) toAttributeValueMap() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"addr": {
			S: aws.String(i.Key),
		},
		"email": {
			S: aws.String(i.Email),
		},
		"password": {
			S: aws.String(i.Password),
		},
		"public": {
			BOOL: aws.Bool(i.Public),
		},
		"temporary": {
			BOOL: aws.Bool(i.Temp),
		},
		"page": {
			S: aws.String(i.Page),
		},
	}
}

// No fields can be empty string to avoid the default value problem and is consistent with DynamoDB's empty string policy.
// Key will never be included in the updated values.
func (i *Item) toAttributeValueUpdateMap() map[string]*dynamodb.AttributeValueUpdate {
	values := map[string]*dynamodb.AttributeValueUpdate{}

	if i.Email != "" {
		values["email"] = &dynamodb.AttributeValueUpdate{
			Value: &dynamodb.AttributeValue{
				S: aws.String(i.Email),
			},
		}
	}

	if i.Password != "" {
		values["password"] = &dynamodb.AttributeValueUpdate{
			Value: &dynamodb.AttributeValue{
				S: aws.String(i.Password),
			},
		}
	}

	if i.PublicHasBeenSet {
		values["public"] = &dynamodb.AttributeValueUpdate{
			Value: &dynamodb.AttributeValue{
				BOOL: aws.Bool(i.Public),
			},
		}
	}

	if i.TempHasBeenSet {
		values["temporary"] = &dynamodb.AttributeValueUpdate{
			Value: &dynamodb.AttributeValue{
				BOOL: aws.Bool(i.Temp),
			},
		}
	}

	if i.Page != "" {
		values["page"] = &dynamodb.AttributeValueUpdate{
			Value: &dynamodb.AttributeValue{
				S: aws.String(i.Page),
			},
		}
	}

	return values
}
