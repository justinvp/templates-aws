package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kinesis"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			Acl: pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		firehoseRole, err := iam.NewRole(ctx, "firehoseRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"firehose.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = kinesis.NewFirehoseDeliveryStream(ctx, "testStream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("s3"),
			S3Configuration: &kinesis.FirehoseDeliveryStreamS3ConfigurationArgs{
				BucketArn: bucket.Arn,
				RoleArn:   firehoseRole.Arn,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

