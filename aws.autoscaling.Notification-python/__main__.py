import pulumi
import pulumi_aws as aws

example = aws.sns.Topic("example")
bar = aws.autoscaling.Group("bar")
foo = aws.autoscaling.Group("foo")
example_notifications = aws.autoscaling.Notification("exampleNotifications",
    group_names=[
        bar.name,
        foo.name,
    ],
    notifications=[
        "autoscaling:EC2_INSTANCE_LAUNCH",
        "autoscaling:EC2_INSTANCE_TERMINATE",
        "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
        "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
    ],
    topic_arn=example.arn)

