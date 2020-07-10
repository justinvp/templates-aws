import pulumi
import pulumi_aws as aws

foo = aws.ssm.get_parameter(name="foo")

