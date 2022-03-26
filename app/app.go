package app

import (
	"bufio"
	"flag"
	"os"

	"github.com/keremdokumaci/sqs-random-message-generator/app/helper"
	"github.com/keremdokumaci/sqs-random-message-generator/app/publisher"
	"github.com/keremdokumaci/sqs-random-message-generator/app/validator"
)

type Cli struct {
	MessageOptions publisher.MessageOptions
}

func NewCli() Cli {
	filePath := flag.String("file", "", "message format yaml")
	flag.Parse()

	cli := Cli{
		MessageOptions: publisher.MessageOptions{
			FilePath: *filePath,
		},
	}

	return cli
}

func (cli Cli) Run() {
	// Initialize services, helpers
	validator.NewValidator()

	if cli.MessageOptions.FilePath == "" {
		helper.ColorizedText(helper.ColorGreen, "Insert a sample message.")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		cli.MessageOptions.SampleMessage = input.Text()
	}

	// validation
	cli.MessageOptions.Validate()
}
