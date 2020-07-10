using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.SecretsManager.Secret("example", new Aws.SecretsManager.SecretArgs
        {
        });
    }

}

