package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// const (
// 	QueueUrl = "https://sqs.us-east-1.amazonaws.com/385697007281/sync-md.fifo"
// )

func main() {
	access := os.Getenv("AWS_ACCESS_KEY_ID")
	secret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	qUrl := flag.String("q", "", "queue url")

	flag.Parse()

	// Check required environment variables
	if access == "" {
		panic("AWS_ACCESS_KEY_ID is undefined")
	}

	if secret == "" {
		panic("AWS_SECRET_ACCESS_KEY is undefined")
	}

	if region == "" {
		panic("AWS_REGION is undefined")
	}

	var queue_url string
	if *qUrl == "" {
		queue_url = string("https://sqs.us-east-1.amazonaws.com/385697007281/sync-md.fifo")
	} else {
		queue_url = *qUrl
	}

	fmt.Println("1")

	// B is the filename
	var b []byte
	var err error
	b, err = ioutil.ReadAll(os.Stdin)
	fmt.Println("2")

	if err != nil {
		panic(err)
	}

	fmt.Println("3")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(access, secret, ""),
		MaxRetries:  aws.Int(2),
	})

	svc := sqs.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
	//svc := sqs.New(sess)

	// Send message
	send_params := &sqs.SendMessageInput{
		MessageBody:    aws.String(string(b)),
		QueueUrl:       aws.String(queue_url),
		MessageGroupId: aws.String("1"),
	}
	send_resp, err := svc.SendMessage(send_params)
	if err != nil {
		fmt.Println("[Error message]\n", err)
	}
	fmt.Println("[Sent message]\n", send_resp)
}
