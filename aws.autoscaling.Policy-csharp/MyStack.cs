using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bar = new Aws.AutoScaling.Group("bar", new Aws.AutoScaling.GroupArgs
        {
            AvailabilityZones = 
            {
                "us-east-1a",
            },
            ForceDelete = true,
            HealthCheckGracePeriod = 300,
            HealthCheckType = "ELB",
            LaunchConfiguration = aws_launch_configuration.Foo.Name,
            MaxSize = 5,
            MinSize = 2,
        });
        var bat = new Aws.AutoScaling.Policy("bat", new Aws.AutoScaling.PolicyArgs
        {
            AdjustmentType = "ChangeInCapacity",
            AutoscalingGroupName = bar.Name,
            Cooldown = 300,
            ScalingAdjustment = 4,
        });
    }

}

