import pulumi
import pulumi_aws as aws

example = aws.secretsmanager.get_secret_rotation(secret_id=data["aws_secretsmanager_secret"]["example"]["id"])

