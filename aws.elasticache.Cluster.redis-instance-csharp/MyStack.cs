using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ElastiCache.Cluster("example", new Aws.ElastiCache.ClusterArgs
        {
            Engine = "redis",
            EngineVersion = "3.2.10",
            NodeType = "cache.m4.large",
            NumCacheNodes = 1,
            ParameterGroupName = "default.redis3.2",
            Port = 6379,
        });
    }

}

