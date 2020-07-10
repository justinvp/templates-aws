package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.NewAccountAlias(ctx, "alias", &iam.AccountAliasArgs{
			AccountAlias: pulumi.String("my-account-alias"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

