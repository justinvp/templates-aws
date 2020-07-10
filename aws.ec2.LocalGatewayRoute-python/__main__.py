import pulumi
import pulumi_aws as aws

example = aws.ec2.LocalGatewayRoute("example",
    destination_cidr_block="172.16.0.0/16",
    local_gateway_route_table_id=data["aws_ec2_local_gateway_route_table"]["example"]["id"],
    local_gateway_virtual_interface_group_id=data["aws_ec2_local_gateway_virtual_interface_group"]["example"]["id"])

