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
            Engine = "aurora-mysql",
            EngineVersion = "5.7.mysql_aurora.2.03.2",
            MasterPassword = "bar",
            MasterUsername = "foo",
            PreferredBackupWindow = "07:00-09:00",
        });
    }

}

