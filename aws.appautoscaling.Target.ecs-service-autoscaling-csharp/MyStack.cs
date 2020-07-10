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
            ResourceId = $"service/{aws_ecs_cluster.Example.Name}/{aws_ecs_service.Example.Name}",
            ScalableDimension = "ecs:service:DesiredCount",
            ServiceNamespace = "ecs",
        });
    }

}

