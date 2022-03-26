package app

type AwsOptions struct {
	AccessKey   string
	SecretKey   string
	Region      string `validate:"required"`
	QueueUrl    string
	SnsTopicArn string
}
