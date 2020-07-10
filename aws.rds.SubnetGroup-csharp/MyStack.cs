using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.Rds.SubnetGroup("default", new Aws.Rds.SubnetGroupArgs
        {
            SubnetIds = 
            {
                aws_subnet.Frontend.Id,
                aws_subnet.Backend.Id,
            },
            Tags = 
            {
                { "Name", "My DB subnet group" },
            },
        });
    }

}

