import pulumi
import pulumi_aws as aws

example = aws.servicediscovery.PublicDnsNamespace("example", description="example")

