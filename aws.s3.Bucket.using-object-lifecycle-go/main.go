package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			Acl: pulumi.String("private"),
			LifecycleRules: s3.BucketLifecycleRuleArray{
				&s3.BucketLifecycleRuleArgs{
					Enabled: pulumi.Bool(true),
					Expiration: &s3.BucketLifecycleRuleExpirationArgs{
						Days: pulumi.Int(90),
					},
					Id:     pulumi.String("log"),
					Prefix: pulumi.String("log/"),
					Tags: pulumi.StringMap{
						"autoclean": pulumi.String("true"),
						"rule":      pulumi.String("log"),
					},
					Transition: pulumi.MapArray{
						pulumi.Map{
							"days":         pulumi.Float64(30),
							"storageClass": pulumi.String("STANDARD_IA"),
						},
						pulumi.Map{
							"days":         pulumi.Float64(60),
							"storageClass": pulumi.String("GLACIER"),
						},
					},
				},
				&s3.BucketLifecycleRuleArgs{
					Enabled: pulumi.Bool(true),
					Expiration: &s3.BucketLifecycleRuleExpirationArgs{
						Date: pulumi.String("2016-01-12"),
					},
					Id:     pulumi.String("tmp"),
					Prefix: pulumi.String("tmp/"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucket(ctx, "versioningBucket", &s3.BucketArgs{
			Acl: pulumi.String("private"),
			LifecycleRules: s3.BucketLifecycleRuleArray{
				&s3.BucketLifecycleRuleArgs{
					Enabled: pulumi.Bool(true),
					NoncurrentVersionExpiration: &s3.BucketLifecycleRuleNoncurrentVersionExpirationArgs{
						Days: pulumi.Int(90),
					},
					NoncurrentVersionTransition: pulumi.MapArray{
						pulumi.Map{
							"days":         pulumi.Float64(30),
							"storageClass": pulumi.String("STANDARD_IA"),
						},
						pulumi.Map{
							"days":         pulumi.Float64(60),
							"storageClass": pulumi.String("GLACIER"),
						},
					},
					Prefix: pulumi.String("config/"),
				},
			},
			Versioning: &s3.BucketVersioningArgs{
				Enabled: pulumi.Bool(true),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

