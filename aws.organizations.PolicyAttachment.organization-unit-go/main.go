package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := organizations.NewPolicyAttachment(ctx, "unit", &organizations.PolicyAttachmentArgs{
			PolicyId: pulumi.Any(aws_organizations_policy.Example.Id),
			TargetId: pulumi.Any(aws_organizations_organizational_unit.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

