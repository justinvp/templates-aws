import pulumi
import pulumi_aws as aws

app = aws.opsworks.JavaAppLayer("app", stack_id=aws_opsworks_stack["main"]["id"])

