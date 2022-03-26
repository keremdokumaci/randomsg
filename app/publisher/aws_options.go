package publisher

import (
	"os"

	"github.com/keremdokumaci/sqs-random-message-generator/app/helper"
	"github.com/keremdokumaci/sqs-random-message-generator/app/validator"
)

type AwsOptions struct {
	AccessKey   string
	SecretKey   string
	Region      string `validate:"required"`
	QueueUrl    string
	SnsTopicArn string
}

func (a AwsOptions) validate() {
	if a.QueueUrl == "" && a.SnsTopicArn == "" {
		helper.ErrorText("Queue Url or Sns Topic Arn must be specified !")
	}

	hasError, errorFields := validator.CustomValidator.ValidateStruct(a)
	if hasError {
		for _, field := range errorFields {
			switch field {
			case "Region":
				helper.ErrorText("Region must be specified !")
				os.Exit(1)
			}
		}
	}
}
