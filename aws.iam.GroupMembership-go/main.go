package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		group, err := iam.NewGroup(ctx, "group", nil)
		if err != nil {
			return err
		}
		userOne, err := iam.NewUser(ctx, "userOne", nil)
		if err != nil {
			return err
		}
		userTwo, err := iam.NewUser(ctx, "userTwo", nil)
		if err != nil {
			return err
		}
		_, err = iam.NewGroupMembership(ctx, "team", &iam.GroupMembershipArgs{
			Group: group.Name,
			Users: pulumi.StringArray{
				userOne.Name,
				userTwo.Name,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

