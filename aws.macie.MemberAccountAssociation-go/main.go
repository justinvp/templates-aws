package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/macie"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := macie.NewMemberAccountAssociation(ctx, "example", &macie.MemberAccountAssociationArgs{
			MemberAccountId: pulumi.String("123456789012"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

