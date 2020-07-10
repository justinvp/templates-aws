import pulumi
import pulumi_aws as aws

s3 = aws.kms.get_alias(name="alias/aws/s3")

