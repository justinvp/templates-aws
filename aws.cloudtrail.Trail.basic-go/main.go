package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudtrail"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		foo, err := s3.NewBucket(ctx, "foo", &s3.BucketArgs{
			ForceDestroy: pulumi.Bool(true),
			Policy:       pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2012-10-17\",\n", "    \"Statement\": [\n", "        {\n", "            \"Sid\": \"AWSCloudTrailAclCheck\",\n", "            \"Effect\": \"Allow\",\n", "            \"Principal\": {\n", "              \"Service\": \"cloudtrail.amazonaws.com\"\n", "            },\n", "            \"Action\": \"s3:GetBucketAcl\",\n", "            \"Resource\": \"arn:aws:s3:::tf-test-trail\"\n", "        },\n", "        {\n", "            \"Sid\": \"AWSCloudTrailWrite\",\n", "            \"Effect\": \"Allow\",\n", "            \"Principal\": {\n", "              \"Service\": \"cloudtrail.amazonaws.com\"\n", "            },\n", "            \"Action\": \"s3:PutObject\",\n", "            \"Resource\": \"arn:aws:s3:::tf-test-trail/prefix/AWSLogs/", current.AccountId, "/*\",\n", "            \"Condition\": {\n", "                \"StringEquals\": {\n", "                    \"s3:x-amz-acl\": \"bucket-owner-full-control\"\n", "                }\n", "            }\n", "        }\n", "    ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		_, err = cloudtrail.NewTrail(ctx, "foobar", &cloudtrail.TrailArgs{
			IncludeGlobalServiceEvents: pulumi.Bool(false),
			S3BucketName:               foo.ID(),
			S3KeyPrefix:                pulumi.String("prefix"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

