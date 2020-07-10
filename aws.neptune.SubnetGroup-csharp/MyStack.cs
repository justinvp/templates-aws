using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.Neptune.SubnetGroup("default", new Aws.Neptune.SubnetGroupArgs
        {
            SubnetIds = 
            {
                aws_subnet.Frontend.Id,
                aws_subnet.Backend.Id,
            },
            Tags = 
            {
                { "Name", "My neptune subnet group" },
            },
        });
    }

}

