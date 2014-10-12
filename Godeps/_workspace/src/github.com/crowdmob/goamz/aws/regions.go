package aws

var USGovWest = Region{
	"us-gov-west-1",
	"https://ec2.us-gov-west-1.amazonaws.com",
	"https://s3-fips-us-gov-west-1.amazonaws.com",
	"",
	true,
	true,
	"",
	"https://sns.us-gov-west-1.amazonaws.com",
	"https://sqs.us-gov-west-1.amazonaws.com",
	"https://iam.us-gov.amazonaws.com",
	"https://elasticloadbalancing.us-gov-west-1.amazonaws.com",
	"https://dynamodb.us-gov-west-1.amazonaws.com",
	ServiceInfo{"https://monitoring.us-gov-west-1.amazonaws.com", V2Signature},
	"https://autoscaling.us-gov-west-1.amazonaws.com",
	ServiceInfo{"https://rds.us-gov-west-1.amazonaws.com", V2Signature},
	"",
}

var USEast = Region{
	"us-east-1",
	"https://ec2.us-east-1.amazonaws.com",
	"https://s3.amazonaws.com",
	"",
	false,
	false,
	"https://sdb.amazonaws.com",
	"https://sns.us-east-1.amazonaws.com",
	"https://sqs.us-east-1.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.us-east-1.amazonaws.com",
	"https://dynamodb.us-east-1.amazonaws.com",
	ServiceInfo{"https://monitoring.us-east-1.amazonaws.com", V2Signature},
	"https://autoscaling.us-east-1.amazonaws.com",
	ServiceInfo{"https://rds.us-east-1.amazonaws.com", V2Signature},
	"https://kinesis.us-east-1.amazonaws.com",
}

var USWest = Region{
	"us-west-1",
	"https://ec2.us-west-1.amazonaws.com",
	"https://s3-us-west-1.amazonaws.com",
	"",
	true,
	true,
	"https://sdb.us-west-1.amazonaws.com",
	"https://sns.us-west-1.amazonaws.com",
	"https://sqs.us-west-1.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.us-west-1.amazonaws.com",
	"https://dynamodb.us-west-1.amazonaws.com",
	ServiceInfo{"https://monitoring.us-west-1.amazonaws.com", V2Signature},
	"https://autoscaling.us-west-1.amazonaws.com",
	ServiceInfo{"https://rds.us-west-1.amazonaws.com", V2Signature},
	"",
}

var USWest2 = Region{
	"us-west-2",
	"https://ec2.us-west-2.amazonaws.com",
	"https://s3-us-west-2.amazonaws.com",
	"",
	true,
	true,
	"https://sdb.us-west-2.amazonaws.com",
	"https://sns.us-west-2.amazonaws.com",
	"https://sqs.us-west-2.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.us-west-2.amazonaws.com",
	"https://dynamodb.us-west-2.amazonaws.com",
	ServiceInfo{"https://monitoring.us-west-2.amazonaws.com", V2Signature},
	"https://autoscaling.us-west-2.amazonaws.com",
	ServiceInfo{"https://rds.us-west-2.amazonaws.com", V2Signature},
	"https://kinesis.us-west-2.amazonaws.com",
}

var EUWest = Region{
	"eu-west-1",
	"https://ec2.eu-west-1.amazonaws.com",
	"https://s3-eu-west-1.amazonaws.com",
	"",
	true,
	true,
	"https://sdb.eu-west-1.amazonaws.com",
	"https://sns.eu-west-1.amazonaws.com",
	"https://sqs.eu-west-1.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.eu-west-1.amazonaws.com",
	"https://dynamodb.eu-west-1.amazonaws.com",
	ServiceInfo{"https://monitoring.eu-west-1.amazonaws.com", V2Signature},
	"https://autoscaling.eu-west-1.amazonaws.com",
	ServiceInfo{"https://rds.eu-west-1.amazonaws.com", V2Signature},
	"https://kinesis.eu-west-1.amazonaws.com",
}

var APSoutheast = Region{
	"ap-southeast-1",
	"https://ec2.ap-southeast-1.amazonaws.com",
	"https://s3-ap-southeast-1.amazonaws.com",
	"",
	true,
	true,
	"https://sdb.ap-southeast-1.amazonaws.com",
	"https://sns.ap-southeast-1.amazonaws.com",
	"https://sqs.ap-southeast-1.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.ap-southeast-1.amazonaws.com",
	"https://dynamodb.ap-southeast-1.amazonaws.com",
	ServiceInfo{"https://monitoring.ap-southeast-1.amazonaws.com", V2Signature},
	"https://autoscaling.ap-southeast-1.amazonaws.com",
	ServiceInfo{"https://rds.ap-southeast-1.amazonaws.com", V2Signature},
	"https://kinesis.ap-southeast-1.amazonaws.com",
}

var APSoutheast2 = Region{
	"ap-southeast-2",
	"https://ec2.ap-southeast-2.amazonaws.com",
	"https://s3-ap-southeast-2.amazonaws.com",
	"",
	true,
	true,
	"https://sdb.ap-southeast-2.amazonaws.com",
	"https://sns.ap-southeast-2.amazonaws.com",
	"https://sqs.ap-southeast-2.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.ap-southeast-2.amazonaws.com",
	"https://dynamodb.ap-southeast-2.amazonaws.com",
	ServiceInfo{"https://monitoring.ap-southeast-2.amazonaws.com", V2Signature},
	"https://autoscaling.ap-southeast-2.amazonaws.com",
	ServiceInfo{"https://rds.ap-southeast-2.amazonaws.com", V2Signature},
	"https://kinesis.ap-southeast-2.amazonaws.com",
}

var APNortheast = Region{
	"ap-northeast-1",
	"https://ec2.ap-northeast-1.amazonaws.com",
	"https://s3-ap-northeast-1.amazonaws.com",
	"",
	true,
	true,
	"https://sdb.ap-northeast-1.amazonaws.com",
	"https://sns.ap-northeast-1.amazonaws.com",
	"https://sqs.ap-northeast-1.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.ap-northeast-1.amazonaws.com",
	"https://dynamodb.ap-northeast-1.amazonaws.com",
	ServiceInfo{"https://monitoring.ap-northeast-1.amazonaws.com", V2Signature},
	"https://autoscaling.ap-northeast-1.amazonaws.com",
	ServiceInfo{"https://rds.ap-northeast-1.amazonaws.com", V2Signature},
	"https://kinesis.ap-northeast-1.amazonaws.com",
}

var SAEast = Region{
	"sa-east-1",
	"https://ec2.sa-east-1.amazonaws.com",
	"https://s3-sa-east-1.amazonaws.com",
	"",
	true,
	true,
	"https://sdb.sa-east-1.amazonaws.com",
	"https://sns.sa-east-1.amazonaws.com",
	"https://sqs.sa-east-1.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.sa-east-1.amazonaws.com",
	"https://dynamodb.sa-east-1.amazonaws.com",
	ServiceInfo{"https://monitoring.sa-east-1.amazonaws.com", V2Signature},
	"https://autoscaling.sa-east-1.amazonaws.com",
	ServiceInfo{"https://rds.sa-east-1.amazonaws.com", V2Signature},
	"",
}
