import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_route_table(filters=[
    {
        "name": "default-association-route-table",
        "values": ["true"],
    },
    {
        "name": "transit-gateway-id",
        "values": ["tgw-12345678"],
    },
])

