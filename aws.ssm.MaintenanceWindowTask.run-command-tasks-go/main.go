package main

import (
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
			TaskArn: pulumi.String("AWS-RunShellScript"),
			TaskInvocationParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersArgs{
				RunCommandParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersArgs{
					NotificationConfig: &ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigArgs{
						NotificationArn: pulumi.Any(aws_sns_topic.Example.Arn),
						NotificationEvents: pulumi.StringArray{
							pulumi.String("All"),
						},
						NotificationType: pulumi.String("Command"),
					},
					OutputS3Bucket:    pulumi.Any(aws_s3_bucket.Example.Bucket),
					OutputS3KeyPrefix: pulumi.String("output"),
					Parameter: pulumi.MapArray{
						pulumi.Map{
							"name": pulumi.String("commands"),
							"values": pulumi.StringArray{
								pulumi.String("date"),
							},
						},
					},
					ServiceRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
					TimeoutSeconds: pulumi.Int(600),
				},
			},
			TaskType: pulumi.String("RUN_COMMAND"),
			WindowId: pulumi.Any(aws_ssm_maintenance_window.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

