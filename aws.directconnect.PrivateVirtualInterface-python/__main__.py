import pulumi
import pulumi_aws as aws

foo = aws.directconnect.PrivateVirtualInterface("foo",
    address_family="ipv4",
    bgp_asn=65352,
    connection_id="dxcon-zzzzzzzz",
    vlan=4094)

