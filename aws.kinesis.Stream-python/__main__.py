import pulumi
import pulumi_aws as aws

test_stream = aws.kinesis.Stream("testStream",
    retention_period=48,
    shard_count=1,
    shard_level_metrics=[
        "IncomingBytes",
        "OutgoingBytes",
    ],
    tags={
        "Environment": "test",
    })

