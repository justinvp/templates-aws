package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testBucket, err := s3.NewBucket(ctx, "testBucket", nil)
		if err != nil {
			return err
		}
		inventory, err := s3.NewBucket(ctx, "inventory", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewInventory(ctx, "testInventory", &s3.InventoryArgs{
			Bucket: testBucket.ID(),
			Destination: &s3.InventoryDestinationArgs{
				Bucket: &s3.InventoryDestinationBucketArgs{
					BucketArn: inventory.Arn,
					Format:    pulumi.String("ORC"),
				},
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

