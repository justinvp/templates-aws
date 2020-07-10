using System.Collections.Generic;
using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.DocDB.Cluster("default", new Aws.DocDB.ClusterArgs
        {
            AvailabilityZones = 
            {
                "us-west-2a",
                "us-west-2b",
                "us-west-2c",
            },
            ClusterIdentifier = "docdb-cluster-demo",
            MasterPassword = "barbut8chars",
            MasterUsername = "foo",
        });
        var clusterInstances = new List<Aws.DocDB.ClusterInstance>();
        for (var rangeIndex = 0; rangeIndex < 2; rangeIndex++)
        {
            var range = new { Value = rangeIndex };
            clusterInstances.Add(new Aws.DocDB.ClusterInstance($"clusterInstances-{range.Value}", new Aws.DocDB.ClusterInstanceArgs
            {
                ClusterIdentifier = @default.Id,
                Identifier = $"docdb-cluster-demo-{range.Value}",
                InstanceClass = "db.r5.large",
            }));
        }
    }

}

