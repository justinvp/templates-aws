package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directconnect"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
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
		exampleTransitGateway, err := ec2transitgateway.NewTransitGateway(ctx, "exampleTransitGateway", nil)
		if err != nil {
			return err
		}
		_, err = directconnect.NewGatewayAssociation(ctx, "exampleGatewayAssociation", &directconnect.GatewayAssociationArgs{
			AllowedPrefixes: pulumi.StringArray{
				pulumi.String("10.255.255.0/30"),
				pulumi.String("10.255.255.8/30"),
			},
			AssociatedGatewayId: exampleTransitGateway.ID(),
			DxGatewayId:         exampleGateway.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

