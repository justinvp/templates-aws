import pulumi
import pulumi_aws as aws

example_vpc = aws.ec2.Vpc("exampleVpc",
    assign_generated_ipv6_cidr_block=True,
    cidr_block="10.1.0.0/16")
example_egress_only_internet_gateway = aws.ec2.EgressOnlyInternetGateway("exampleEgressOnlyInternetGateway",
    tags={
        "Name": "main",
    },
    vpc_id=example_vpc.id)

