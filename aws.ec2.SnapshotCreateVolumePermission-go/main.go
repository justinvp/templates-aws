package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ebs"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := ebs.NewVolume(ctx, "example", &ebs.VolumeArgs{
			AvailabilityZone: pulumi.String("us-west-2a"),
			Size:             pulumi.Int(40),
		})
		if err != nil {
			return err
		}
		exampleSnapshot, err := ebs.NewSnapshot(ctx, "exampleSnapshot", &ebs.SnapshotArgs{
			VolumeId: example.ID(),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewSnapshotCreateVolumePermission(ctx, "examplePerm", &ec2.SnapshotCreateVolumePermissionArgs{
			AccountId:  pulumi.String("12345678"),
			SnapshotId: exampleSnapshot.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

