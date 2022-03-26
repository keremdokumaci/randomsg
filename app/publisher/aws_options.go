package publisher

import (
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

func (a AwsOptions) validate() bool {
	if a.QueueUrl == "" && a.SnsTopicArn == "" {
		helper.ErrorText("Queue Url or Sns Topic Arn must be specified !")
		return true
	}

	hasError, errorFields := validator.CustomValidator.ValidateStruct(a)
	if hasError {
		for _, field := range errorFields {
			switch field {
			case "Region":
				helper.ErrorText("Region must be specified !")
			}
		}
	}

	return hasError
}
