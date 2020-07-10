import pulumi
import pulumi_aws as aws

example = aws.s3.AccountPublicAccessBlock("example",
    block_public_acls=True,
    block_public_policy=True)

