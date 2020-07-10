package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directconnect"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := directconnect.LookupGateway(ctx, &directconnect.LookupGatewayArgs{
			Name: "example",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

