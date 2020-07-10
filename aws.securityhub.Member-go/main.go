package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		_, err = securityhub.NewMember(ctx, "exampleMember", &securityhub.MemberArgs{
			AccountId: pulumi.String("123456789012"),
			Email:     pulumi.String("example@example.com"),
			Invite:    pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_securityhub_account.example",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

