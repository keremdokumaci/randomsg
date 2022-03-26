package publisher

import "flag"

type SqsPublisher struct {
	Options AwsOptions
}

func NewSqsPublisher() SqsPublisher {
	return SqsPublisher{}
}

func (p SqsPublisher) Publish(message string) {

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

	p.Options.validate()
}
