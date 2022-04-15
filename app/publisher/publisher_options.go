package publisher

import (
	"os"

	"github.com/keremdokumaci/sqs-random-message-generator/app/helper"
)

type IPublisherOptions interface {
	validate() bool
}

func ConvertCredentials[T IPublisherOptions](creds interface{}) T {
	converted, ok := creds.(T)
	if !ok {
		helper.ErrorText("couldn't get related parameters for your service ! Please check the parameters or run --help to see details !")
		os.Exit(1)
	}

	hasValidationErr := converted.validate()
	if hasValidationErr {
		os.Exit(1)
	}

	return converted
}
