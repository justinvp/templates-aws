package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "YAML"
		foo, err := ssm.LookupDocument(ctx, &ssm.LookupDocumentArgs{
			DocumentFormat: &opt0,
			Name:           "AWS-GatherSoftwareInventory",
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("content", foo.Content)
		return nil
	})
}

