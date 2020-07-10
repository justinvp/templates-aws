package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directconnect"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := directconnect.NewHostedPrivateVirtualInterface(ctx, "foo", &directconnect.HostedPrivateVirtualInterfaceArgs{
			AddressFamily: pulumi.String("ipv4"),
			BgpAsn:        pulumi.Int(65352),
			ConnectionId:  pulumi.String("dxcon-zzzzzzzz"),
			Vlan:          pulumi.Int(4094),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

