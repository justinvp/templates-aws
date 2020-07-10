package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/guardduty"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		master, err := guardduty.NewDetector(ctx, "master", nil)
		if err != nil {
			return err
		}
		memberDetector, err := guardduty.NewDetector(ctx, "memberDetector", nil, pulumi.Provider("aws.dev"))
		if err != nil {
			return err
		}
		_, err = guardduty.NewMember(ctx, "dev", &guardduty.MemberArgs{
			AccountId:  memberDetector.AccountId,
			DetectorId: master.ID(),
			Email:      pulumi.String("required@example.com"),
			Invite:     pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = guardduty.NewInviteAccepter(ctx, "memberInviteAccepter", &guardduty.InviteAccepterArgs{
			DetectorId:      memberDetector.ID(),
			MasterAccountId: master.AccountId,
		}, pulumi.Provider("aws.dev"), pulumi.DependsOn([]pulumi.Resource{
			"aws_guardduty_member.dev",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

