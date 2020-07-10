import pulumi
import pulumi_aws as aws

log_bucket = aws.s3.Bucket("logBucket", acl="log-delivery-write")
bucket = aws.s3.Bucket("bucket",
    acl="private",
    loggings=[{
        "targetBucket": log_bucket.id,
        "targetPrefix": "log/",
    }])

