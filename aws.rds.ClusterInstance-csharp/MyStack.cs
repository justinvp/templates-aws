using System.Collections.Generic;
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
            ClusterIdentifier = "aurora-cluster-demo",
            DatabaseName = "mydb",
            MasterPassword = "barbut8chars",
            MasterUsername = "foo",
        });
        var clusterInstances = new List<Aws.Rds.ClusterInstance>();
        for (var rangeIndex = 0; rangeIndex < 2; rangeIndex++)
        {
            var range = new { Value = rangeIndex };
            clusterInstances.Add(new Aws.Rds.ClusterInstance($"clusterInstances-{range.Value}", new Aws.Rds.ClusterInstanceArgs
            {
                ClusterIdentifier = @default.Id,
                Identifier = $"aurora-cluster-demo-{range.Value}",
                InstanceClass = "db.r4.large",
            }));
        }
    }

}

