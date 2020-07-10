package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/servicediscovery"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicediscovery.NewHttpNamespace(ctx, "example", &servicediscovery.HttpNamespaceArgs{
			Description: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

