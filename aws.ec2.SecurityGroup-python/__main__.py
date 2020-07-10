import pulumi
import pulumi_aws as aws

allow_tls = aws.ec2.SecurityGroup("allowTls",
    description="Allow TLS inbound traffic",
    vpc_id=aws_vpc["main"]["id"],
    ingress=[{
        "description": "TLS from VPC",
        "from_port": 443,
        "to_port": 443,
        "protocol": "tcp",
        "cidr_blocks": [aws_vpc["main"]["cidr_block"]],
    }],
    egress=[{
        "from_port": 0,
        "to_port": 0,
        "protocol": "-1",
        "cidr_blocks": ["0.0.0.0/0"],
    }],
    tags={
        "Name": "allow_tls",
    })

