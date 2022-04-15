package publisher

import (
	"github.com/keremdokumaci/sqs-random-message-generator/app/helper"
	"github.com/keremdokumaci/sqs-random-message-generator/app/validator"
)

type MessageOptions struct {
	FilePath       string `validate:"required"`
	MessageCount   int    `validate:"gt=0"`
	DelayInSeconds int
}

func (m MessageOptions) Validate() bool {
	hasError, errorFields := validator.CustomValidator.ValidateStruct(m)
	if hasError {
		for _, field := range errorFields {
			switch field {
			case "MessageCount":
				helper.ErrorText("Message Count must be greater than zero !")
				break
			case "FilePath":
				helper.ErrorText("File path must be defined !")
				break
			}
		}
	}

	return hasError
}
