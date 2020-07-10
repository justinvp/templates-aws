package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleEventCategories, err := rds.GetEventCategories(ctx, nil, nil)
		if err != nil {
			return err
		}
		ctx.Export("example", exampleEventCategories.EventCategories)
		return nil
	})
}

