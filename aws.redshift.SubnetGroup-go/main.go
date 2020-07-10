package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooVpc, err := ec2.NewVpc(ctx, "fooVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.1.0.0/16"),
		})
		if err != nil {
			return err
		}
		fooSubnet, err := ec2.NewSubnet(ctx, "fooSubnet", &ec2.SubnetArgs{
			AvailabilityZone: pulumi.String("us-west-2a"),
			CidrBlock:        pulumi.String("10.1.1.0/24"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-dbsubnet-test-1"),
			},
			VpcId: fooVpc.ID(),
		})
		if err != nil {
			return err
		}
		bar, err := ec2.NewSubnet(ctx, "bar", &ec2.SubnetArgs{
			AvailabilityZone: pulumi.String("us-west-2b"),
			CidrBlock:        pulumi.String("10.1.2.0/24"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-dbsubnet-test-2"),
			},
			VpcId: fooVpc.ID(),
		})
		if err != nil {
			return err
		}
		_, err = redshift.NewSubnetGroup(ctx, "fooSubnetGroup", &redshift.SubnetGroupArgs{
			SubnetIds: pulumi.StringArray{
				fooSubnet.ID(),
				bar.ID(),
			},
			Tags: pulumi.StringMap{
				"environment": pulumi.String("Production"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

