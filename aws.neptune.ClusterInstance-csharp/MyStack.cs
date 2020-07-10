using System.Collections.Generic;
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
        var example = new List<Aws.Neptune.ClusterInstance>();
        for (var rangeIndex = 0; rangeIndex < 2; rangeIndex++)
        {
            var range = new { Value = rangeIndex };
            example.Add(new Aws.Neptune.ClusterInstance($"example-{range.Value}", new Aws.Neptune.ClusterInstanceArgs
            {
                ApplyImmediately = true,
                ClusterIdentifier = @default.Id,
                Engine = "neptune",
                InstanceClass = "db.r4.large",
            }));
        }
    }

}

