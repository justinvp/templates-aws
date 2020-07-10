using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var cluster = Output.Create(Aws.CloudHsmV2.GetCluster.InvokeAsync(new Aws.CloudHsmV2.GetClusterArgs
        {
            ClusterId = @var.Cloudhsm_cluster_id,
        }));
        var cloudhsmV2Hsm = new Aws.CloudHsmV2.Hsm("cloudhsmV2Hsm", new Aws.CloudHsmV2.HsmArgs
        {
            ClusterId = cluster.Apply(cluster => cluster.ClusterId),
            SubnetId = cluster.Apply(cluster => cluster.SubnetIds[0]),
        });
    }

}

