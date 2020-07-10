import pulumi
import pulumi_aws as aws

by_version_stage = aws.secretsmanager.get_secret_version(secret_id=data["aws_secretsmanager_secret"]["example"]["id"],
    version_stage="example")

