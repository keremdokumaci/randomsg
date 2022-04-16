package publisher

import (
	"context"
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

func NewSqsPublisher(options AwsOptions) SqsPublisher {
	publisher := SqsPublisher{}
	publisher.options = options

	hasErr := publisher.options.validate()
	if hasErr {
		helper.ErrorText("Check AWS related parameters in your file.")
	}

	if publisher.options.AccessKey != "" && publisher.options.SecretKey != "" {
		cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(publisher.options.Region), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(publisher.options.AccessKey, publisher.options.SecretKey, "")))
		if err != nil {
			helper.ErrorText(err.Error())
		}

		publisher.client = sqs.NewFromConfig(cfg)
	} else {
		helper.ColorizedText(helper.ColorYellow, "Looking for .aws folder to get credentials..")
		time.Sleep(time.Second * 2)

		cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(publisher.options.Region))
		if err != nil {
			helper.ErrorText(err.Error())
		}

		publisher.client = sqs.NewFromConfig(cfg)
	}

	return publisher
}

func (p SqsPublisher) Publish(message string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5)) //todo: dynamic timeout
	defer cancel()

	output, err := p.client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:          &p.options.QueueUrl,
		MessageAttributes: map[string]types.MessageAttributeValue{},
		MessageBody:       aws.String(message),
	})

	if err != nil {
		helper.ErrorText(err.Error())
	}

	helper.ColorizedText(helper.ColorGreen, *output.MessageId+" has published successfuly.")

}
