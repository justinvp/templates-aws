using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.S3.AccountPublicAccessBlock("example", new Aws.S3.AccountPublicAccessBlockArgs
        {
            BlockPublicAcls = true,
            BlockPublicPolicy = true,
        });
    }

}

