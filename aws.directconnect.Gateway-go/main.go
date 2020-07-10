package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directconnect"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := directconnect.NewGateway(ctx, "example", &directconnect.GatewayArgs{
			AmazonSideAsn: pulumi.String("64512"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

