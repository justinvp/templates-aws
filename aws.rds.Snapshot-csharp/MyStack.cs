using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bar = new Aws.Rds.Instance("bar", new Aws.Rds.InstanceArgs
        {
            AllocatedStorage = 10,
            BackupRetentionPeriod = 0,
            Engine = "MySQL",
            EngineVersion = "5.6.21",
            InstanceClass = "db.t2.micro",
            MaintenanceWindow = "Fri:09:00-Fri:09:30",
            Name = "baz",
            ParameterGroupName = "default.mysql5.6",
            Password = "barbarbarbar",
            Username = "foo",
        });
        var test = new Aws.Rds.Snapshot("test", new Aws.Rds.SnapshotArgs
        {
            DbInstanceIdentifier = bar.Id,
            DbSnapshotIdentifier = "testsnapshot1234",
        });
    }

}

