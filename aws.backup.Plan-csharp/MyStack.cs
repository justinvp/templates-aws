using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Backup.Plan("example", new Aws.Backup.PlanArgs
        {
            Rules = 
            {
                new Aws.Backup.Inputs.PlanRuleArgs
                {
                    RuleName = "tf_example_backup_rule",
                    Schedule = "cron(0 12 * * ? *)",
                    TargetVaultName = aws_backup_vault.Test.Name,
                },
            },
        });
    }

}

