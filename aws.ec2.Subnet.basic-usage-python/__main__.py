import pulumi
import pulumi_aws as aws

main = aws.ec2.Subnet("main",
    cidr_block="10.0.1.0/24",
    tags={
        "Name": "Main",
    },
    vpc_id=aws_vpc["main"]["id"])

