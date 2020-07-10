import pulumi
import pulumi_aws as aws

gw = aws.ec2.InternetGateway("gw",
    tags={
        "Name": "main",
    },
    vpc_id=aws_vpc["main"]["id"])

