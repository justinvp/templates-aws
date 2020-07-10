package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.LookupRole(ctx, &iam.LookupRoleArgs{
			Name: "an_example_role_name",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

