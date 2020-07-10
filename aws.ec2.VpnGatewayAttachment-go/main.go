package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		network, err := ec2.NewVpc(ctx, "network", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		vpn, err := ec2.NewVpnGateway(ctx, "vpn", &ec2.VpnGatewayArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example-vpn-gateway"),
			},
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpnGatewayAttachment(ctx, "vpnAttachment", &ec2.VpnGatewayAttachmentArgs{
			VpcId:        network.ID(),
			VpnGatewayId: vpn.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

