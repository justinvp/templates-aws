import pulumi
import pulumi_aws as aws

rotation_example = aws.secretsmanager.Secret("rotation-example",
    rotation_lambda_arn=aws_lambda_function["example"]["arn"],
    rotation_rules={
        "automaticallyAfterDays": 7,
    })

