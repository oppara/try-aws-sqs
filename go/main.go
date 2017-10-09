package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joho/godotenv"
	"github.com/satori/go.uuid"
	"log"
	"os"
	"sync"
)

type sqsMessage struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

func usage() {
	fmt.Println("Usage:")
	fmt.Printf("\t$ %s (send|recv)\n", os.Args[0])
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func sendMessages(svc *sqs.SQS) error {
	for i := 0; i < 5; i++ {
		json, err := json.Marshal(sqsMessage{
			ID:   fmt.Sprintf("%s", uuid.NewV4()),
			Body: fmt.Sprintf("body %d", i),
		})
		if err != nil {
			return err
		}

		if err := sendMessage(svc, string(json)); err != nil {
			return err
		}
	}

	return nil
}

func sendMessage(svc *sqs.SQS, msg string) error {
	params := &sqs.SendMessageInput{
		MessageBody:    aws.String(msg),
		MessageGroupId: aws.String("golang"),
		QueueUrl:       aws.String(os.Getenv("QUEUE_URL")),
	}
	result, err := svc.SendMessage(params)
	if err != nil {
		return err
	}

	fmt.Println("SQS MessageId", *result.MessageId, msg)

	return nil
}

func receiveMessage(svc *sqs.SQS) error {
	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(os.Getenv("QUEUE_URL")),
		MaxNumberOfMessages: aws.Int64(10),
	}
	result, err := svc.ReceiveMessage(params)
	if err != nil {
		return err
	}

	fmt.Printf("messages count: %d\n", len(result.Messages))
	if len(result.Messages) == 0 {
		fmt.Println("empty queue.")
	}

	var wg sync.WaitGroup
	for _, m := range result.Messages {
		wg.Add(1)
		go func(msg *sqs.Message) {
			defer wg.Done()
			fmt.Println("SQS MessageId", *msg.MessageId, *msg.Body)
			if err := deleteMesage(svc, msg); err != nil {
				fmt.Println(err)
			}
		}(m)
	}
	wg.Wait()

	return nil
}

func deleteMesage(svc *sqs.SQS, msg *sqs.Message) error {
	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(os.Getenv("QUEUE_URL")),
		ReceiptHandle: aws.String(*msg.ReceiptHandle),
	}
	_, err := svc.DeleteMessage(params)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	if len(os.Args) < 2 || (os.Args[1] != "send" && os.Args[1] != "recv") {
		usage()
		os.Exit(1)
	}

	envLoad()

	sess := session.Must(session.NewSession())
	creds := credentials.NewStaticCredentials(os.Getenv("AWS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")
	svc := sqs.New(
		sess,
		aws.NewConfig().WithRegion("us-west-2").WithCredentials(creds),
	)

	action := os.Args[1]
	if action == "send" {
		if err := sendMessages(svc); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	if err := receiveMessage(svc); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)

}
