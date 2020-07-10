package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/emr"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := emr.NewSecurityConfiguration(ctx, "foo", &emr.SecurityConfigurationArgs{
			Configuration: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"EncryptionConfiguration\": {\n", "    \"AtRestEncryptionConfiguration\": {\n", "      \"S3EncryptionConfiguration\": {\n", "        \"EncryptionMode\": \"SSE-S3\"\n", "      },\n", "      \"LocalDiskEncryptionConfiguration\": {\n", "        \"EncryptionKeyProviderType\": \"AwsKms\",\n", "        \"AwsKmsKey\": \"arn:aws:kms:us-west-2:187416307283:alias/tf_emr_test_key\"\n", "      }\n", "    },\n", "    \"EnableInTransitEncryption\": false,\n", "    \"EnableAtRestEncryption\": true\n", "  }\n", "}\n", "\n")),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

