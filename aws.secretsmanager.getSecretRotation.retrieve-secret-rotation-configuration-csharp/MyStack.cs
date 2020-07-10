using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.SecretsManager.GetSecretRotation.InvokeAsync(new Aws.SecretsManager.GetSecretRotationArgs
        {
            SecretId = data.Aws_secretsmanager_secret.Example.Id,
        }));
    }

}

