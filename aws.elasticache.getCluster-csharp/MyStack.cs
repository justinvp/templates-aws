using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var myCluster = Output.Create(Aws.ElastiCache.GetCluster.InvokeAsync(new Aws.ElastiCache.GetClusterArgs
        {
            ClusterId = "my-cluster-id",
        }));
    }

}

