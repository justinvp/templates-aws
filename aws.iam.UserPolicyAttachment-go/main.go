package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		user, err := iam.NewUser(ctx, "user", nil)
		if err != nil {
			return err
		}
		policy, err := iam.NewPolicy(ctx, "policy", &iam.PolicyArgs{
			Description: pulumi.String("A test policy"),
			Policy:      pulumi.String(""),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewUserPolicyAttachment(ctx, "test_attach", &iam.UserPolicyAttachmentArgs{
			PolicyArn: policy.Arn,
			User:      user.Name,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

