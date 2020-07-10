package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.NewRouteTableAssociation(ctx, "example", &ec2transitgateway.RouteTableAssociationArgs{
			TransitGatewayAttachmentId: pulumi.Any(aws_ec2_transit_gateway_vpc_attachment.Example.Id),
			TransitGatewayRouteTableId: pulumi.Any(aws_ec2_transit_gateway_route_table.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

