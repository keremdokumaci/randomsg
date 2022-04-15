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
	// Initialize services, helpers
	validator.NewValidator()
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

	cli.MessageOptions.FilePath = *filePath
	cli.MessageOptions.MessageCount = *messageCount
	cli.MessageOptions.DelayInSeconds = *delayInSeconds

	hasValidationErr := cli.MessageOptions.Validate()
	if hasValidationErr {
		os.Exit(1)
	}

	var options interface{}

	switch publisher.PublisherType(*publisherType) {
	case publisher.AwsSQS:
		options = awsOptions
		break
	default:
		helper.ErrorText("Services except SQS and SNS are not supported yet !")
		os.Exit(1)
	}

	cli.Publisher = publisher.NewPublisher(publisher.PublisherType(*publisherType), options)

	return cli
}

func (cli Cli) Run() {

	// publish
	os.Exit(1)
}
