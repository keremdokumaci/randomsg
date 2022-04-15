package publisher

import (
	"os"

	"github.com/keremdokumaci/sqs-random-message-generator/app/helper"
)

type PublisherType string

const (
	AwsSQS PublisherType = "sqs"
	AwsSNS               = "sns"
)

type IPublisher interface {
	Publish(message MessageOptions)
	SetCredentials(credentials interface{})
}

func NewPublisher(publisherType PublisherType, credentials interface{}) IPublisher {
	var p IPublisher

	switch publisherType {
	case AwsSQS:
		awsCreds, ok := credentials.(AwsOptions)
		if !ok {
			helper.ErrorText("couldn't get AWS related parameters ! Please check the parameters or run --help to see details !")
			os.Exit(1)
		}
		hasValidationErr := awsCreds.validate()
		if hasValidationErr {
			os.Exit(1)
		}

		p = NewSqsPublisher()
		p.SetCredentials(credentials)
	default:
		helper.ErrorText("Couldn't find publisher !")
		os.Exit(1)
	}

	return p
}
