using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var by_version_stage = Output.Create(Aws.SecretsManager.GetSecretVersion.InvokeAsync(new Aws.SecretsManager.GetSecretVersionArgs
        {
            SecretId = data.Aws_secretsmanager_secret.Example.Id,
            VersionStage = "example",
        }));
    }

}

