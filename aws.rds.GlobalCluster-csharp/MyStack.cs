using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var primary = new Aws.Provider("primary", new Aws.ProviderArgs
        {
            Region = "us-east-2",
        });
        var secondary = new Aws.Provider("secondary", new Aws.ProviderArgs
        {
            Region = "us-west-2",
        });
        var example = new Aws.Rds.GlobalCluster("example", new Aws.Rds.GlobalClusterArgs
        {
            GlobalClusterIdentifier = "example",
        }, new CustomResourceOptions
        {
            Provider = "aws.primary",
        });
        var primaryCluster = new Aws.Rds.Cluster("primaryCluster", new Aws.Rds.ClusterArgs
        {
            EngineMode = "global",
            GlobalClusterIdentifier = example.Id,
        }, new CustomResourceOptions
        {
            Provider = "aws.primary",
        });
        var primaryClusterInstance = new Aws.Rds.ClusterInstance("primaryClusterInstance", new Aws.Rds.ClusterInstanceArgs
        {
            ClusterIdentifier = primaryCluster.Id,
        }, new CustomResourceOptions
        {
            Provider = "aws.primary",
        });
        var secondaryCluster = new Aws.Rds.Cluster("secondaryCluster", new Aws.Rds.ClusterArgs
        {
            EngineMode = "global",
            GlobalClusterIdentifier = example.Id,
        }, new CustomResourceOptions
        {
            Provider = "aws.secondary",
            DependsOn = 
            {
                "aws_rds_cluster_instance.primary",
            },
        });
        var secondaryClusterInstance = new Aws.Rds.ClusterInstance("secondaryClusterInstance", new Aws.Rds.ClusterInstanceArgs
        {
            ClusterIdentifier = secondaryCluster.Id,
        }, new CustomResourceOptions
        {
            Provider = "aws.secondary",
        });
    }

}

