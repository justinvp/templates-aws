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
                { "JDBC_CONNECTION_URL", "jdbc:mysql://example.com/exampledatabase" },
                { "PASSWORD", "examplepassword" },
                { "USERNAME", "exampleusername" },
            },
        });
    }

}

