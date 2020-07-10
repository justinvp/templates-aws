package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directconnect"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := directconnect.NewPublicVirtualInterface(ctx, "foo", &directconnect.PublicVirtualInterfaceArgs{
			AddressFamily:   pulumi.String("ipv4"),
			AmazonAddress:   pulumi.String("175.45.176.2/30"),
			BgpAsn:          pulumi.Int(65352),
			ConnectionId:    pulumi.String("dxcon-zzzzzzzz"),
			CustomerAddress: pulumi.String("175.45.176.1/30"),
			RouteFilterPrefixes: pulumi.StringArray{
				pulumi.String("210.52.109.0/24"),
				pulumi.String("175.45.176.0/22"),
			},
			Vlan: pulumi.Int(4094),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

