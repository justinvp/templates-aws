import pulumi
import pulumi_aws as aws

current = aws.ebs.get_encryption_by_default()

