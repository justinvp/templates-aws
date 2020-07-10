using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var testSnapshotCopyGrant = new Aws.RedShift.SnapshotCopyGrant("testSnapshotCopyGrant", new Aws.RedShift.SnapshotCopyGrantArgs
        {
            SnapshotCopyGrantName = "my-grant",
        });
        var testCluster = new Aws.RedShift.Cluster("testCluster", new Aws.RedShift.ClusterArgs
        {
            SnapshotCopy = new Aws.RedShift.Inputs.ClusterSnapshotCopyArgs
            {
                DestinationRegion = "us-east-2",
                GrantName = testSnapshotCopyGrant.SnapshotCopyGrantName,
            },
        });
    }

}

