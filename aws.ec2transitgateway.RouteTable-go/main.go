package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.NewRouteTable(ctx, "example", &ec2transitgateway.RouteTableArgs{
			TransitGatewayId: pulumi.Any(aws_ec2_transit_gateway.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

