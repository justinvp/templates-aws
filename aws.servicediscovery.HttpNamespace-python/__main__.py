import pulumi
import pulumi_aws as aws

example = aws.servicediscovery.HttpNamespace("example", description="example")

