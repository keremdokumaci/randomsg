package publisher

import (
	"github.com/keremdokumaci/sqs-random-message-generator/app/helper"
	"github.com/keremdokumaci/sqs-random-message-generator/app/validator"
)

type MessageOptions struct {
	FilePath       string
	MessageFormat  string
	MessageCount   int `validate:"gt=0"`
	DelayInSeconds int
}

func (m MessageOptions) Validate() bool {
	if m.FilePath == "" && m.MessageFormat == "" {
		helper.ErrorText("Must specify file path or message format!")
		return true
	}

	hasError, errorFields := validator.CustomValidator.ValidateStruct(m)
	if hasError {
		for _, field := range errorFields {
			switch field {
			case "MessageCount":
				helper.ErrorText("Message Count must be greater than zero !")
			}
		}
	}

	return hasError
}
