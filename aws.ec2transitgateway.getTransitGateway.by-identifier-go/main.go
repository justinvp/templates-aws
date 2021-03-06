package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "tgw-12345678"
		_, err := ec2transitgateway.LookupTransitGateway(ctx, &ec2transitgateway.LookupTransitGatewayArgs{
			Id: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

