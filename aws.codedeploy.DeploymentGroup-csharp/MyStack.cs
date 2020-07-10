using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleRole = new Aws.Iam.Role("exampleRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = @"{
  ""Version"": ""2012-10-17"",
  ""Statement"": [
    {
      ""Sid"": """",
      ""Effect"": ""Allow"",
      ""Principal"": {
        ""Service"": ""codedeploy.amazonaws.com""
      },
      ""Action"": ""sts:AssumeRole""
    }
  ]
}

",
        });
        var aWSCodeDeployRole = new Aws.Iam.RolePolicyAttachment("aWSCodeDeployRole", new Aws.Iam.RolePolicyAttachmentArgs
        {
            PolicyArn = "arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole",
            Role = exampleRole.Name,
        });
        var exampleApplication = new Aws.CodeDeploy.Application("exampleApplication", new Aws.CodeDeploy.ApplicationArgs
        {
        });
        var exampleTopic = new Aws.Sns.Topic("exampleTopic", new Aws.Sns.TopicArgs
        {
        });
        var exampleDeploymentGroup = new Aws.CodeDeploy.DeploymentGroup("exampleDeploymentGroup", new Aws.CodeDeploy.DeploymentGroupArgs
        {
            AlarmConfiguration = new Aws.CodeDeploy.Inputs.DeploymentGroupAlarmConfigurationArgs
            {
                Alarms = 
                {
                    "my-alarm-name",
                },
                Enabled = true,
            },
            AppName = exampleApplication.Name,
            AutoRollbackConfiguration = new Aws.CodeDeploy.Inputs.DeploymentGroupAutoRollbackConfigurationArgs
            {
                Enabled = true,
                Events = 
                {
                    "DEPLOYMENT_FAILURE",
                },
            },
            DeploymentGroupName = "example-group",
            Ec2TagSets = 
            {
                new Aws.CodeDeploy.Inputs.DeploymentGroupEc2TagSetArgs
                {
                    Ec2TagFilter = 
                    {
                        
                        {
                            { "key", "filterkey1" },
                            { "type", "KEY_AND_VALUE" },
                            { "value", "filtervalue" },
                        },
                        
                        {
                            { "key", "filterkey2" },
                            { "type", "KEY_AND_VALUE" },
                            { "value", "filtervalue" },
                        },
                    },
                },
            },
            ServiceRoleArn = exampleRole.Arn,
            TriggerConfigurations = 
            {
                new Aws.CodeDeploy.Inputs.DeploymentGroupTriggerConfigurationArgs
                {
                    TriggerEvents = 
                    {
                        "DeploymentFailure",
                    },
                    TriggerName = "example-trigger",
                    TriggerTargetArn = exampleTopic.Arn,
                },
            },
        });
    }

}

