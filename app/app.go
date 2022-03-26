package app

import (
	"bufio"
	"flag"
	"os"
)

type MessageOptions struct {
	FilePath       string
	SampleMessage  string
	MessageCount   int `validate:"required"`
	DelayInSeconds int
}

type AwsOptions struct {
	AccessKey   string
	SecretKey   string
	Region      string `validate:"required"`
	QueueUrl    string
	SnsTopicArn string
}

type Cli struct {
	MessageOptions MessageOptions
	AwsOptions     AwsOptions
}

func NewCli() Cli {
	filePath := flag.String("file", "", "message format yaml")
	accessKey := flag.String("accessKey", "", "access key for aws")
	secretKey := flag.String("secretKey", "", "secret key for aws")
	region := flag.String("region", "", "region for related service")
	queueUrl := flag.String("queue", "", "queue url")
	snsArn := flag.String("snsArn", "", "sns arn")
	flag.Parse()

	cli := Cli{
		MessageOptions: MessageOptions{
			FilePath: *filePath,
		},
		AwsOptions: AwsOptions{
			AccessKey:   *accessKey,
			SecretKey:   *secretKey,
			Region:      *region,
			QueueUrl:    *queueUrl,
			SnsTopicArn: *snsArn,
		},
	}

	return cli
}

func (cli Cli) Run() {
	if cli.MessageOptions.FilePath == "" {
		colorizedText(ColorGreen, "Insert a sample message.")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		cli.MessageOptions.SampleMessage = input.Text()
	}

	// call business func.
}
