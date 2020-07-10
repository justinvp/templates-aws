using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var fooDeploymentConfig = new Aws.CodeDeploy.DeploymentConfig("fooDeploymentConfig", new Aws.CodeDeploy.DeploymentConfigArgs
        {
            DeploymentConfigName = "test-deployment-config",
            MinimumHealthyHosts = new Aws.CodeDeploy.Inputs.DeploymentConfigMinimumHealthyHostsArgs
            {
                Type = "HOST_COUNT",
                Value = 2,
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
                    "DEPLOYMENT_FAILURE",
                },
            },
            DeploymentConfigName = fooDeploymentConfig.Id,
            DeploymentGroupName = "bar",
            Ec2TagFilters = 
            {
                new Aws.CodeDeploy.Inputs.DeploymentGroupEc2TagFilterArgs
                {
                    Key = "filterkey",
                    Type = "KEY_AND_VALUE",
                    Value = "filtervalue",
                },
            },
            ServiceRoleArn = aws_iam_role.Foo_role.Arn,
            TriggerConfigurations = 
            {
                new Aws.CodeDeploy.Inputs.DeploymentGroupTriggerConfigurationArgs
                {
                    TriggerEvents = 
                    {
                        "DeploymentFailure",
                    },
                    TriggerName = "foo-trigger",
                    TriggerTargetArn = "foo-topic-arn",
                },
            },
        });
    }

}

