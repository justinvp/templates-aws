using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var cluster = Output.Create(Aws.CloudHsmV2.GetCluster.InvokeAsync(new Aws.CloudHsmV2.GetClusterArgs
        {
            ClusterId = "cluster-testclusterid",
        }));
    }

}

