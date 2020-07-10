package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kinesis.NewFirehoseDeliveryStream(ctx, "testStream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("splunk"),
			S3Configuration: &kinesis.FirehoseDeliveryStreamS3ConfigurationArgs{
				BucketArn:         pulumi.Any(aws_s3_bucket.Bucket.Arn),
				BufferInterval:    pulumi.Int(400),
				BufferSize:        pulumi.Int(10),
				CompressionFormat: pulumi.String("GZIP"),
				RoleArn:           pulumi.Any(aws_iam_role.Firehose.Arn),
			},
			SplunkConfiguration: &kinesis.FirehoseDeliveryStreamSplunkConfigurationArgs{
				HecAcknowledgmentTimeout: pulumi.Int(600),
				HecEndpoint:              pulumi.String("https://http-inputs-mydomain.splunkcloud.com:443"),
				HecEndpointType:          pulumi.String("Event"),
				HecToken:                 pulumi.String("51D4DA16-C61B-4F5F-8EC7-ED4301342A4A"),
				S3BackupMode:             pulumi.String("FailedEventsOnly"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

