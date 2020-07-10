using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var production = new Aws.Ssm.MaintenanceWindow("production", new Aws.Ssm.MaintenanceWindowArgs
        {
            Cutoff = 1,
            Duration = 3,
            Schedule = "cron(0 16 ? * TUE *)",
        });
    }

}

