package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := organizations.NewAccount(ctx, "account", &organizations.AccountArgs{
			Email: pulumi.String("john@doe.org"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

