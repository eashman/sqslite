sqslite
=======

Amazon SQS Lite/Simple CLI for getting messages in and out of Amazon SQS. Useful for creating polyglot jobs/workers in BASH/MS-DOS/Command line.

### Send a message

```
export AWS_ACCESS_KEY_ID=whatever
export AWS_SECRET_ACCESS_KEY=whatever
echo "message" | sqslite -q queue-name -c s
```

### Receive a message

```
sqslite -q queue-name
```

### Delete a message
```
echo "ReceiptHandlerId" | sqslite -q queue-name -c d
```

### Full Example

See examples folder for the work script. You can run the script like this:

```
QUEUE_NAME=sqslite QUEUE_REGION=us-east-1 ./work job
```

### Installation 

sqslite was written in Go so you can compile it on any platform Go supports (windows, linux or OS X)! To get started, install go and then go the following:

```
go get github.com/crowdmob/goamz/sqs

# Make sure you're in the cloned sqs directory
go build
```

Install Go: https://golang.org/doc/install

You can also run go install which will install the binary universally if you have it setup properly. See Go manual for instructions.
