package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/storagegateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagegateway.NewGateway(ctx, "example", &storagegateway.GatewayArgs{
			GatewayIpAddress: pulumi.String("1.2.3.4"),
			GatewayName:      pulumi.String("example"),
			GatewayTimezone:  pulumi.String("GMT"),
			GatewayType:      pulumi.String("STORED"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

