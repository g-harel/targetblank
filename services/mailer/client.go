package mailer

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

var client = ses.New(
	session.New(),
	aws.NewConfig().WithRegion(os.Getenv("AWS_REGION")),
)
