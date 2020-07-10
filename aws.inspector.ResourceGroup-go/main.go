package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/inspector"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := inspector.NewResourceGroup(ctx, "example", &inspector.ResourceGroupArgs{
			Tags: pulumi.StringMap{
				"Env":  pulumi.String("bar"),
				"Name": pulumi.String("foo"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

