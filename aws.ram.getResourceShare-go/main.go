package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ram"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ram.LookupResourceShare(ctx, &ram.LookupResourceShareArgs{
			Name:          "example",
			ResourceOwner: "SELF",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

