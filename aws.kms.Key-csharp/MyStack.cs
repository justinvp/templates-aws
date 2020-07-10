using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var key = new Aws.Kms.Key("key", new Aws.Kms.KeyArgs
        {
            DeletionWindowInDays = 10,
            Description = "KMS key 1",
        });
    }

}

