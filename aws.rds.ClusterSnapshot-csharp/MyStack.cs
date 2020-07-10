using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Rds.ClusterSnapshot("example", new Aws.Rds.ClusterSnapshotArgs
        {
            DbClusterIdentifier = aws_rds_cluster.Example.Id,
            DbClusterSnapshotIdentifier = "resourcetestsnapshot1234",
        });
    }

}

