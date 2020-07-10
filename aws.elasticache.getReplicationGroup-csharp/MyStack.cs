using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bar = Output.Create(Aws.ElastiCache.GetReplicationGroup.InvokeAsync(new Aws.ElastiCache.GetReplicationGroupArgs
        {
            ReplicationGroupId = "example",
        }));
    }

}

