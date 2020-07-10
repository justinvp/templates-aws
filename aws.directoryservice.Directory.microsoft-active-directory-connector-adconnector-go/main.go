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
		bar, err := ec2.NewSubnet(ctx, "bar", &ec2.SubnetArgs{
			AvailabilityZone: pulumi.String("us-west-2b"),
			CidrBlock:        pulumi.String("10.0.2.0/24"),
			VpcId:            main.ID(),
		})
		if err != nil {
			return err
		}
		_, err = directoryservice.NewDirectory(ctx, "connector", &directoryservice.DirectoryArgs{
			ConnectSettings: &directoryservice.DirectoryConnectSettingsArgs{
				CustomerDnsIps: pulumi.StringArray{
					pulumi.String("A.B.C.D"),
				},
				CustomerUsername: pulumi.String("Admin"),
				SubnetIds: pulumi.StringArray{
					foo.ID(),
					bar.ID(),
				},
				VpcId: main.ID(),
			},
			Password: pulumi.String("SuperSecretPassw0rd"),
			Size:     pulumi.String("Small"),
			Type:     pulumi.String("ADConnector"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

