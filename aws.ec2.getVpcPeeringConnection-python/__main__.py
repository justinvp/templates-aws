import pulumi
import pulumi_aws as aws

pc = aws.ec2.get_vpc_peering_connection(peer_cidr_block="10.0.1.0/22",
    vpc_id=aws_vpc["foo"]["id"])
# Create a route table
rt = aws.ec2.RouteTable("rt", vpc_id=aws_vpc["foo"]["id"])
# Create a route
route = aws.ec2.Route("route",
    destination_cidr_block=pc.peer_cidr_block,
    route_table_id=rt.id,
    vpc_peering_connection_id=pc.id)

