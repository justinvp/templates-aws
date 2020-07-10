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
            TaskArn = aws_sfn_activity.Example.Id,
            TaskInvocationParameters = new Aws.Ssm.Inputs.MaintenanceWindowTaskTaskInvocationParametersArgs
            {
                StepFunctionsParameters = new Aws.Ssm.Inputs.MaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersArgs
                {
                    Input = "{\"key1\":\"value1\"}",
                    Name = "example",
                },
            },
            TaskType = "STEP_FUNCTIONS",
            WindowId = aws_ssm_maintenance_window.Example.Id,
        });
    }

}

