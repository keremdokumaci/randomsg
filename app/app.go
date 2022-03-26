package app

import (
	"bufio"
	"flag"
	"os"
	"time"

	"github.com/keremdokumaci/sqs-random-message-generator/app/helper"
	"github.com/keremdokumaci/sqs-random-message-generator/app/publisher"
	"github.com/keremdokumaci/sqs-random-message-generator/app/validator"
)

type Cli struct {
	MessageOptions publisher.MessageOptions
	PublisherType  publisher.PublisherType
}

func NewCli() Cli {
	var sampleMessage string
	filePath := flag.String("file", "", "message option file")
	publisherType := flag.String("service", "", "service to push message (like sqs or sns)")
	flag.Parse()

	if *filePath == "" {
		helper.ColorizedText(helper.ColorGreen, "Insert a sample message.")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		sampleMessage = input.Text()
	}

	cli := Cli{
		MessageOptions: publisher.MessageOptions{
			FilePath:      *filePath,
			SampleMessage: sampleMessage,
		},
		PublisherType: publisher.PublisherType(*publisherType),
	}

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
	publisher.NewPublisher(cli.PublisherType).Publisher.Publish(cli.MessageOptions)

	time.Sleep(2)
	os.Exit(1)
}
