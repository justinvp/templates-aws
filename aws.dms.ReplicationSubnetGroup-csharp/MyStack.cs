using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Create a new replication subnet group
        var test = new Aws.Dms.ReplicationSubnetGroup("test", new Aws.Dms.ReplicationSubnetGroupArgs
        {
            ReplicationSubnetGroupDescription = "Test replication subnet group",
            ReplicationSubnetGroupId = "test-dms-replication-subnet-group-tf",
            SubnetIds = 
            {
                "subnet-12345678",
            },
            Tags = 
            {
                { "Name", "test" },
            },
        });
    }

}

