package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		test, err := s3.NewBucket(ctx, "test", nil)
		if err != nil {
			return err
		}
		inventory, err := s3.NewBucket(ctx, "inventory", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewInventory(ctx, "test_prefix", &s3.InventoryArgs{
			Bucket: test.ID(),
			Destination: &s3.InventoryDestinationArgs{
				Bucket: &s3.InventoryDestinationBucketArgs{
					BucketArn: inventory.Arn,
					Format:    pulumi.String("ORC"),
					Prefix:    pulumi.String("inventory"),
				},
			},
			Filter: &s3.InventoryFilterArgs{
				Prefix: pulumi.String("documents/"),
			},
			IncludedObjectVersions: pulumi.String("All"),
			Schedule: &s3.InventoryScheduleArgs{
				Frequency: pulumi.String("Daily"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

