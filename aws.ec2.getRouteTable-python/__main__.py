import pulumi
import pulumi_aws as aws

config = pulumi.Config()
subnet_id = config.require_object("subnetId")
selected = aws.ec2.get_route_table(subnet_id=subnet_id)
route = aws.ec2.Route("route",
    destination_cidr_block="10.0.1.0/22",
    route_table_id=selected.id,
    vpc_peering_connection_id="pcx-45ff3dc1")

