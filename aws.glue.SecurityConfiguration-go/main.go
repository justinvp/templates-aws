package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewSecurityConfiguration(ctx, "example", &glue.SecurityConfigurationArgs{
			EncryptionConfiguration: &glue.SecurityConfigurationEncryptionConfigurationArgs{
				CloudwatchEncryption: &glue.SecurityConfigurationEncryptionConfigurationCloudwatchEncryptionArgs{
					CloudwatchEncryptionMode: pulumi.String("DISABLED"),
				},
				JobBookmarksEncryption: &glue.SecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionArgs{
					JobBookmarksEncryptionMode: pulumi.String("DISABLED"),
				},
				S3Encryption: &glue.SecurityConfigurationEncryptionConfigurationS3EncryptionArgs{
					KmsKeyArn:        pulumi.Any(data.Aws_kms_key.Example.Arn),
					S3EncryptionMode: pulumi.String("SSE-KMS"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

