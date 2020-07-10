import pulumi
import pulumi_aws as aws

test_stream = aws.kinesis.Stream("testStream", shard_count=1)
test_application = aws.kinesis.AnalyticsApplication("testApplication", inputs={
    "kinesisStream": {
        "resource_arn": test_stream.arn,
        "role_arn": aws_iam_role["test"]["arn"],
    },
    "name_prefix": "test_prefix",
    "parallelism": {
        "count": 1,
    },
    "schema": {
        "recordColumns": [{
            "mapping": "$.test",
            "name": "test",
            "sqlType": "VARCHAR(8)",
        }],
        "recordEncoding": "UTF-8",
        "recordFormat": {
            "mappingParameters": {
                "json": {
                    "recordRowPath": "$",
                },
            },
        },
    },
})

