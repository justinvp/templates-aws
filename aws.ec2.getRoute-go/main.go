package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := subnetId
		_, err := ec2.LookupRouteTable(ctx, &ec2.LookupRouteTableArgs{
			SubnetId: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		opt1 := "10.0.1.0/24"
		route, err := ec2.LookupRoute(ctx, &ec2.LookupRouteArgs{
			DestinationCidrBlock: &opt1,
			RouteTableId:         aws_route_table.Selected.Id,
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.LookupNetworkInterface(ctx, &ec2.LookupNetworkInterfaceArgs{
			NetworkInterfaceId: route.NetworkInterfaceId,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

