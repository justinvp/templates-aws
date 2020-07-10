package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2clientvpn"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2clientvpn.NewNetworkAssociation(ctx, "example", &ec2clientvpn.NetworkAssociationArgs{
			ClientVpnEndpointId: pulumi.Any(aws_ec2_client_vpn_endpoint.Example.Id),
			SubnetId:            pulumi.Any(aws_subnet.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

