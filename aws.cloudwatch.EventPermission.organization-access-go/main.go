package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewEventPermission(ctx, "organizationAccess", &cloudwatch.EventPermissionArgs{
			Condition: &cloudwatch.EventPermissionConditionArgs{
				Key:   pulumi.String("aws:PrincipalOrgID"),
				Type:  pulumi.String("StringEquals"),
				Value: pulumi.Any(aws_organizations_organization.Example.Id),
			},
			Principal:   pulumi.String("*"),
			StatementId: pulumi.String("OrganizationAccess"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

