package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := organizations.NewOrganization(ctx, "org", &organizations.OrganizationArgs{
			AwsServiceAccessPrincipals: pulumi.StringArray{
				pulumi.String("cloudtrail.amazonaws.com"),
				pulumi.String("config.amazonaws.com"),
			},
			FeatureSet: pulumi.String("ALL"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

