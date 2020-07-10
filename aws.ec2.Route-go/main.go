package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewRoute(ctx, "route", &ec2.RouteArgs{
			RouteTableId:           pulumi.String("rtb-4fbb3ac4"),
			DestinationCidrBlock:   pulumi.String("10.0.1.0/22"),
			VpcPeeringConnectionId: pulumi.String("pcx-45ff3dc1"),
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_route_table.testing",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

