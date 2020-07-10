using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ram.ResourceAssociation("example", new Aws.Ram.ResourceAssociationArgs
        {
            ResourceArn = aws_subnet.Example.Arn,
            ResourceShareArn = aws_ram_resource_share.Example.Arn,
        });
    }

}

