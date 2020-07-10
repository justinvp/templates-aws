package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := s3.NewBucket(ctx, "bucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketPolicy(ctx, "bucketPolicy", &s3.BucketPolicyArgs{
			Bucket: bucket.ID(),
			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Id\": \"MYBUCKETPOLICY\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"IPAllow\",\n", "      \"Effect\": \"Deny\",\n", "      \"Principal\": \"*\",\n", "      \"Action\": \"s3:*\",\n", "      \"Resource\": \"arn:aws:s3:::my_tf_test_bucket/*\",\n", "      \"Condition\": {\n", "         \"IpAddress\": {\"aws:SourceIp\": \"8.8.8.8/32\"}\n", "      }\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

