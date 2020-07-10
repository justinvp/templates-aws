package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/athena"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := athena.NewWorkgroup(ctx, "example", &athena.WorkgroupArgs{
			Configuration: &athena.WorkgroupConfigurationArgs{
				EnforceWorkgroupConfiguration:   pulumi.Bool(true),
				PublishCloudwatchMetricsEnabled: pulumi.Bool(true),
				ResultConfiguration: &athena.WorkgroupConfigurationResultConfigurationArgs{
					EncryptionConfiguration: &athena.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs{
						EncryptionOption: pulumi.String("SSE_KMS"),
						KmsKeyArn:        pulumi.Any(aws_kms_key.Example.Arn),
					},
					OutputLocation: pulumi.String("s3://{aws_s3_bucket.example.bucket}/output/"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

