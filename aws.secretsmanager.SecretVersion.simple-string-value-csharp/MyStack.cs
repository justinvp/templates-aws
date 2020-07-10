using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.SecretsManager.SecretVersion("example", new Aws.SecretsManager.SecretVersionArgs
        {
            SecretId = aws_secretsmanager_secret.Example.Id,
            SecretString = "example-string-to-protect",
        });
    }

}

