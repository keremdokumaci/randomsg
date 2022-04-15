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
		awsOptions := ConvertCredentials[AwsOptions](credentials)
		p = NewSqsPublisher()
		p.SetCredentials(awsOptions)
	default:
		helper.ErrorText("Couldn't find publisher !")
		os.Exit(1)
	}

	return p
}
