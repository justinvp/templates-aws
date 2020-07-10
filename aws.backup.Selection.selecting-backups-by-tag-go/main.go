package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/backup"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := backup.NewSelection(ctx, "example", &backup.SelectionArgs{
			IamRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			PlanId:     pulumi.Any(aws_backup_plan.Example.Id),
			SelectionTags: backup.SelectionSelectionTagArray{
				&backup.SelectionSelectionTagArgs{
					Key:   pulumi.String("foo"),
					Type:  pulumi.String("STRINGEQUALS"),
					Value: pulumi.String("bar"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

