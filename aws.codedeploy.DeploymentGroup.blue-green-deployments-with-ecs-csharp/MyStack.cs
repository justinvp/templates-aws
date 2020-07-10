using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleApplication = new Aws.CodeDeploy.Application("exampleApplication", new Aws.CodeDeploy.ApplicationArgs
        {
            ComputePlatform = "ECS",
        });
        var exampleDeploymentGroup = new Aws.CodeDeploy.DeploymentGroup("exampleDeploymentGroup", new Aws.CodeDeploy.DeploymentGroupArgs
        {
            AppName = exampleApplication.Name,
            AutoRollbackConfiguration = new Aws.CodeDeploy.Inputs.DeploymentGroupAutoRollbackConfigurationArgs
            {
                Enabled = true,
                Events = 
                {
                    "DEPLOYMENT_FAILURE",
                },
            },
            BlueGreenDeploymentConfig = new Aws.CodeDeploy.Inputs.DeploymentGroupBlueGreenDeploymentConfigArgs
            {
                DeploymentReadyOption = new Aws.CodeDeploy.Inputs.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs
                {
                    ActionOnTimeout = "CONTINUE_DEPLOYMENT",
                },
                TerminateBlueInstancesOnDeploymentSuccess = new Aws.CodeDeploy.Inputs.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs
                {
                    Action = "TERMINATE",
                    TerminationWaitTimeInMinutes = 5,
                },
            },
            DeploymentConfigName = "CodeDeployDefault.ECSAllAtOnce",
            DeploymentGroupName = "example",
            DeploymentStyle = new Aws.CodeDeploy.Inputs.DeploymentGroupDeploymentStyleArgs
            {
                DeploymentOption = "WITH_TRAFFIC_CONTROL",
                DeploymentType = "BLUE_GREEN",
            },
            EcsService = new Aws.CodeDeploy.Inputs.DeploymentGroupEcsServiceArgs
            {
                ClusterName = aws_ecs_cluster.Example.Name,
                ServiceName = aws_ecs_service.Example.Name,
            },
            LoadBalancerInfo = new Aws.CodeDeploy.Inputs.DeploymentGroupLoadBalancerInfoArgs
            {
                TargetGroupPairInfo = new Aws.CodeDeploy.Inputs.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoArgs
                {
                    ProdTrafficRoute = new Aws.CodeDeploy.Inputs.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteArgs
                    {
                        ListenerArns = 
                        {
                            aws_lb_listener.Example.Arn,
                        },
                    },
                    TargetGroup = 
                    {
                        
                        {
                            { "name", aws_lb_target_group.Blue.Name },
                        },
                        
                        {
                            { "name", aws_lb_target_group.Green.Name },
                        },
                    },
                },
            },
            ServiceRoleArn = aws_iam_role.Example.Arn,
        });
    }

}

