using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var replicas = new Aws.AppAutoScaling.Target("replicas", new Aws.AppAutoScaling.TargetArgs
        {
            MaxCapacity = 15,
            MinCapacity = 1,
            ResourceId = $"cluster:{aws_rds_cluster.Example.Id}",
            ScalableDimension = "rds:cluster:ReadReplicaCount",
            ServiceNamespace = "rds",
        });
    }

}

