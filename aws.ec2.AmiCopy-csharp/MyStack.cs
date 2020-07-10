using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ec2.AmiCopy("example", new Aws.Ec2.AmiCopyArgs
        {
            Description = "A copy of ami-xxxxxxxx",
            SourceAmiId = "ami-xxxxxxxx",
            SourceAmiRegion = "us-west-1",
            Tags = 
            {
                { "Name", "HelloWorld" },
            },
        });
    }

}

