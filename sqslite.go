package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/jessevdk/go-flags"
)

func main() {

	var opts struct {
		File    string `short:"f" long:"file" description:"A file" value-name:"FILE"`
		Queue   string `short:"q" long:"queue" description:"Queue URL" value-name:"QUEUE" default:"https://sqs.us-east-1.amazonaws.com/385697007281/sync-md.fifo"`
		Config  string `short:"c" long:"config" description:"Config file location" value-name:"CONFIG"`
		Profile string `short:"p" long:"profile" description:"Config Profile name" value-name:"PROFILE"`
	}

	flags.Parse(&opts)
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	fmt.Println(currentTime, " - FILE - name: ", strings.TrimSpace(opts.File))

	creds := credentials.NewStaticCredentials("", "", "")

	if opts.Config == "" {
		creds = credentials.NewEnvCredentials()
	} else {
		creds = credentials.NewSharedCredentials(opts.Config, opts.Profile)
	}

	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		MaxRetries:  aws.Int(2),
	})

	//svc := sqs.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
	svc := sqs.New(sess)

	// Send message
	send_params := &sqs.SendMessageInput{
		MessageBody:    aws.String(strings.TrimSpace(opts.File)),
		QueueUrl:       aws.String(opts.Queue),
		MessageGroupId: aws.String("1"),
	}
	send_resp, err := svc.SendMessage(send_params)
	if err != nil {
		fmt.Println(currentTime, " - ERROR - ", err)
	}
	fmt.Println(currentTime, " - SENT - msgid: ", send_resp.MessageId)
}
