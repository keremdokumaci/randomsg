package publisher

import (
	"github.com/keremdokumaci/random-message-generator/app/helper"
	"github.com/keremdokumaci/random-message-generator/app/validator"
)

type AwsOptions struct {
	AccessKey   string `json:"access_key"`
	SecretKey   string `json:"secret_key"`
	Region      string `validate:"required"`
	QueueUrl    string `json:"queue"`
	SnsTopicArn string `json:"sns"`
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
