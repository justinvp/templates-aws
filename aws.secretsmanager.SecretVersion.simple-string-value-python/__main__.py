import pulumi
import pulumi_aws as aws

example = aws.secretsmanager.SecretVersion("example",
    secret_id=aws_secretsmanager_secret["example"]["id"],
    secret_string="example-string-to-protect")

