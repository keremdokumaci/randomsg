package app

import "github.com/keremdokumaci/sqs-random-message-generator/app/validator"

type AwsOptions struct {
	AccessKey   string
	SecretKey   string
	Region      string `validate:"required"`
	QueueUrl    string
	SnsTopicArn string
}

func (a AwsOptions) validate() {
	if a.QueueUrl == "" && a.SnsTopicArn == "" {
		errorText("Queue Url or Sns Topic Arn must be specified !")
	}

	hasError, errorFields := validator.CustomValidator.ValidateStruct(a)
	if hasError {
		for _, field := range errorFields {
			switch field {
			case "Region":
				errorText("Region must be specified !")
			}
		}
	}
}
