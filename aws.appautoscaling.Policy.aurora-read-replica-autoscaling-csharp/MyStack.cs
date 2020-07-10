using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var replicasTarget = new Aws.AppAutoScaling.Target("replicasTarget", new Aws.AppAutoScaling.TargetArgs
        {
            MaxCapacity = 15,
            MinCapacity = 1,
            ResourceId = $"cluster:{aws_rds_cluster.Example.Id}",
            ScalableDimension = "rds:cluster:ReadReplicaCount",
            ServiceNamespace = "rds",
        });
        var replicasPolicy = new Aws.AppAutoScaling.Policy("replicasPolicy", new Aws.AppAutoScaling.PolicyArgs
        {
            PolicyType = "TargetTrackingScaling",
            ResourceId = replicasTarget.ResourceId,
            ScalableDimension = replicasTarget.ScalableDimension,
            ServiceNamespace = replicasTarget.ServiceNamespace,
            TargetTrackingScalingPolicyConfiguration = new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationArgs
            {
                PredefinedMetricSpecification = new Aws.AppAutoScaling.Inputs.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs
                {
                    PredefinedMetricType = "RDSReaderAverageCPUUtilization",
                },
                ScaleInCooldown = 300,
                ScaleOutCooldown = 300,
                TargetValue = 75,
            },
        });
    }

}

