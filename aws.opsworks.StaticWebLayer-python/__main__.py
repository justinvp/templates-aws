import pulumi
import pulumi_aws as aws

web = aws.opsworks.StaticWebLayer("web", stack_id=aws_opsworks_stack["main"]["id"])

