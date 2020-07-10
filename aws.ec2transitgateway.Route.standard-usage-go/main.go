package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.NewRoute(ctx, "example", &ec2transitgateway.RouteArgs{
			DestinationCidrBlock:       pulumi.String("0.0.0.0/0"),
			TransitGatewayAttachmentId: pulumi.Any(aws_ec2_transit_gateway_vpc_attachment.Example.Id),
			TransitGatewayRouteTableId: pulumi.Any(aws_ec2_transit_gateway.Example.Association_default_route_table_id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

