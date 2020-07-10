package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		mykey, err := kms.NewKey(ctx, "mykey", &kms.KeyArgs{
			DeletionWindowInDays: pulumi.Int(10),
			Description:          pulumi.String("This key is used to encrypt bucket objects"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucket(ctx, "mybucket", &s3.BucketArgs{
			ServerSideEncryptionConfiguration: &s3.BucketServerSideEncryptionConfigurationArgs{
				Rule: &s3.BucketServerSideEncryptionConfigurationRuleArgs{
					ApplyServerSideEncryptionByDefault: &s3.BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs{
						KmsMasterKeyId: mykey.Arn,
						SseAlgorithm:   pulumi.String("aws:kms"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

