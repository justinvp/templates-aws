package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := iam.LookupAccountAlias(ctx, nil, nil)
		if err != nil {
			return err
		}
		ctx.Export("accountId", current.AccountAlias)
		return nil
	})
}

