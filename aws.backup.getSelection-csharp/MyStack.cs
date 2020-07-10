using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Backup.GetSelection.InvokeAsync(new Aws.Backup.GetSelectionArgs
        {
            PlanId = data.Aws_backup_plan.Example.Id,
            SelectionId = "selection-id-example",
        }));
    }

}

