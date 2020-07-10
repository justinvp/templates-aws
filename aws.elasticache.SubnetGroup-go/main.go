package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooVpc, err := ec2.NewVpc(ctx, "fooVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-test"),
			},
		})
		if err != nil {
			return err
		}
		fooSubnet, err := ec2.NewSubnet(ctx, "fooSubnet", &ec2.SubnetArgs{
			AvailabilityZone: pulumi.String("us-west-2a"),
			CidrBlock:        pulumi.String("10.0.0.0/24"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-test"),
			},
			VpcId: fooVpc.ID(),
		})
		if err != nil {
			return err
		}
		_, err = elasticache.NewSubnetGroup(ctx, "bar", &elasticache.SubnetGroupArgs{
			SubnetIds: pulumi.StringArray{
				fooSubnet.ID(),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

