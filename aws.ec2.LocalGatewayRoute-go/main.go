package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewLocalGatewayRoute(ctx, "example", &ec2.LocalGatewayRouteArgs{
			DestinationCidrBlock:                pulumi.String("172.16.0.0/16"),
			LocalGatewayRouteTableId:            pulumi.Any(data.Aws_ec2_local_gateway_route_table.Example.Id),
			LocalGatewayVirtualInterfaceGroupId: pulumi.Any(data.Aws_ec2_local_gateway_virtual_interface_group.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

