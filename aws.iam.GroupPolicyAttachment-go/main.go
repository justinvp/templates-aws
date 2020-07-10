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
		policy, err := iam.NewPolicy(ctx, "policy", &iam.PolicyArgs{
			Description: pulumi.String("A test policy"),
			Policy:      pulumi.String(""),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewGroupPolicyAttachment(ctx, "test_attach", &iam.GroupPolicyAttachmentArgs{
			Group:     group.Name,
			PolicyArn: policy.Arn,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

