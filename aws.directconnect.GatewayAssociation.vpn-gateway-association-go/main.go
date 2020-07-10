package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directconnect"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
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
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.255.255.0/28"),
		})
		if err != nil {
			return err
		}
		exampleVpnGateway, err := ec2.NewVpnGateway(ctx, "exampleVpnGateway", &ec2.VpnGatewayArgs{
			VpcId: exampleVpc.ID(),
		})
		if err != nil {
			return err
		}
		_, err = directconnect.NewGatewayAssociation(ctx, "exampleGatewayAssociation", &directconnect.GatewayAssociationArgs{
			AssociatedGatewayId: exampleVpnGateway.ID(),
			DxGatewayId:         exampleGateway.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

