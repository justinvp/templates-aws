using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var clusterName = Output.Create(Aws.Rds.GetCluster.InvokeAsync(new Aws.Rds.GetClusterArgs
        {
            ClusterIdentifier = "clusterName",
        }));
    }

}

