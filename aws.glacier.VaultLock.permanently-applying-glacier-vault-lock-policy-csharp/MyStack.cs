using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Glacier.VaultLock("example", new Aws.Glacier.VaultLockArgs
        {
            CompleteLock = true,
            Policy = data.Aws_iam_policy_document.Example.Json,
            VaultName = aws_glacier_vault.Example.Name,
        });
    }

}

