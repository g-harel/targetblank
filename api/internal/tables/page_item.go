package tables

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// PageItem represents a document stored in the page table.
type PageItem struct {
	Key       string `json:"addr"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Published bool   `json:"published"`
	TempPass  bool   `json:"temp_pass"`
	Page      string `json:"page"`

	// Flags used when using Item object as a source of updates.
	PublishedHasBeenSetForUpdateExpression bool `json:"-"`
	TempPassHasBeenSetForUpdateExpression  bool `json:"-"`
}

func (i *PageItem) toCreateMap() map[string]*dynamodb.AttributeValue {
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
		"published": {
			BOOL: aws.Bool(i.Published),
		},
		"temp_pass": {
			BOOL: aws.Bool(i.TempPass),
		},
		"page": {
			S: aws.String(i.Page),
		},
	}
}

// No fields can be empty string to avoid the default value problem and is consistent with DynamoDB's no empty string policy.
// Key will never be included in the updated values.
func (i *PageItem) toUpdateExpression() (string, map[string]*dynamodb.AttributeValue) {
	expression := "SET "
	values := map[string]*dynamodb.AttributeValue{}

	if i.Email != "" {
		values[":email"] = &dynamodb.AttributeValue{
			S: aws.String(i.Email),
		}
		expression += "email = :email,"
	}

	if i.Password != "" {
		values[":password"] = &dynamodb.AttributeValue{
			S: aws.String(i.Password),
		}
		expression += "password = :password,"
	}

	if i.PublishedHasBeenSetForUpdateExpression {
		values[":published"] = &dynamodb.AttributeValue{
			BOOL: aws.Bool(i.Published),
		}
		expression += "published = :published,"
	}

	if i.TempPassHasBeenSetForUpdateExpression {
		values[":temp_pass"] = &dynamodb.AttributeValue{
			BOOL: aws.Bool(i.TempPass),
		}
		expression += "temp_pass = :temp_pass,"
	}

	if i.Page != "" {
		values[":page"] = &dynamodb.AttributeValue{
			S: aws.String(i.Page),
		}
		expression += "page = :page,"
	}

	return strings.TrimRight(expression, ","), values
}
