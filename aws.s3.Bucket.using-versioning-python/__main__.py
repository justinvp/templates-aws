import pulumi
import pulumi_aws as aws

bucket = aws.s3.Bucket("bucket",
    acl="private",
    versioning={
        "enabled": True,
    })

