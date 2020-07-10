import pulumi
import pulumi_aws as aws

route = aws.ec2.Route("route",
    route_table_id="rtb-4fbb3ac4",
    destination_cidr_block="10.0.1.0/22",
    vpc_peering_connection_id="pcx-45ff3dc1",
    opts=ResourceOptions(depends_on=["aws_route_table.testing"]))

