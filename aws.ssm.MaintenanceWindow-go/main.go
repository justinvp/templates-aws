package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewMaintenanceWindow(ctx, "production", &ssm.MaintenanceWindowArgs{
			Cutoff:   pulumi.Int(1),
			Duration: pulumi.Int(3),
			Schedule: pulumi.String("cron(0 16 ? * TUE *)"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

