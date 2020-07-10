using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var key = new Aws.Kms.Key("key", new Aws.Kms.KeyArgs
        {
        });
        var @alias = new Aws.Kms.Alias("alias", new Aws.Kms.AliasArgs
        {
            TargetKeyId = key.KeyId,
        });
    }

}

