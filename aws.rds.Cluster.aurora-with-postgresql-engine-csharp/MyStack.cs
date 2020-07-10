using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var postgresql = new Aws.Rds.Cluster("postgresql", new Aws.Rds.ClusterArgs
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
            Engine = "aurora-postgresql",
            MasterPassword = "bar",
            MasterUsername = "foo",
            PreferredBackupWindow = "07:00-09:00",
        });
    }

}

