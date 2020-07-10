using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var @default = new Aws.RedShift.SnapshotSchedule("default", new Aws.RedShift.SnapshotScheduleArgs
        {
            Definitions = 
            {
                "rate(12 hours)",
            },
            Identifier = "tf-redshift-snapshot-schedule",
        });
    }

}

