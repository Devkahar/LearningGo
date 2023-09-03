package main

import (
	"context"
	"log"
	"time"

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

	queue, err := svc.CreateQueue(context.TODO(), &sqs.CreateQueueInput{
		QueueName: aws.String("BobSQS"),
	})

	if err != nil {
		log.Println("Error while creating queue ", err)
	}

	queueUrl := queue.QueueUrl
	reciveInput := &sqs.ReceiveMessageInput{
		QueueUrl:            queueUrl,
		MaxNumberOfMessages: 10,
	}
	for {
		messages, _ := svc.ReceiveMessage(context.TODO(), reciveInput)
		for _, message := range messages.Messages {
			log.Println("Message recived ", *message.Body)
		}
		time.Sleep(time.Second * 2)
	}

}
