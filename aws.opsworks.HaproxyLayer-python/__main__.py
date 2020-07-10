import pulumi
import pulumi_aws as aws

lb = aws.opsworks.HaproxyLayer("lb",
    stack_id=aws_opsworks_stack["main"]["id"],
    stats_password="foobarbaz")

