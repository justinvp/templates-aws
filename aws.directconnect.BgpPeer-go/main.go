package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directconnect"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := directconnect.NewBgpPeer(ctx, "peer", &directconnect.BgpPeerArgs{
			AddressFamily:      pulumi.String("ipv6"),
			BgpAsn:             pulumi.Int(65351),
			VirtualInterfaceId: pulumi.Any(aws_dx_private_virtual_interface.Foo.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

