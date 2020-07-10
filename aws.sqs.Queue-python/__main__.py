import pulumi
import json
import pulumi_aws as aws

queue = aws.sqs.Queue("queue",
    delay_seconds=90,
    max_message_size=2048,
    message_retention_seconds=86400,
    receive_wait_time_seconds=10,
    redrive_policy=json.dumps({
        "deadLetterTargetArn": aws_sqs_queue["queue_deadletter"]["arn"],
        "maxReceiveCount": 4,
    }),
    tags={
        "Environment": "production",
    })

