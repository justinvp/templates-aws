package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/backup"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := backup.NewPlan(ctx, "example", &backup.PlanArgs{
			Rules: backup.PlanRuleArray{
				&backup.PlanRuleArgs{
					RuleName:        pulumi.String("tf_example_backup_rule"),
					Schedule:        pulumi.String("cron(0 12 * * ? *)"),
					TargetVaultName: pulumi.Any(aws_backup_vault.Test.Name),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

