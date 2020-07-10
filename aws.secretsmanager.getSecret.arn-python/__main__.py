import pulumi
import pulumi_aws as aws

by_arn = aws.secretsmanager.get_secret(arn="arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456")

