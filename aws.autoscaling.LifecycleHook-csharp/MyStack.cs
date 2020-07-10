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
            HealthCheckType = "EC2",
            Tags = 
            {
                new Aws.AutoScaling.Inputs.GroupTagArgs
                {
                    Key = "Foo",
                    PropagateAtLaunch = true,
                    Value = "foo-bar",
                },
            },
            TerminationPolicies = 
            {
                "OldestInstance",
            },
        });
        var foobarLifecycleHook = new Aws.AutoScaling.LifecycleHook("foobarLifecycleHook", new Aws.AutoScaling.LifecycleHookArgs
        {
            AutoscalingGroupName = foobarGroup.Name,
            DefaultResult = "CONTINUE",
            HeartbeatTimeout = 2000,
            LifecycleTransition = "autoscaling:EC2_INSTANCE_LAUNCHING",
            NotificationMetadata = @"{
  ""foo"": ""bar""
}

",
            NotificationTargetArn = "arn:aws:sqs:us-east-1:444455556666:queue1*",
            RoleArn = "arn:aws:iam::123456789012:role/S3Access",
        });
    }

}

