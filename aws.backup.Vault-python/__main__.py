import pulumi
import pulumi_aws as aws

example = aws.backup.Vault("example", kms_key_arn=aws_kms_key["example"]["arn"])

