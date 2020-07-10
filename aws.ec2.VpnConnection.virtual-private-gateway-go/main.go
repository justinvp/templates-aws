package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		vpc, err := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		vpnGateway, err := ec2.NewVpnGateway(ctx, "vpnGateway", &ec2.VpnGatewayArgs{
			VpcId: vpc.ID(),
		})
		if err != nil {
			return err
		}
		customerGateway, err := ec2.NewCustomerGateway(ctx, "customerGateway", &ec2.CustomerGatewayArgs{
			BgpAsn:    pulumi.Int(65000),
			IpAddress: pulumi.String("172.0.0.1"),
			Type:      pulumi.String("ipsec.1"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpnConnection(ctx, "main", &ec2.VpnConnectionArgs{
			CustomerGatewayId: customerGateway.ID(),
			StaticRoutesOnly:  pulumi.Bool(true),
			Type:              pulumi.String("ipsec.1"),
			VpnGatewayId:      vpnGateway.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

