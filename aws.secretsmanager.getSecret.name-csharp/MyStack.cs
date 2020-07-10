using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var by_name = Output.Create(Aws.SecretsManager.GetSecret.InvokeAsync(new Aws.SecretsManager.GetSecretArgs
        {
            Name = "example",
        }));
    }

}

