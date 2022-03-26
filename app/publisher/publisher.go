package publisher

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
	publisher IPublisher
}

func NewPublisher(publisherType PublisherType) Publisher {
	p := Publisher{}

	switch publisherType {
	case AwsSQS:
		p.publisher = NewSqsPublisher()
	default:
		panic("Couldn't find publisher !")
	}

	return p
}
