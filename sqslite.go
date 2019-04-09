package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/crowdmob/goamz/sqs"
)

// formatResp takes the response and formats it in JSON or XML.
func formatResp(format string, resp interface{}) ([]byte, error) {
	if format == "json" {
		return json.Marshal(resp)
	}
	return xml.Marshal(resp)
}

func main() {
	access := os.Getenv("AWS_ACCESS_KEY_ID")
	secret := os.Getenv("AWS_SECRET_ACCESS_KEY")

	cmd := flag.String("c", "r", "command (r=receive, s=send, d=delete)")
	qName := flag.String("q", "", "queue name")
	region := flag.String("re", "us-east-1", "region")
	format := flag.String("f", "xml", "response format (xml or json)")
	maxNumberOfMessages := flag.Int("mN", 1, "maximum messages")
	flag.Parse()

	// Check required environment variables
	if access == "" {
		panic("AWS_ACCESS_KEY_ID is undefined")
	}

	if secret == "" {
		panic("AWS_SECRET_ACCESS_KEY is undefined")
	}

	// If required flags are are not filled
	if *qName == "" || *region == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var b []byte
	var err error
	if *cmd == "s" || *cmd == "d" {
		b, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
	}

	const (
		QueueUrl    = "	https://sqs.us-east-1.amazonaws.com/385697007281/sync-md.fifo"
		Region      = "us-east-1"
		CredPath    = "~/.aws/credentials"
		CredProfile = "sync"
	)

	c, err := sqs.NewFrom(access, secret, *region)
	if err != nil {
		panic(err)
	}
	q, err := c.GetQueue(*qName)
	if err != nil {
		panic(err)
	}

	if *cmd == "s" {

		sess := session.New(&aws.Config{
			Region:      aws.String(Region),
			Credentials: credentials.NewSharedCredentials(CredPath, CredProfile),
			MaxRetries:  aws.Int(5),
		})

		svc := sqs.New(sess)

		// Send message
		send_params := &sqs.SendMessageInput{
			MessageBody: aws.String(string(b)), // Required
			QueueUrl:    aws.String(QueueUrl),  // Required
			MessageGroupId:    aws.String("1")  // Required
			//DelaySeconds: aws.Int64(3), // (optional) 傳進去的 message 延遲 n 秒才會被取出, 0 ~ 900s (15 minutes)
		}
		send_resp, err := svc.SendMessage(send_params)
		if err != nil {
			fmt.Println(err)
		}
		os.Stdout.Write("[Send message] \n%v \n\n", send_resp)

	} else if *cmd == "r" {
		resp, err := q.ReceiveMessage(*maxNumberOfMessages)
		if err != nil {
			panic(err)
		}
		b, err := formatResp(*format, resp)
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(b)
	} else if *cmd == "d" {
		m := &sqs.Message{ReceiptHandle: string(b)}
		resp, err := q.DeleteMessage(m)
		if err != nil {
			panic(err)
		}
		b, err := formatResp(*format, resp)
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(b)
	} else {
		flag.PrintDefaults()
		panic("Invalid command")
	}
}
