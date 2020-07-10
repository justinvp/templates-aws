import pulumi
import pulumi_aws as aws

foo = aws.autoscaling.get_group(name="foo")

