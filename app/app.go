package app

import (
	"encoding/json"
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

	publisherCreds := make(map[string]interface{})
	publisherCreds["accessKey"] = *accessKey
	publisherCreds["secretKey"] = *secretKey
	publisherCreds["region"] = *region
	publisherCreds["queueUrl"] = *queueUrl
	publisherCreds["topic"] = *topic

	marshaledCreds, err := json.Marshal(publisherCreds)
	if err != nil {
		helper.ErrorText(err.Error())
		os.Exit(1)
	}

	cli.Publisher = publisher.NewPublisher(publisher.PublisherType(*publisherType), string(marshaledCreds))
	cli.MessageOptions = publisher.MessageOptions{
		FilePath:       *filePath,
		MessageCount:   *messageCount,
		DelayInSeconds: *delayInSeconds,
	}
	return cli
}

func (cli Cli) Run() {

	// publish
	os.Exit(1)
}
