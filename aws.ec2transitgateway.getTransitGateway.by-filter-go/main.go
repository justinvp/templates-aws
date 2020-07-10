package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.LookupTransitGateway(ctx, &ec2transitgateway.LookupTransitGatewayArgs{
			Filters: []ec2transitgateway.GetTransitGatewayFilter{
				ec2transitgateway.GetTransitGatewayFilter{
					Name: "options.amazon-side-asn",
					Values: []string{
						"64512",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

