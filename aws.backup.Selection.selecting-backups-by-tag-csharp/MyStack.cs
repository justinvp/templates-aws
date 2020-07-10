using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Backup.Selection("example", new Aws.Backup.SelectionArgs
        {
            IamRoleArn = aws_iam_role.Example.Arn,
            PlanId = aws_backup_plan.Example.Id,
            SelectionTags = 
            {
                new Aws.Backup.Inputs.SelectionSelectionTagArgs
                {
                    Key = "foo",
                    Type = "STRINGEQUALS",
                    Value = "bar",
                },
            },
        });
    }

}

