package publisher

import (
	"github.com/keremdokumaci/sqs-random-message-generator/app/helper"
)

type PublisherType string

const (
	AwsSQS PublisherType = "sqs"
	AwsSNS               = "sns"
)

type IPublisher interface {
	Publish(message string)
}

func NewPublisher(publisherType PublisherType, credentials string) IPublisher {
	var p IPublisher

	switch publisherType {
	case AwsSQS:
		awsOptions := ConvertCredentials[AwsOptions](credentials)
		p = NewSqsPublisher(awsOptions)
	default:
		helper.ErrorText("Couldn't find publisher !")
	}

	return p
}
