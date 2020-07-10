using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.ElastiCache.Cluster("example", new Aws.ElastiCache.ClusterArgs
        {
            Engine = "memcached",
            NodeType = "cache.m4.large",
            NumCacheNodes = 2,
            ParameterGroupName = "default.memcached1.4",
            Port = 11211,
        });
    }

}

