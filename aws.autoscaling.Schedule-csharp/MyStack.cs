using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var foobarGroup = new Aws.AutoScaling.Group("foobarGroup", new Aws.AutoScaling.GroupArgs
        {
            AvailabilityZones = 
            {
                "us-west-2a",
            },
            ForceDelete = true,
            HealthCheckGracePeriod = 300,
            HealthCheckType = "ELB",
            MaxSize = 1,
            MinSize = 1,
            TerminationPolicies = 
            {
                "OldestInstance",
            },
        });
        var foobarSchedule = new Aws.AutoScaling.Schedule("foobarSchedule", new Aws.AutoScaling.ScheduleArgs
        {
            AutoscalingGroupName = foobarGroup.Name,
            DesiredCapacity = 0,
            EndTime = "2016-12-12T06:00:00Z",
            MaxSize = 1,
            MinSize = 0,
            ScheduledActionName = "foobar",
            StartTime = "2016-12-11T18:00:00Z",
        });
    }

}

