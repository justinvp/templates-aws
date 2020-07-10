import pulumi
import pulumi_aws as aws

example_bucket = aws.s3.Bucket("exampleBucket")
example_access_point = aws.s3.AccessPoint("exampleAccessPoint", bucket=example_bucket.id)

