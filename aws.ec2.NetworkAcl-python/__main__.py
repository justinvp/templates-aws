import pulumi
import pulumi_aws as aws

main = aws.ec2.NetworkAcl("main",
    egress=[{
        "action": "allow",
        "cidr_block": "10.3.0.0/18",
        "from_port": 443,
        "protocol": "tcp",
        "ruleNo": 200,
        "to_port": 443,
    }],
    ingress=[{
        "action": "allow",
        "cidr_block": "10.3.0.0/18",
        "from_port": 80,
        "protocol": "tcp",
        "ruleNo": 100,
        "to_port": 80,
    }],
    tags={
        "Name": "main",
    },
    vpc_id=aws_vpc["main"]["id"])

