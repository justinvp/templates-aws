package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ram"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleResourceShare, err := ram.NewResourceShare(ctx, "exampleResourceShare", &ram.ResourceShareArgs{
			AllowExternalPrincipals: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = ram.NewPrincipalAssociation(ctx, "examplePrincipalAssociation", &ram.PrincipalAssociationArgs{
			Principal:        pulumi.String("111111111111"),
			ResourceShareArn: exampleResourceShare.Arn,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

