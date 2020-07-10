import pulumi
import pulumi_aws as aws

default_route_table = aws.ec2.DefaultRouteTable("defaultRouteTable",
    default_route_table_id=aws_vpc["foo"]["default_route_table_id"],
    routes=[{}],
    tags={
        "Name": "default table",
    })

