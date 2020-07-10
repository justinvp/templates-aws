package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directconnect"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/providers"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := providers.Newaws(ctx, "accepter", nil)
		if err != nil {
			return err
		}
		accepterCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		creator, err := directconnect.NewHostedPrivateVirtualInterface(ctx, "creator", &directconnect.HostedPrivateVirtualInterfaceArgs{
			AddressFamily:  pulumi.String("ipv4"),
			BgpAsn:         pulumi.Int(65352),
			ConnectionId:   pulumi.String("dxcon-zzzzzzzz"),
			OwnerAccountId: pulumi.String(accepterCallerIdentity.AccountId),
			Vlan:           pulumi.Int(4094),
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_vpn_gateway.vpn_gw",
		}))
		if err != nil {
			return err
		}
		vpnGw, err := ec2.NewVpnGateway(ctx, "vpnGw", nil, pulumi.Provider("aws.accepter"))
		if err != nil {
			return err
		}
		_, err = directconnect.NewHostedPrivateVirtualInterfaceAccepter(ctx, "accepterHostedPrivateVirtualInterfaceAccepter", &directconnect.HostedPrivateVirtualInterfaceAccepterArgs{
			Tags: pulumi.StringMap{
				"Side": pulumi.String("Accepter"),
			},
			VirtualInterfaceId: creator.ID(),
			VpnGatewayId:       vpnGw.ID(),
		}, pulumi.Provider("aws.accepter"))
		if err != nil {
			return err
		}
		return nil
	})
}

