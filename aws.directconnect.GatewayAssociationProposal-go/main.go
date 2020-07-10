package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directconnect"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := directconnect.NewGatewayAssociationProposal(ctx, "example", &directconnect.GatewayAssociationProposalArgs{
			AssociatedGatewayId:     pulumi.Any(aws_vpn_gateway.Example.Id),
			DxGatewayId:             pulumi.Any(aws_dx_gateway.Example.Id),
			DxGatewayOwnerAccountId: pulumi.Any(aws_dx_gateway.Example.Owner_account_id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

