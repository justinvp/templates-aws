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
		_, err = s3.NewBucketMetric(ctx, "example_filtered", &s3.BucketMetricArgs{
			Bucket: example.Bucket,
			Filter: &s3.BucketMetricFilterArgs{
				Prefix: pulumi.String("documents/"),
				Tags: pulumi.StringMap{
					"class":    pulumi.String("blue"),
					"priority": pulumi.String("high"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

