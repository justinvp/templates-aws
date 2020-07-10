using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Dax.SubnetGroup("example", new Aws.Dax.SubnetGroupArgs
        {
            SubnetIds = 
            {
                aws_subnet.Example1.Id,
                aws_subnet.Example2.Id,
            },
        });
    }

}

