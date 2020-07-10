package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cfg"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := organizations.NewOrganization(ctx, "exampleOrganization", &organizations.OrganizationArgs{
			AwsServiceAccessPrincipals: pulumi.StringArray{
				pulumi.String("config-multiaccountsetup.amazonaws.com"),
			},
			FeatureSet: pulumi.String("ALL"),
		})
		if err != nil {
			return err
		}
		_, err = cfg.NewOrganizationManagedRule(ctx, "exampleOrganizationManagedRule", &cfg.OrganizationManagedRuleArgs{
			RuleIdentifier: pulumi.String("IAM_PASSWORD_POLICY"),
		}, pulumi.DependsOn([]pulumi.Resource{
			"aws_organizations_organization.example",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

