package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.LookupUser(ctx, &iam.LookupUserArgs{
			UserName: "an_example_user_name",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

