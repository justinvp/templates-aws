package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/backup"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := backup.LookupSelection(ctx, &backup.LookupSelectionArgs{
			PlanId:      data.Aws_backup_plan.Example.Id,
			SelectionId: "selection-id-example",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

