import pulumi
import pulumi_aws as aws

service = aws.ecr.get_repository(name="ecr-repository")

