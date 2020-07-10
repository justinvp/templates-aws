package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ram"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ram.NewResourceShare(ctx, "example", &ram.ResourceShareArgs{
			AllowExternalPrincipals: pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Production"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

