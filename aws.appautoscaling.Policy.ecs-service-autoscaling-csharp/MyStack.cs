using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ecsTarget = new Aws.AppAutoScaling.Target("ecsTarget", new Aws.AppAutoScaling.TargetArgs
        {
            MaxCapacity = 4,
            MinCapacity = 1,
            ResourceId = "service/clusterName/serviceName",
            ScalableDimension = "ecs:service:DesiredCount",
            ServiceNamespace = "ecs",
        });
        var ecsPolicy = new Aws.AppAutoScaling.Policy("ecsPolicy", new Aws.AppAutoScaling.PolicyArgs
        {
            PolicyType = "StepScaling",
            ResourceId = ecsTarget.ResourceId,
            ScalableDimension = ecsTarget.ScalableDimension,
            ServiceNamespace = ecsTarget.ServiceNamespace,
            StepScalingPolicyConfiguration = new Aws.AppAutoScaling.Inputs.PolicyStepScalingPolicyConfigurationArgs
            {
                AdjustmentType = "ChangeInCapacity",
                Cooldown = 60,
                MetricAggregationType = "Maximum",
                StepAdjustment = 
                {
                    
                    {
                        { "metricIntervalUpperBound", 0 },
                        { "scalingAdjustment", -1 },
                    },
                },
            },
        });
    }

}

