package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/efs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		foo, err := ec2.NewVpc(ctx, "foo", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		alphaSubnet, err := ec2.NewSubnet(ctx, "alphaSubnet", &ec2.SubnetArgs{
			AvailabilityZone: pulumi.String("us-west-2a"),
			CidrBlock:        pulumi.String("10.0.1.0/24"),
			VpcId:            foo.ID(),
		})
		if err != nil {
			return err
		}
		_, err = efs.NewMountTarget(ctx, "alphaMountTarget", &efs.MountTargetArgs{
			FileSystemId: pulumi.Any(aws_efs_file_system.Foo.Id),
			SubnetId:     alphaSubnet.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

