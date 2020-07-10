package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudtrail"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := cloudtrail.GetServiceAccount(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			ForceDestroy: pulumi.Bool(true),
			Policy:       pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2008-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"Put bucket policy needed for trails\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"AWS\": \"", main.Arn, "\"\n", "      },\n", "      \"Action\": \"s3:PutObject\",\n", "      \"Resource\": \"arn:aws:s3:::tf-cloudtrail-logging-test-bucket/*\"\n", "    },\n", "    {\n", "      \"Sid\": \"Get bucket policy needed for trails\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"AWS\": \"", main.Arn, "\"\n", "      },\n", "      \"Action\": \"s3:GetBucketAcl\",\n", "      \"Resource\": \"arn:aws:s3:::tf-cloudtrail-logging-test-bucket\"\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

