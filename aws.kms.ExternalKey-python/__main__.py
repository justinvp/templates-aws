import pulumi
import pulumi_aws as aws

example = aws.kms.ExternalKey("example", description="KMS EXTERNAL for AMI encryption")

