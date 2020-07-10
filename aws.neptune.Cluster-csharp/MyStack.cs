using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.Neptune.Cluster("default", new Aws.Neptune.ClusterArgs
        {
            ApplyImmediately = true,
            BackupRetentionPeriod = 5,
            ClusterIdentifier = "neptune-cluster-demo",
            Engine = "neptune",
            IamDatabaseAuthenticationEnabled = true,
            PreferredBackupWindow = "07:00-09:00",
            SkipFinalSnapshot = true,
        });
    }

}

