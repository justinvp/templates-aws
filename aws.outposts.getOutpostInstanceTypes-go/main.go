package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/outposts"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := outposts.GetOutpostInstanceTypes(ctx, &outposts.GetOutpostInstanceTypesArgs{
			Arn: data.Aws_outposts_outpost.Example.Arn,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

