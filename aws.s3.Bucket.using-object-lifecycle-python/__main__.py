import pulumi
import pulumi_aws as aws

bucket = aws.s3.Bucket("bucket",
    acl="private",
    lifecycle_rules=[
        {
            "enabled": True,
            "expiration": {
                "days": 90,
            },
            "id": "log",
            "prefix": "log/",
            "tags": {
                "autoclean": "true",
                "rule": "log",
            },
            "transition": [
                {
                    "days": 30,
                    "storage_class": "STANDARD_IA",
                },
                {
                    "days": 60,
                    "storage_class": "GLACIER",
                },
            ],
        },
        {
            "enabled": True,
            "expiration": {
                "date": "2016-01-12",
            },
            "id": "tmp",
            "prefix": "tmp/",
        },
    ])
versioning_bucket = aws.s3.Bucket("versioningBucket",
    acl="private",
    lifecycle_rules=[{
        "enabled": True,
        "noncurrentVersionExpiration": {
            "days": 90,
        },
        "noncurrentVersionTransition": [
            {
                "days": 30,
                "storage_class": "STANDARD_IA",
            },
            {
                "days": 60,
                "storage_class": "GLACIER",
            },
        ],
        "prefix": "config/",
    }],
    versioning={
        "enabled": True,
    })

