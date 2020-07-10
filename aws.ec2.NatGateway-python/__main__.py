import pulumi
import pulumi_aws as aws

gw = aws.ec2.NatGateway("gw",
    allocation_id=aws_eip["nat"]["id"],
    subnet_id=aws_subnet["example"]["id"])

