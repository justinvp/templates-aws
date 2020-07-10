package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewEventPermission(ctx, "devAccountAccess", &cloudwatch.EventPermissionArgs{
			Principal:   pulumi.String("123456789012"),
			StatementId: pulumi.String("DevAccountAccess"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

