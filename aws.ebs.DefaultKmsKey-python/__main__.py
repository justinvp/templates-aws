import pulumi
import pulumi_aws as aws

example = aws.ebs.DefaultKmsKey("example", key_arn=aws_kms_key["example"]["arn"])

