package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := organizations.NewOrganizationalUnit(ctx, "example", &organizations.OrganizationalUnitArgs{
			ParentId: pulumi.Any(aws_organizations_organization.Example.Roots[0].Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

