package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ram"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ram.NewPrincipalAssociation(ctx, "example", &ram.PrincipalAssociationArgs{
			Principal:        pulumi.Any(aws_organizations_organization.Example.Arn),
			ResourceShareArn: pulumi.Any(aws_ram_resource_share.Example.Arn),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

