package publisher

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/keremdokumaci/randomsg/app/helper"
)

type SnsPublisher struct {
	options AwsOptions
	client  *sns.Client
}

func NewSnsPublisher(options AwsOptions) SnsPublisher {
	publisher := SnsPublisher{}
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

		publisher.client = sns.NewFromConfig(cfg)
	} else {
		helper.ColorizedText(helper.ColorYellow, "Looking for .aws folder to get credentials..")
		time.Sleep(time.Second * 2)

		cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(publisher.options.Region))
		if err != nil {
			helper.ErrorText(err.Error())
		}

		publisher.client = sns.NewFromConfig(cfg)
	}

	return publisher
}

func (p SnsPublisher) Publish(message string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5)) //todo: dynamic timeout
	defer cancel()

	output, err := p.client.Publish(ctx, &sns.PublishInput{
		TopicArn:          &p.options.SnsTopicArn,
		MessageAttributes: map[string]types.MessageAttributeValue{},
		Message:           aws.String(message),
	})

	if err != nil {
		helper.ErrorText(err.Error())
	}

	helper.ColorizedText(helper.ColorGreen, *output.MessageId+" has published successfuly.")

}
