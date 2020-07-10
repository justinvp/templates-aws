using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Backup.Vault("example", new Aws.Backup.VaultArgs
        {
            KmsKeyArn = aws_kms_key.Example.Arn,
        });
    }

}

