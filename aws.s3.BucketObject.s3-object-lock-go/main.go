package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		examplebucket, err := s3.NewBucket(ctx, "examplebucket", &s3.BucketArgs{
			Acl: pulumi.String("private"),
			ObjectLockConfiguration: &s3.BucketObjectLockConfigurationArgs{
				ObjectLockEnabled: pulumi.String("Enabled"),
			},
			Versioning: &s3.BucketVersioningArgs{
				Enabled: pulumi.Bool(true),
			},
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "examplebucketObject", &s3.BucketObjectArgs{
			Bucket:                    examplebucket.ID(),
			ForceDestroy:              pulumi.Bool(true),
			Key:                       pulumi.String("someobject"),
			ObjectLockLegalHoldStatus: pulumi.String("ON"),
			ObjectLockMode:            pulumi.String("GOVERNANCE"),
			ObjectLockRetainUntilDate: pulumi.String("2021-12-31T23:59:60Z"),
			Source:                    pulumi.NewFileAsset("important.txt"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

