package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var appAWSConfig aws.Config

func setUpAWS() (err error) {
	// Configure AWS SDK to use LocalStack's endpoints and dummy credentials
	appAWSConfig, err = config.LoadDefaultConfig(
		context.TODO(),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           "http://localhost:4566",
				SigningRegion: "ap-south-1",
			}, nil
		})),
		config.WithCredentialsProvider(
			aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
				return aws.Credentials{
					AccessKeyID:     "dummy",
					SecretAccessKey: "dummy",
				}, nil
			}),
		),
	)
	if err != nil {
		log.Println("Error in aws config", err)
	}
	return
}
func main() {
	err := setUpAWS()
	if err != nil {
		return
	}
	svc := sqs.NewFromConfig(appAWSConfig)

	svc.CreateQueue(context.TODO(), &sqs.CreateQueueInput{
		QueueName: aws.String("AliceSQS"),
	})

	queueUrl, err := svc.GetQueueUrl(context.TODO(), &sqs.GetQueueUrlInput{
		QueueName: aws.String("BobSQS"),
	})

	if err != nil {
		log.Println("Error while getting queue url", err)
		return
	}
	log.Println("Queue url ", *queueUrl.QueueUrl)
	svc.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    queueUrl.QueueUrl,
		MessageBody: aws.String("Hello Dev here"),
	})

}
