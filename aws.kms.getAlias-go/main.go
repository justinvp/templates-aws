package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kms.LookupAlias(ctx, &kms.LookupAliasArgs{
			Name: "alias/aws/s3",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

