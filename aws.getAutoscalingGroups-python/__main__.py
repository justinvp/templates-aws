import pulumi
import pulumi_aws as aws

groups = aws.get_autoscaling_groups(filters=[
    {
        "name": "key",
        "values": ["Team"],
    },
    {
        "name": "value",
        "values": ["Pets"],
    },
])
slack_notifications = aws.autoscaling.Notification("slackNotifications",
    group_names=groups.names,
    notifications=[
        "autoscaling:EC2_INSTANCE_LAUNCH",
        "autoscaling:EC2_INSTANCE_TERMINATE",
        "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
        "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
    ],
    topic_arn="TOPIC ARN")

