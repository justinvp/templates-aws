package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directoryservice"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		foo, err := ec2.NewSubnet(ctx, "foo", &ec2.SubnetArgs{
			AvailabilityZone: pulumi.String("us-west-2a"),
			CidrBlock:        pulumi.String("10.0.1.0/24"),
			VpcId:            main.ID(),
		})
		if err != nil {
			return err
		}
		barSubnet, err := ec2.NewSubnet(ctx, "barSubnet", &ec2.SubnetArgs{
			AvailabilityZone: pulumi.String("us-west-2b"),
			CidrBlock:        pulumi.String("10.0.2.0/24"),
			VpcId:            main.ID(),
		})
		if err != nil {
			return err
		}
		_, err = directoryservice.NewDirectory(ctx, "barDirectory", &directoryservice.DirectoryArgs{
			Edition:  pulumi.String("Standard"),
			Password: pulumi.String("SuperSecretPassw0rd"),
			Tags: pulumi.StringMap{
				"Project": pulumi.String("foo"),
			},
			Type: pulumi.String("MicrosoftAD"),
			VpcSettings: &directoryservice.DirectoryVpcSettingsArgs{
				SubnetIds: pulumi.StringArray{
					foo.ID(),
					barSubnet.ID(),
				},
				VpcId: main.ID(),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

