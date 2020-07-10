import pulumi
import pulumi_aws as aws

example_bucket = aws.s3.Bucket("exampleBucket")
example_bucket_public_access_block = aws.s3.BucketPublicAccessBlock("exampleBucketPublicAccessBlock",
    block_public_acls=True,
    block_public_policy=True,
    bucket=example_bucket.id)

