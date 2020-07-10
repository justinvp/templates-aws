package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.LookupPolicy(ctx, &iam.LookupPolicyArgs{
			Arn: "arn:aws:iam::123456789012:policy/UsersManageOwnCredentials",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

