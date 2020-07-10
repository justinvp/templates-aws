package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := aws.GetBillingServiceAccount(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucket(ctx, "billingLogs", &s3.BucketArgs{
			Acl:    pulumi.String("private"),
			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Id\": \"Policy\",\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"s3:GetBucketAcl\", \"s3:GetBucketPolicy\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"arn:aws:s3:::my-billing-tf-test-bucket\",\n", "      \"Principal\": {\n", "        \"AWS\": [\n", "          \"", main.Arn, "\"\n", "        ]\n", "      }\n", "    },\n", "    {\n", "      \"Action\": [\n", "        \"s3:PutObject\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"arn:aws:s3:::my-billing-tf-test-bucket/*\",\n", "      \"Principal\": {\n", "        \"AWS\": [\n", "          \"", main.Arn, "\"\n", "        ]\n", "      }\n", "    }\n", "  ]\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

