package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := s3.NewBucket(ctx, "example", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketMetric(ctx, "example_entire_bucket", &s3.BucketMetricArgs{
			Bucket: example.Bucket,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

