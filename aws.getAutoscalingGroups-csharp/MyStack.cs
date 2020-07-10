using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var groups = Output.Create(Aws.GetAutoscalingGroups.InvokeAsync(new Aws.GetAutoscalingGroupsArgs
        {
            Filters = 
            {
                new Aws.Inputs.GetAutoscalingGroupsFilterArgs
                {
                    Name = "key",
                    Values = 
                    {
                        "Team",
                    },
                },
                new Aws.Inputs.GetAutoscalingGroupsFilterArgs
                {
                    Name = "value",
                    Values = 
                    {
                        "Pets",
                    },
                },
            },
        }));
        var slackNotifications = new Aws.AutoScaling.Notification("slackNotifications", new Aws.AutoScaling.NotificationArgs
        {
            GroupNames = groups.Apply(groups => groups.Names),
            Notifications = 
            {
                "autoscaling:EC2_INSTANCE_LAUNCH",
                "autoscaling:EC2_INSTANCE_TERMINATE",
                "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
                "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
            },
            TopicArn = "TOPIC ARN",
        });
    }

}

