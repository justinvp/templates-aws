using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Sns.Topic("example", new Aws.Sns.TopicArgs
        {
        });
        var bar = new Aws.AutoScaling.Group("bar", new Aws.AutoScaling.GroupArgs
        {
        });
        var foo = new Aws.AutoScaling.Group("foo", new Aws.AutoScaling.GroupArgs
        {
        });
        var exampleNotifications = new Aws.AutoScaling.Notification("exampleNotifications", new Aws.AutoScaling.NotificationArgs
        {
            GroupNames = 
            {
                bar.Name,
                foo.Name,
            },
            Notifications = 
            {
                "autoscaling:EC2_INSTANCE_LAUNCH",
                "autoscaling:EC2_INSTANCE_TERMINATE",
                "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
                "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
            },
            TopicArn = example.Arn,
        });
    }

}

