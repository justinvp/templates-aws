package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		selected, err := ec2.LookupVpnGateway(ctx, &ec2.LookupVpnGatewayArgs{
			Filters: []ec2.GetVpnGatewayFilter{
				ec2.GetVpnGatewayFilter{
					Name: "tag:Name",
					Values: []string{
						"vpn-gw",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("vpnGatewayId", selected.Id)
		return nil
	})
}

