package publisher

import (
	"encoding/json"

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
	}

	publisherOptions.validate()

	return publisherOptions
}
