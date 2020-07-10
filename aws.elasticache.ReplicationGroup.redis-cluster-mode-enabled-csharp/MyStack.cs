using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var baz = new Aws.ElastiCache.ReplicationGroup("baz", new Aws.ElastiCache.ReplicationGroupArgs
        {
            AutomaticFailoverEnabled = true,
            ClusterMode = new Aws.ElastiCache.Inputs.ReplicationGroupClusterModeArgs
            {
                NumNodeGroups = 2,
                ReplicasPerNodeGroup = 1,
            },
            NodeType = "cache.t2.small",
            ParameterGroupName = "default.redis3.2.cluster.on",
            Port = 6379,
            ReplicationGroupDescription = "test description",
        });
    }

}

