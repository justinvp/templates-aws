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
        var ecsScheduledAction = new Aws.AppAutoScaling.ScheduledAction("ecsScheduledAction", new Aws.AppAutoScaling.ScheduledActionArgs
        {
            ResourceId = ecsTarget.ResourceId,
            ScalableDimension = ecsTarget.ScalableDimension,
            ScalableTargetAction = new Aws.AppAutoScaling.Inputs.ScheduledActionScalableTargetActionArgs
            {
                MaxCapacity = 10,
                MinCapacity = 1,
            },
            Schedule = "at(2006-01-02T15:04:05)",
            ServiceNamespace = ecsTarget.ServiceNamespace,
        });
    }

}

