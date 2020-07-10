package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.NewTransitGateway(ctx, "example", &ec2transitgateway.TransitGatewayArgs{
			Description: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

