using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ssm.MaintenanceWindowTask("example", new Aws.Ssm.MaintenanceWindowTaskArgs
        {
            MaxConcurrency = "2",
            MaxErrors = "1",
            Priority = 1,
            ServiceRoleArn = aws_iam_role.Example.Arn,
            Targets = 
            {
                new Aws.Ssm.Inputs.MaintenanceWindowTaskTargetArgs
                {
                    Key = "InstanceIds",
                    Values = 
                    {
                        aws_instance.Example.Id,
                    },
                },
            },
            TaskArn = "AWS-RunShellScript",
            TaskInvocationParameters = new Aws.Ssm.Inputs.MaintenanceWindowTaskTaskInvocationParametersArgs
            {
                RunCommandParameters = new Aws.Ssm.Inputs.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersArgs
                {
                    NotificationConfig = new Aws.Ssm.Inputs.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigArgs
                    {
                        NotificationArn = aws_sns_topic.Example.Arn,
                        NotificationEvents = 
                        {
                            "All",
                        },
                        NotificationType = "Command",
                    },
                    OutputS3Bucket = aws_s3_bucket.Example.Bucket,
                    OutputS3KeyPrefix = "output",
                    Parameter = 
                    {
                        
                        {
                            { "name", "commands" },
                            { "values", 
                            {
                                "date",
                            } },
                        },
                    },
                    ServiceRoleArn = aws_iam_role.Example.Arn,
                    TimeoutSeconds = 600,
                },
            },
            TaskType = "RUN_COMMAND",
            WindowId = aws_ssm_maintenance_window.Example.Id,
        });
    }

}

