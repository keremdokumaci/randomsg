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
	GetCredentials()
}

type Publisher struct {
	Publisher IPublisher
}

func NewPublisher(publisherType PublisherType) Publisher {
	p := Publisher{}

	switch publisherType {
	case AwsSQS:
		p.Publisher = NewSqsPublisher()
	default:
		helper.ErrorText("Couldn't find publisher !")
		os.Exit(1)
	}

	return p
}
