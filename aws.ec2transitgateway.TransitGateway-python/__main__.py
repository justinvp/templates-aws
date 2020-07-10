import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.TransitGateway("example", description="example")

