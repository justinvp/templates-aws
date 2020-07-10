using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.DocDB.SubnetGroup("default", new Aws.DocDB.SubnetGroupArgs
        {
            SubnetIds = 
            {
                aws_subnet.Frontend.Id,
                aws_subnet.Backend.Id,
            },
            Tags = 
            {
                { "Name", "My docdb subnet group" },
            },
        });
    }

}

