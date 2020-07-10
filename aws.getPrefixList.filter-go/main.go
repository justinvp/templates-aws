package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.GetPrefixList(ctx, &aws.GetPrefixListArgs{
			Filters: []aws.GetPrefixListFilter{
				aws.GetPrefixListFilter{
					Name: "prefix-list-id",
					Values: []string{
						"pl-68a54001",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

