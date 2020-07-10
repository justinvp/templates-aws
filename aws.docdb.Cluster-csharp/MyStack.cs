using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var docdb = new Aws.DocDB.Cluster("docdb", new Aws.DocDB.ClusterArgs
        {
            BackupRetentionPeriod = 5,
            ClusterIdentifier = "my-docdb-cluster",
            Engine = "docdb",
            MasterPassword = "mustbeeightchars",
            MasterUsername = "foo",
            PreferredBackupWindow = "07:00-09:00",
            SkipFinalSnapshot = true,
        });
    }

}

