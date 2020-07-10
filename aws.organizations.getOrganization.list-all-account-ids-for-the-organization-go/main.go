package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := organizations.LookupOrganization(ctx, nil, nil)
		if err != nil {
			return err
		}
		var splat0 []string
		for _, val0 := range example.Accounts {
			splat0 = append(splat0, val0.Id)
		}
		ctx.Export("accountIds", splat0)
		return nil
	})
}

