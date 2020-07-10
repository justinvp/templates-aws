package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewVpnGatewayRoutePropagation(ctx, "example", &ec2.VpnGatewayRoutePropagationArgs{
			RouteTableId: pulumi.Any(aws_route_table.Example.Id),
			VpnGatewayId: pulumi.Any(aws_vpn_gateway.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

