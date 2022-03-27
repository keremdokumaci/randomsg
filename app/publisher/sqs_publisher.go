package publisher

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/keremdokumaci/sqs-random-message-generator/app/helper"
)

type SqsPublisher struct {
	options AwsOptions
	client  *sqs.Client
}

func NewSqsPublisher() SqsPublisher {
	publisher := SqsPublisher{}

	if publisher.options.AccessKey != "" && publisher.options.SecretKey != "" {
		cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(publisher.options.Region), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(publisher.options.AccessKey, publisher.options.SecretKey, "")))
		if err != nil {
			helper.ErrorText(err.Error())
			os.Exit(1)
		}

		publisher.client = sqs.NewFromConfig(cfg)
	} else {
		helper.ColorizedText(helper.ColorYellow, "Looking for .aws folder to get credentials..")
		time.Sleep(time.Second * 2)

		cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(publisher.options.Region))
		if err != nil {
			helper.ErrorText(err.Error())
			os.Exit(1)
		}

		publisher.client = sqs.NewFromConfig(cfg)
	}

	return publisher
}

func (p SqsPublisher) Publish(message MessageOptions) {
	ctx, cancel := context.WithTimeout(context.Background(), 5) //todo: dynamic timeout
	defer cancel()

	for i := 0; i < message.MessageCount; i++ {
		output, err := p.client.SendMessage(ctx, &sqs.SendMessageInput{
			QueueUrl:          &p.options.QueueUrl,
			MessageAttributes: map[string]types.MessageAttributeValue{},
			MessageBody:       aws.String(message.MessageFormat),
		})

		if err != nil {
			helper.ErrorText(err.Error())
			os.Exit(1)
		}

		helper.ColorizedText(helper.ColorBlue, *output.MessageId+" has published.")

		if message.DelayInSeconds > 0 {
			time.Sleep(time.Duration(message.DelayInSeconds) * time.Second)
		}
	}

}

func (p SqsPublisher) SetCredentials(credentials interface{}) {
	p.options = credentials.(AwsOptions)
	hasValidationErr := p.options.validate()
	if hasValidationErr {
		os.Exit(1)
	}
}
