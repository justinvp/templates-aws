using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleVault = new Aws.Glacier.Vault("exampleVault", new Aws.Glacier.VaultArgs
        {
        });
        var examplePolicyDocument = exampleVault.Arn.Apply(arn => Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
        {
            Statements = 
            {
                new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
                {
                    Actions = 
                    {
                        "glacier:DeleteArchive",
                    },
                    Condition = 
                    {
                        
                        {
                            { "test", "NumericLessThanEquals" },
                            { "values", 
                            {
                                "365",
                            } },
                            { "variable", "glacier:ArchiveAgeinDays" },
                        },
                    },
                    Effect = "Deny",
                    Resources = 
                    {
                        arn,
                    },
                },
            },
        }));
        var exampleVaultLock = new Aws.Glacier.VaultLock("exampleVaultLock", new Aws.Glacier.VaultLockArgs
        {
            CompleteLock = false,
            Policy = examplePolicyDocument.Apply(examplePolicyDocument => examplePolicyDocument.Json),
            VaultName = exampleVault.Name,
        });
    }

}

