import pulumi
import pulumi_aws as aws

example = aws.directconnect.HostedTransitVirtualInterface("example",
    address_family="ipv4",
    bgp_asn=65352,
    connection_id=aws_dx_connection["example"]["id"],
    vlan=4094)

