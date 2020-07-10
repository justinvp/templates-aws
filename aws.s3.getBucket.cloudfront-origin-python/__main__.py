import pulumi
import pulumi_aws as aws

selected = aws.s3.get_bucket(bucket="a-test-bucket")
test = aws.cloudfront.Distribution("test", origins=[{
    "domain_name": selected.bucket_domain_name,
    "originId": "s3-selected-bucket",
}])

