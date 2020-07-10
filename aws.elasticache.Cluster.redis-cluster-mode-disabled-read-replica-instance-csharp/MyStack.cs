using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var replica = new Aws.ElastiCache.Cluster("replica", new Aws.ElastiCache.ClusterArgs
        {
            ReplicationGroupId = aws_elasticache_replication_group.Example.Id,
        });
    }

}

