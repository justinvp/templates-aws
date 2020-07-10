import pulumi
import pulumi_aws as aws

route_table = aws.ec2.RouteTable("routeTable",
    routes=[
        {
            "cidr_block": "10.0.1.0/24",
            "gateway_id": aws_internet_gateway["main"]["id"],
        },
        {
            "egress_only_gateway_id": aws_egress_only_internet_gateway["foo"]["id"],
            "ipv6_cidr_block": "::/0",
        },
    ],
    tags={
        "Name": "main",
    },
    vpc_id=aws_vpc["default"]["id"])

