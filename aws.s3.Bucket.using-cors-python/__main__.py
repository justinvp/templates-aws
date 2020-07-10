import pulumi
import pulumi_aws as aws

bucket = aws.s3.Bucket("bucket",
    acl="public-read",
    cors_rules=[{
        "allowedHeaders": ["*"],
        "allowedMethods": [
            "PUT",
            "POST",
        ],
        "allowedOrigins": ["https://s3-website-test.mydomain.com"],
        "exposeHeaders": ["ETag"],
        "maxAgeSeconds": 3000,
    }])

