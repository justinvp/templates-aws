using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Glue.Connection("example", new Aws.Glue.ConnectionArgs
        {
            ConnectionProperties = 
            {
                { "JDBC_CONNECTION_URL", $"jdbc:mysql://{aws_rds_cluster.Example.Endpoint}/exampledatabase" },
                { "PASSWORD", "examplepassword" },
                { "USERNAME", "exampleusername" },
            },
            PhysicalConnectionRequirements = new Aws.Glue.Inputs.ConnectionPhysicalConnectionRequirementsArgs
            {
                AvailabilityZone = aws_subnet.Example.Availability_zone,
                SecurityGroupIdList = 
                {
                    aws_security_group.Example.Id,
                },
                SubnetId = aws_subnet.Example.Id,
            },
        });
    }

}

