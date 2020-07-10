import pulumi
import pulumi_aws as aws

vpn_gw = aws.ec2.VpnGateway("vpnGw",
    tags={
        "Name": "main",
    },
    vpc_id=aws_vpc["main"]["id"])

