using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Msk.GetCluster.InvokeAsync(new Aws.Msk.GetClusterArgs
        {
            ClusterName = "example",
        }));
    }

}

