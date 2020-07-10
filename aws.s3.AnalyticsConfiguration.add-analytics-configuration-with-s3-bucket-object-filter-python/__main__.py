import pulumi
import pulumi_aws as aws

example = aws.s3.Bucket("example")
example_filtered = aws.s3.AnalyticsConfiguration("example-filtered",
    bucket=example.bucket,
    filter={
        "prefix": "documents/",
        "tags": {
            "priority": "high",
            "class": "blue",
        },
    })

