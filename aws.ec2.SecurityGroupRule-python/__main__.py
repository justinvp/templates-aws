import pulumi
import pulumi_aws as aws

example = aws.ec2.SecurityGroupRule("example",
    type="ingress",
    from_port=0,
    to_port=65535,
    protocol="tcp",
    cidr_blocks=aws_vpc["example"]["cidr_block"],
    security_group_id="sg-123456")

