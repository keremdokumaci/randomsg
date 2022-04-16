package publisher

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/keremdokumaci/sqs-random-message-generator/app/helper"
)

type IPublisherOptions interface {
	validate() bool
}

func ConvertCredentials[T IPublisherOptions](creds string) T {
	var publisherOptions T
	err := json.Unmarshal([]byte(creds), &publisherOptions)

	if err != nil {
		helper.ErrorText(err.Error())
		os.Exit(1)
	}
	fmt.Println(publisherOptions)
	hasValidationErr := publisherOptions.validate()
	if hasValidationErr {
		os.Exit(1)
	}

	return publisherOptions
}
