package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directoryservice"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/workspaces"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		mainVpc, err := ec2.NewVpc(ctx, "mainVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewSubnet(ctx, "private_a", &ec2.SubnetArgs{
			AvailabilityZone: pulumi.String("us-east-1a"),
			CidrBlock:        pulumi.String("10.0.0.0/24"),
			VpcId:            mainVpc.ID(),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewSubnet(ctx, "private_b", &ec2.SubnetArgs{
			AvailabilityZone: pulumi.String("us-east-1b"),
			CidrBlock:        pulumi.String("10.0.1.0/24"),
			VpcId:            mainVpc.ID(),
		})
		if err != nil {
			return err
		}
		mainDirectory, err := directoryservice.NewDirectory(ctx, "mainDirectory", &directoryservice.DirectoryArgs{
			Password: pulumi.String("#S1ncerely"),
			Size:     pulumi.String("Small"),
			VpcSettings: &directoryservice.DirectoryVpcSettingsArgs{
				SubnetIds: pulumi.StringArray{
					private_a.ID(),
					private_b.ID(),
				},
				VpcId: mainVpc.ID(),
			},
		})
		if err != nil {
			return err
		}
		_, err = workspaces.NewDirectory(ctx, "mainWorkspaces_directoryDirectory", &workspaces.DirectoryArgs{
			DirectoryId: mainDirectory.ID(),
			SelfServicePermissions: &workspaces.DirectorySelfServicePermissionsArgs{
				IncreaseVolumeSize: pulumi.Bool(true),
				RebuildWorkspace:   pulumi.Bool(true),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

