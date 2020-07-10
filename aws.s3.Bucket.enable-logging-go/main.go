package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		logBucket, err := s3.NewBucket(ctx, "logBucket", &s3.BucketArgs{
			Acl: pulumi.String("log-delivery-write"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			Acl: pulumi.String("private"),
			Loggings: s3.BucketLoggingArray{
				&s3.BucketLoggingArgs{
					TargetBucket: logBucket.ID(),
					TargetPrefix: pulumi.String("log/"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

