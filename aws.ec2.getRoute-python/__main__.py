import pulumi
import pulumi_aws as aws

config = pulumi.Config()
subnet_id = config.require_object("subnetId")
selected = aws.ec2.get_route_table(subnet_id=subnet_id)
route = aws.ec2.get_route(destination_cidr_block="10.0.1.0/24",
    route_table_id=aws_route_table["selected"]["id"])
interface = aws.ec2.get_network_interface(network_interface_id=route.network_interface_id)

