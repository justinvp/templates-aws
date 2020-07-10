import pulumi
import pulumi_aws as aws

example = aws.dynamodb.Table("example",
    attributes=[{
        "name": "TestTableHashKey",
        "type": "S",
    }],
    billing_mode="PAY_PER_REQUEST",
    hash_key="TestTableHashKey",
    replicas=[
        {
            "regionName": "us-east-2",
        },
        {
            "regionName": "us-west-2",
        },
    ],
    stream_enabled=True,
    stream_view_type="NEW_AND_OLD_IMAGES")

