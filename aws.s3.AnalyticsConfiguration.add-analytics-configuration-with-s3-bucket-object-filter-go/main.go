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
		_, err = s3.NewAnalyticsConfiguration(ctx, "example_filtered", &s3.AnalyticsConfigurationArgs{
			Bucket: example.Bucket,
			Filter: &s3.AnalyticsConfigurationFilterArgs{
				Prefix: pulumi.String("documents/"),
				Tags: pulumi.StringMap{
					"priority": pulumi.String("high"),
					"class":    pulumi.String("blue"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

