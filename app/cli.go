package app

import (
	"flag"
	"os"
	"path/filepath"

	filereader "github.com/keremdokumaci/sqs-random-message-generator/app/file_reader"
	"github.com/keremdokumaci/sqs-random-message-generator/app/publisher"
	"github.com/keremdokumaci/sqs-random-message-generator/app/validator"
)

type Cli struct {
	CliOptions CliOptions
	Publisher  publisher.IPublisher
	FileReader filereader.IReader
}

func NewCli() Cli {
	// Initialize services, helpers
	validator.NewValidator()
	cli := Cli{}

	filePath := flag.String("file", "", "message format rules file")
	publisherType := flag.String("service", "", "service to push message (like sqs or sns)")
	messageCount := flag.Int("count", 0, "message count to push")
	delayInSeconds := flag.Int("delay", 0, "delay in seconds for each push")
	flag.Parse()

	cli.CliOptions = CliOptions{
		FilePath:       *filePath,
		MessageCount:   *messageCount,
		DelayInSeconds: *delayInSeconds,
		ServiceType:    *publisherType,
	}
	hasError := cli.CliOptions.Validate()
	if hasError {
		os.Exit(1)
	}

	cli.FileReader = filereader.NewFileReader(filepath.Ext(*filePath))
	return cli
}

func (cli Cli) Run() {
	// cli.Publisher = publisher.NewPublisher()

	// publish
	os.Exit(1)
}
