package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		user1, err := iam.NewUser(ctx, "user1", nil)
		if err != nil {
			return err
		}
		group1, err := iam.NewGroup(ctx, "group1", nil)
		if err != nil {
			return err
		}
		group2, err := iam.NewGroup(ctx, "group2", nil)
		if err != nil {
			return err
		}
		_, err = iam.NewUserGroupMembership(ctx, "example1", &iam.UserGroupMembershipArgs{
			Groups: pulumi.StringArray{
				group1.Name,
				group2.Name,
			},
			User: user1.Name,
		})
		if err != nil {
			return err
		}
		group3, err := iam.NewGroup(ctx, "group3", nil)
		if err != nil {
			return err
		}
		_, err = iam.NewUserGroupMembership(ctx, "example2", &iam.UserGroupMembershipArgs{
			Groups: pulumi.StringArray{
				group3.Name,
			},
			User: user1.Name,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

