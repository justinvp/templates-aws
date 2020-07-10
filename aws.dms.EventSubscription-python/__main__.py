import pulumi
import pulumi_aws as aws

example = aws.dms.EventSubscription("example",
    enabled=True,
    event_categories=[
        "creation",
        "failure",
    ],
    sns_topic_arn=aws_sns_topic["example"]["arn"],
    source_ids=[aws_dms_replication_task["example"]["replication_task_id"]],
    source_type="replication-task",
    tags={
        "Name": "example",
    })

