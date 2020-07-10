package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewRoleAssociation(ctx, "example", &rds.RoleAssociationArgs{
			DbInstanceIdentifier: pulumi.Any(aws_db_instance.Example.Id),
			FeatureName:          pulumi.String("S3_INTEGRATION"),
			RoleArn:              pulumi.Any(aws_iam_role.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

