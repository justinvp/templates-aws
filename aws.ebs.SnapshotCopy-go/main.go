package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ebs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := ebs.NewVolume(ctx, "example", &ebs.VolumeArgs{
			AvailabilityZone: pulumi.String("us-west-2a"),
			Size:             pulumi.Int(40),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("HelloWorld"),
			},
		})
		if err != nil {
			return err
		}
		exampleSnapshot, err := ebs.NewSnapshot(ctx, "exampleSnapshot", &ebs.SnapshotArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("HelloWorld_snap"),
			},
			VolumeId: example.ID(),
		})
		if err != nil {
			return err
		}
		_, err = ebs.NewSnapshotCopy(ctx, "exampleCopy", &ebs.SnapshotCopyArgs{
			SourceRegion:     pulumi.String("us-west-2"),
			SourceSnapshotId: exampleSnapshot.ID(),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("HelloWorld_copy_snap"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

