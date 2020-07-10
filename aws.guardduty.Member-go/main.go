package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/guardduty"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		master, err := guardduty.NewDetector(ctx, "master", &guardduty.DetectorArgs{
			Enable: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		memberDetector, err := guardduty.NewDetector(ctx, "memberDetector", &guardduty.DetectorArgs{
			Enable: pulumi.Bool(true),
		}, pulumi.Provider("aws.dev"))
		if err != nil {
			return err
		}
		_, err = guardduty.NewMember(ctx, "memberMember", &guardduty.MemberArgs{
			AccountId:         memberDetector.AccountId,
			DetectorId:        master.ID(),
			Email:             pulumi.String("required@example.com"),
			Invite:            pulumi.Bool(true),
			InvitationMessage: pulumi.String("please accept guardduty invitation"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

