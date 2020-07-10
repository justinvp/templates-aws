package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := subnetId
		selected, err := ec2.LookupRouteTable(ctx, &ec2.LookupRouteTableArgs{
			SubnetId: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewRoute(ctx, "route", &ec2.RouteArgs{
			DestinationCidrBlock:   pulumi.String("10.0.1.0/22"),
			RouteTableId:           pulumi.String(selected.Id),
			VpcPeeringConnectionId: pulumi.String("pcx-45ff3dc1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

