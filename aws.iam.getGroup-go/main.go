package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.LookupGroup(ctx, &iam.LookupGroupArgs{
			GroupName: "an_example_group_name",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

