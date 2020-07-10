import pulumi
import pulumi_aws as aws

example = aws.s3.Bucket("example")
example_entire_bucket = aws.s3.BucketMetric("example-entire-bucket", bucket=example.bucket)

