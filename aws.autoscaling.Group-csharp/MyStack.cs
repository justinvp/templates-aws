using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var test = new Aws.Ec2.PlacementGroup("test", new Aws.Ec2.PlacementGroupArgs
        {
            Strategy = "cluster",
        });
        var bar = new Aws.AutoScaling.Group("bar", new Aws.AutoScaling.GroupArgs
        {
            DesiredCapacity = 4,
            ForceDelete = true,
            HealthCheckGracePeriod = 300,
            HealthCheckType = "ELB",
            InitialLifecycleHooks = 
            {
                new Aws.AutoScaling.Inputs.GroupInitialLifecycleHookArgs
                {
                    DefaultResult = "CONTINUE",
                    HeartbeatTimeout = 2000,
                    LifecycleTransition = "autoscaling:EC2_INSTANCE_LAUNCHING",
                    Name = "foobar",
                    NotificationMetadata = @"{
  ""foo"": ""bar""
}

",
                    NotificationTargetArn = "arn:aws:sqs:us-east-1:444455556666:queue1*",
                    RoleArn = "arn:aws:iam::123456789012:role/S3Access",
                },
            },
            LaunchConfiguration = aws_launch_configuration.Foobar.Name,
            MaxSize = 5,
            MinSize = 2,
            PlacementGroup = test.Id,
            Tags = 
            {
                new Aws.AutoScaling.Inputs.GroupTagArgs
                {
                    Key = "foo",
                    PropagateAtLaunch = true,
                    Value = "bar",
                },
                new Aws.AutoScaling.Inputs.GroupTagArgs
                {
                    Key = "lorem",
                    PropagateAtLaunch = false,
                    Value = "ipsum",
                },
            },
            VpcZoneIdentifiers = 
            {
                aws_subnet.Example1.Id,
                aws_subnet.Example2.Id,
            },
        });
    }

}

