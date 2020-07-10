using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Kms.ExternalKey("example", new Aws.Kms.ExternalKeyArgs
        {
            Description = "KMS EXTERNAL for AMI encryption",
        });
    }

}

