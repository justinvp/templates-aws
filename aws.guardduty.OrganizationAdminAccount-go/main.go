package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/guardduty"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleOrganization, err := organizations.NewOrganization(ctx, "exampleOrganization", &organizations.OrganizationArgs{
			AwsServiceAccessPrincipals: pulumi.StringArray{
				pulumi.String("guardduty.amazonaws.com"),
			},
			FeatureSet: pulumi.String("ALL"),
		})
		if err != nil {
			return err
		}
		_, err = guardduty.NewDetector(ctx, "exampleDetector", nil)
		if err != nil {
			return err
		}
		_, err = guardduty.NewOrganizationAdminAccount(ctx, "exampleOrganizationAdminAccount", &guardduty.OrganizationAdminAccountArgs{
			AdminAccountId: pulumi.String("123456789012"),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleOrganization,
		}))
		if err != nil {
			return err
		}
		return nil
	})
}

