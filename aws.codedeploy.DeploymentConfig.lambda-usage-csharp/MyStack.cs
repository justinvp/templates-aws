using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var fooDeploymentConfig = new Aws.CodeDeploy.DeploymentConfig("fooDeploymentConfig", new Aws.CodeDeploy.DeploymentConfigArgs
        {
            ComputePlatform = "Lambda",
            DeploymentConfigName = "test-deployment-config",
            TrafficRoutingConfig = new Aws.CodeDeploy.Inputs.DeploymentConfigTrafficRoutingConfigArgs
            {
                TimeBasedLinear = new Aws.CodeDeploy.Inputs.DeploymentConfigTrafficRoutingConfigTimeBasedLinearArgs
                {
                    Interval = 10,
                    Percentage = 10,
                },
                Type = "TimeBasedLinear",
            },
        });
        var fooDeploymentGroup = new Aws.CodeDeploy.DeploymentGroup("fooDeploymentGroup", new Aws.CodeDeploy.DeploymentGroupArgs
        {
            AlarmConfiguration = new Aws.CodeDeploy.Inputs.DeploymentGroupAlarmConfigurationArgs
            {
                Alarms = 
                {
                    "my-alarm-name",
                },
                Enabled = true,
            },
            AppName = aws_codedeploy_app.Foo_app.Name,
            AutoRollbackConfiguration = new Aws.CodeDeploy.Inputs.DeploymentGroupAutoRollbackConfigurationArgs
            {
                Enabled = true,
                Events = 
                {
                    "DEPLOYMENT_STOP_ON_ALARM",
                },
            },
            DeploymentConfigName = fooDeploymentConfig.Id,
            DeploymentGroupName = "bar",
            ServiceRoleArn = aws_iam_role.Foo_role.Arn,
        });
    }

}

