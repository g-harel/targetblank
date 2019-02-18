package secrets

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

const parameterName = "targetblank-key"

// Key stores the decrypted key after it has been fetched.
var key = ""

// Key retrieves the decrypted application key secret.
func Key() (string, error) {
	if key != "" {
		return key, nil
	}

	res, err := client.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(parameterName),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "", err
	}

	key = res.String()

	return key, nil
}
