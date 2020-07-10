package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directconnect"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleGateway, err := directconnect.NewGateway(ctx, "exampleGateway", &directconnect.GatewayArgs{
			AmazonSideAsn: pulumi.String("64512"),
		})
		if err != nil {
			return err
		}
		_, err = directconnect.NewTransitVirtualInterface(ctx, "exampleTransitVirtualInterface", &directconnect.TransitVirtualInterfaceArgs{
			AddressFamily: pulumi.String("ipv4"),
			BgpAsn:        pulumi.Int(65352),
			ConnectionId:  pulumi.Any(aws_dx_connection.Example.Id),
			DxGatewayId:   exampleGateway.ID(),
			Vlan:          pulumi.Int(4094),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

