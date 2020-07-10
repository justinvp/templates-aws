using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ebs.EncryptionByDefault("example", new Aws.Ebs.EncryptionByDefaultArgs
        {
            Enabled = true,
        });
    }

}

