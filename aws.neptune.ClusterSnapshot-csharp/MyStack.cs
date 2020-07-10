using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Neptune.ClusterSnapshot("example", new Aws.Neptune.ClusterSnapshotArgs
        {
            DbClusterIdentifier = aws_neptune_cluster.Example.Id,
            DbClusterSnapshotIdentifier = "resourcetestsnapshot1234",
        });
    }

}

