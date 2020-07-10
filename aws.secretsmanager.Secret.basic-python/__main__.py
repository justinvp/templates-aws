import pulumi
import pulumi_aws as aws

example = aws.secretsmanager.Secret("example")

