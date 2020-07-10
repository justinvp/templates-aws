using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.DocDB.ClusterSnapshot("example", new Aws.DocDB.ClusterSnapshotArgs
        {
            DbClusterIdentifier = aws_docdb_cluster.Example.Id,
            DbClusterSnapshotIdentifier = "resourcetestsnapshot1234",
        });
    }

}

