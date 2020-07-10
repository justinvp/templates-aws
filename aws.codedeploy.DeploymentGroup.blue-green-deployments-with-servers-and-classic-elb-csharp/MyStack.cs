using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleApplication = new Aws.CodeDeploy.Application("exampleApplication", new Aws.CodeDeploy.ApplicationArgs
        {
        });
        var exampleDeploymentGroup = new Aws.CodeDeploy.DeploymentGroup("exampleDeploymentGroup", new Aws.CodeDeploy.DeploymentGroupArgs
        {
            AppName = exampleApplication.Name,
            BlueGreenDeploymentConfig = new Aws.CodeDeploy.Inputs.DeploymentGroupBlueGreenDeploymentConfigArgs
            {
                DeploymentReadyOption = new Aws.CodeDeploy.Inputs.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs
                {
                    ActionOnTimeout = "STOP_DEPLOYMENT",
                    WaitTimeInMinutes = 60,
                },
                GreenFleetProvisioningOption = new Aws.CodeDeploy.Inputs.DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionArgs
                {
                    Action = "DISCOVER_EXISTING",
                },
                TerminateBlueInstancesOnDeploymentSuccess = new Aws.CodeDeploy.Inputs.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs
                {
                    Action = "KEEP_ALIVE",
                },
            },
            DeploymentGroupName = "example-group",
            DeploymentStyle = new Aws.CodeDeploy.Inputs.DeploymentGroupDeploymentStyleArgs
            {
                DeploymentOption = "WITH_TRAFFIC_CONTROL",
                DeploymentType = "BLUE_GREEN",
            },
            LoadBalancerInfo = new Aws.CodeDeploy.Inputs.DeploymentGroupLoadBalancerInfoArgs
            {
                ElbInfo = 
                {
                    
                    {
                        { "name", aws_elb.Example.Name },
                    },
                },
            },
            ServiceRoleArn = aws_iam_role.Example.Arn,
        });
    }

}

