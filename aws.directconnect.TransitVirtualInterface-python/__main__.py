import pulumi
import pulumi_aws as aws

example_gateway = aws.directconnect.Gateway("exampleGateway", amazon_side_asn=64512)
example_transit_virtual_interface = aws.directconnect.TransitVirtualInterface("exampleTransitVirtualInterface",
    address_family="ipv4",
    bgp_asn=65352,
    connection_id=aws_dx_connection["example"]["id"],
    dx_gateway_id=example_gateway.id,
    vlan=4094)

