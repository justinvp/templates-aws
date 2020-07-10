using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var by_arn = Output.Create(Aws.SecretsManager.GetSecret.InvokeAsync(new Aws.SecretsManager.GetSecretArgs
        {
            Arn = "arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456",
        }));
    }

}

