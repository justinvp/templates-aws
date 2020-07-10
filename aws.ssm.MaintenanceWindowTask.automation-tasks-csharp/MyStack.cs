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
            TaskArn = "AWS-RestartEC2Instance",
            TaskInvocationParameters = new Aws.Ssm.Inputs.MaintenanceWindowTaskTaskInvocationParametersArgs
            {
                AutomationParameters = new Aws.Ssm.Inputs.MaintenanceWindowTaskTaskInvocationParametersAutomationParametersArgs
                {
                    DocumentVersion = "$LATEST",
                    Parameter = 
                    {
                        
                        {
                            { "name", "InstanceId" },
                            { "values", 
                            {
                                aws_instance.Example.Id,
                            } },
                        },
                    },
                },
            },
            TaskType = "AUTOMATION",
            WindowId = aws_ssm_maintenance_window.Example.Id,
        });
    }

}

