package publisher

import (
	"flag"
	"os"
)

type SqsPublisher struct {
	Options AwsOptions
}

func NewSqsPublisher() SqsPublisher {
	publisher := SqsPublisher{}
	publisher.GetCredentials()
	return publisher
}

func (p SqsPublisher) Publish(message MessageOptions) {

}

func (p SqsPublisher) GetCredentials() {
	accessKey := flag.String("accessKey", "", "access key for aws")
	secretKey := flag.String("secretKey", "", "secret key for aws")
	region := flag.String("region", "", "region for related service")
	queueUrl := flag.String("queue", "", "queue url")
	flag.Parse()

	p.Options = AwsOptions{
		QueueUrl:  *queueUrl,
		Region:    *region,
		AccessKey: *accessKey,
		SecretKey: *secretKey,
	}

	hasValidationErr := p.Options.validate()
	if hasValidationErr {
		os.Exit(1)
	}
}
