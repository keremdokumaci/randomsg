package app

import (
	"flag"
	"os"

	"github.com/keremdokumaci/sqs-random-message-generator/app/helper"
	"github.com/keremdokumaci/sqs-random-message-generator/app/publisher"
	"github.com/keremdokumaci/sqs-random-message-generator/app/validator"
)

type Cli struct {
	MessageOptions publisher.MessageOptions
	Publisher      publisher.IPublisher
}

func NewCli() Cli {
	cli := Cli{}

	filePath := flag.String("file", "", "message format rules file")
	publisherType := flag.String("service", "", "service to push message (like sqs or sns)")
	messageCount := flag.Int("count", 0, "message count to push")
	delayInSeconds := flag.Int("delay", 0, "delay in seconds for each push")
	accessKey := flag.String("accessKey", "", "access key")
	secretKey := flag.String("secretKey", "", "secret key")
	region := flag.String("region", "", "region")
	queueUrl := flag.String("queue", "", "queue url")
	topic := flag.String("topic", "", "topic")

	flag.Parse()

	awsOptions := publisher.AwsOptions{
		QueueUrl:    *queueUrl,
		AccessKey:   *accessKey,
		SecretKey:   *secretKey,
		Region:      *region,
		SnsTopicArn: *topic,
	}
	messageOptions := publisher.MessageOptions{
		FilePath:       *filePath,
		MessageCount:   *messageCount,
		DelayInSeconds: *delayInSeconds,
	}

	var options interface{}

	switch publisher.PublisherType(*publisherType) {
	case publisher.AwsSQS:
		options = awsOptions
		break
	default:
		helper.ErrorText("other services are not supported yet !")
		os.Exit(1)
	}

	cli.Publisher = publisher.NewPublisher(publisher.PublisherType(*publisherType), options)
	cli.MessageOptions = messageOptions

	return cli
}

func (cli Cli) Run() {
	// Initialize services, helpers
	validator.NewValidator()

	// message option validation
	hasValidationErr := cli.MessageOptions.Validate()
	if hasValidationErr {
		os.Exit(1)
	}

	// publish
	os.Exit(1)
}
