package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewMaintenanceWindowTask(ctx, "example", &ssm.MaintenanceWindowTaskArgs{
			MaxConcurrency: pulumi.String("2"),
			MaxErrors:      pulumi.String("1"),
			Priority:       pulumi.Int(1),
			ServiceRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Targets: ssm.MaintenanceWindowTaskTargetArray{
				&ssm.MaintenanceWindowTaskTargetArgs{
					Key: pulumi.String("InstanceIds"),
					Values: pulumi.StringArray{
						pulumi.Any(aws_instance.Example.Id),
					},
				},
			},
			TaskArn: pulumi.String("AWS-RestartEC2Instance"),
			TaskInvocationParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersArgs{
				AutomationParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersAutomationParametersArgs{
					DocumentVersion: pulumi.String(fmt.Sprintf("%v%v", "$", "LATEST")),
					Parameter: pulumi.MapArray{
						pulumi.Map{
							"name": pulumi.String("InstanceId"),
							"values": pulumi.AnyArray{
								pulumi.Any(aws_instance.Example.Id),
							},
						},
					},
				},
			},
			TaskType: pulumi.String("AUTOMATION"),
			WindowId: pulumi.Any(aws_ssm_maintenance_window.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

