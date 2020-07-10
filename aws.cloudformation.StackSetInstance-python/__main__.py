import pulumi
import pulumi_aws as aws

example = aws.cloudformation.StackSetInstance("example",
    account_id="123456789012",
    region="us-east-1",
    stack_set_name=aws_cloudformation_stack_set["example"]["name"])

