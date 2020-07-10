package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		org, err := organizations.LookupOrganization(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = organizations.GetOrganizationalUnits(ctx, &organizations.GetOrganizationalUnitsArgs{
			ParentId: org.Roots[0].Id,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

