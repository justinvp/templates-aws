package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewOptionGroup(ctx, "example", &rds.OptionGroupArgs{
			EngineName:         pulumi.String("sqlserver-ee"),
			MajorEngineVersion: pulumi.String("11.00"),
			Options: rds.OptionGroupOptionArray{
				&rds.OptionGroupOptionArgs{
					OptionName: pulumi.String("Timezone"),
					OptionSettings: rds.OptionGroupOptionOptionSettingArray{
						&rds.OptionGroupOptionOptionSettingArgs{
							Name:  pulumi.String("TIME_ZONE"),
							Value: pulumi.String("UTC"),
						},
					},
				},
				&rds.OptionGroupOptionArgs{
					OptionName: pulumi.String("SQLSERVER_BACKUP_RESTORE"),
					OptionSettings: rds.OptionGroupOptionOptionSettingArray{
						&rds.OptionGroupOptionOptionSettingArgs{
							Name:  pulumi.String("IAM_ROLE_ARN"),
							Value: pulumi.Any(aws_iam_role.Example.Arn),
						},
					},
				},
				&rds.OptionGroupOptionArgs{
					OptionName: pulumi.String("TDE"),
				},
			},
			OptionGroupDescription: pulumi.String("Option Group"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

