using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Backup.GetPlan.InvokeAsync(new Aws.Backup.GetPlanArgs
        {
            PlanId = "tf_example_backup_plan_id",
        }));
    }

}

