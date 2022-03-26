package app

import "github.com/keremdokumaci/sqs-random-message-generator/app/validator"

type MessageOptions struct {
	FilePath       string
	SampleMessage  string
	MessageCount   int `validate:"gt=0"`
	DelayInSeconds int
}

func (m MessageOptions) validate() {
	if m.FilePath == "" && m.SampleMessage == "" {
		errorText("Should specify file path or sample message !")
	}

	hasError, errorFields := validator.CustomValidator.ValidateStruct(m)
	if hasError {
		for _, field := range errorFields {
			switch field {
			case "MessageCount":
				errorText("Message Count should be greater than zero !")
			}
		}
	}
}
