using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Backup.GetVault.InvokeAsync(new Aws.Backup.GetVaultArgs
        {
            Name = "example_backup_vault",
        }));
    }

}

