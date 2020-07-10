package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "available"
		available, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{
			State: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewSubnet(ctx, "primary", &ec2.SubnetArgs{
			AvailabilityZone: pulumi.String(available.Names[0]),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewSubnet(ctx, "secondary", &ec2.SubnetArgs{
			AvailabilityZone: pulumi.String(available.Names[1]),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

