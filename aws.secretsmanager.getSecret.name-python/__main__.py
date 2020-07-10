import pulumi
import pulumi_aws as aws

by_name = aws.secretsmanager.get_secret(name="example")

