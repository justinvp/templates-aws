using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.RedShift.Cluster("default", new Aws.RedShift.ClusterArgs
        {
            ClusterIdentifier = "tf-redshift-cluster",
            ClusterType = "single-node",
            DatabaseName = "mydb",
            MasterPassword = "Mustbe8characters",
            MasterUsername = "foo",
            NodeType = "dc1.large",
        });
    }

}

