using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.Rds.Cluster("default", new Aws.Rds.ClusterArgs
        {
            AvailabilityZones = 
            {
                "us-west-2a",
                "us-west-2b",
                "us-west-2c",
            },
            BackupRetentionPeriod = 5,
            ClusterIdentifier = "aurora-cluster-demo",
            DatabaseName = "mydb",
            MasterPassword = "bar",
            MasterUsername = "foo",
            PreferredBackupWindow = "07:00-09:00",
        });
        var test1 = new Aws.Rds.ClusterInstance("test1", new Aws.Rds.ClusterInstanceArgs
        {
            ApplyImmediately = true,
            ClusterIdentifier = @default.Id,
            Identifier = "test1",
            InstanceClass = "db.t2.small",
        });
        var test2 = new Aws.Rds.ClusterInstance("test2", new Aws.Rds.ClusterInstanceArgs
        {
            ApplyImmediately = true,
            ClusterIdentifier = @default.Id,
            Identifier = "test2",
            InstanceClass = "db.t2.small",
        });
        var test3 = new Aws.Rds.ClusterInstance("test3", new Aws.Rds.ClusterInstanceArgs
        {
            ApplyImmediately = true,
            ClusterIdentifier = @default.Id,
            Identifier = "test3",
            InstanceClass = "db.t2.small",
        });
        var eligible = new Aws.Rds.ClusterEndpoint("eligible", new Aws.Rds.ClusterEndpointArgs
        {
            ClusterEndpointIdentifier = "reader",
            ClusterIdentifier = @default.Id,
            CustomEndpointType = "READER",
            ExcludedMembers = 
            {
                test1.Id,
                test2.Id,
            },
        });
        var @static = new Aws.Rds.ClusterEndpoint("static", new Aws.Rds.ClusterEndpointArgs
        {
            ClusterEndpointIdentifier = "static",
            ClusterIdentifier = @default.Id,
            CustomEndpointType = "READER",
            StaticMembers = 
            {
                test1.Id,
                test3.Id,
            },
        });
    }

}

