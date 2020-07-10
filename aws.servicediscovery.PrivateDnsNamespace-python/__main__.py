import pulumi
import pulumi_aws as aws

example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.0.0.0/16")
example_private_dns_namespace = aws.servicediscovery.PrivateDnsNamespace("examplePrivateDnsNamespace",
    description="example",
    vpc=example_vpc.id)

