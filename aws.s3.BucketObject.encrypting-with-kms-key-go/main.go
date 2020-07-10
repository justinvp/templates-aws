package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		examplekms, err := kms.NewKey(ctx, "examplekms", &kms.KeyArgs{
			DeletionWindowInDays: pulumi.Int(7),
			Description:          pulumi.String("KMS key 1"),
		})
		if err != nil {
			return err
		}
		examplebucket, err := s3.NewBucket(ctx, "examplebucket", &s3.BucketArgs{
			Acl: pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "examplebucketObject", &s3.BucketObjectArgs{
			Bucket:   examplebucket.ID(),
			Key:      pulumi.String("someobject"),
			KmsKeyId: examplekms.Arn,
			Source:   pulumi.NewFileAsset("index.html"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

