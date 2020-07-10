using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var defaultCluster = new Aws.RedShift.Cluster("defaultCluster", new Aws.RedShift.ClusterArgs
        {
            ClusterIdentifier = "tf-redshift-cluster",
            ClusterType = "single-node",
            DatabaseName = "mydb",
            MasterPassword = "Mustbe8characters",
            MasterUsername = "foo",
            NodeType = "dc1.large",
        });
        var defaultSnapshotSchedule = new Aws.RedShift.SnapshotSchedule("defaultSnapshotSchedule", new Aws.RedShift.SnapshotScheduleArgs
        {
            Definitions = 
            {
                "rate(12 hours)",
            },
            Identifier = "tf-redshift-snapshot-schedule",
        });
        var defaultSnapshotScheduleAssociation = new Aws.RedShift.SnapshotScheduleAssociation("defaultSnapshotScheduleAssociation", new Aws.RedShift.SnapshotScheduleAssociationArgs
        {
            ClusterIdentifier = defaultCluster.Id,
            ScheduleIdentifier = defaultSnapshotSchedule.Id,
        });
    }

}

