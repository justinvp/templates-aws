import pulumi
import pulumi_aws as aws

test = aws.s3.Bucket("test")
inventory = aws.s3.Bucket("inventory")
test_prefix = aws.s3.Inventory("test-prefix",
    bucket=test.id,
    destination={
        "bucket": {
            "bucketArn": inventory.arn,
            "format": "ORC",
            "prefix": "inventory",
        },
    },
    filter={
        "prefix": "documents/",
    },
    included_object_versions="All",
    schedule={
        "frequency": "Daily",
    })

