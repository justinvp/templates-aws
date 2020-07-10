package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		hogeBucket, err := s3.NewBucket(ctx, "hogeBucket", &s3.BucketArgs{
			Region: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketPolicy(ctx, "hogeBucketPolicy", &s3.BucketPolicyArgs{
			Bucket: hogeBucket.Bucket,
			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2012-10-17\",\n", "    \"Statement\": [\n", "        {\n", "            \"Sid\": \"SSMBucketPermissionsCheck\",\n", "            \"Effect\": \"Allow\",\n", "            \"Principal\": {\n", "                \"Service\": \"ssm.amazonaws.com\"\n", "            },\n", "            \"Action\": \"s3:GetBucketAcl\",\n", "            \"Resource\": \"arn:aws:s3:::tf-test-bucket-1234\"\n", "        },\n", "        {\n", "            \"Sid\": \" SSMBucketDelivery\",\n", "            \"Effect\": \"Allow\",\n", "            \"Principal\": {\n", "                \"Service\": \"ssm.amazonaws.com\"\n", "            },\n", "            \"Action\": \"s3:PutObject\",\n", "            \"Resource\": [\"arn:aws:s3:::tf-test-bucket-1234/*\"],\n", "            \"Condition\": {\n", "                \"StringEquals\": {\n", "                    \"s3:x-amz-acl\": \"bucket-owner-full-control\"\n", "                }\n", "            }\n", "        }\n", "    ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = ssm.NewResourceDataSync(ctx, "foo", &ssm.ResourceDataSyncArgs{
			S3Destination: &ssm.ResourceDataSyncS3DestinationArgs{
				BucketName: hogeBucket.Bucket,
				Region:     hogeBucket.Region,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

