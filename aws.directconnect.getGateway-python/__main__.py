import pulumi
import pulumi_aws as aws

example = aws.directconnect.get_gateway(name="example")

