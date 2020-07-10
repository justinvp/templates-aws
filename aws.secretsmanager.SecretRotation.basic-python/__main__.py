import pulumi
import pulumi_aws as aws

example = aws.secretsmanager.SecretRotation("example",
    rotation_lambda_arn=aws_lambda_function["example"]["arn"],
    rotation_rules={
        "automaticallyAfterDays": 30,
    },
    secret_id=aws_secretsmanager_secret["example"]["id"])

