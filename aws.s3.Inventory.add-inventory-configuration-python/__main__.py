import pulumi
import pulumi_aws as aws

test_bucket = aws.s3.Bucket("testBucket")
inventory = aws.s3.Bucket("inventory")
test_inventory = aws.s3.Inventory("testInventory",
    bucket=test_bucket.id,
    destination={
        "bucket": {
            "bucketArn": inventory.arn,
            "format": "ORC",
        },
    },
    included_object_versions="All",
    schedule={
        "frequency": "Daily",
    })

