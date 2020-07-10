import pulumi
import pulumi_aws as aws

app = aws.opsworks.PhpAppLayer("app", stack_id=aws_opsworks_stack["main"]["id"])

