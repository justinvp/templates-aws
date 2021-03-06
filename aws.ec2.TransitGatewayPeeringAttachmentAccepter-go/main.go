package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewTransitGatewayPeeringAttachmentAccepter(ctx, "example", &ec2.TransitGatewayPeeringAttachmentAccepterArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example cross-account attachment"),
			},
			TransitGatewayAttachmentId: pulumi.Any(aws_ec2_transit_gateway_peering_attachment.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

