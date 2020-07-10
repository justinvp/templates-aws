using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Rds.Cluster("example", new Aws.Rds.ClusterArgs
        {
            ClusterIdentifier = "example",
            DbSubnetGroupName = aws_db_subnet_group.Example.Name,
            EngineMode = "multimaster",
            MasterPassword = "barbarbarbar",
            MasterUsername = "foo",
            SkipFinalSnapshot = true,
        });
    }

}

