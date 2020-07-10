import pulumi
import pulumi_aws as aws

peer = aws.directconnect.BgpPeer("peer",
    address_family="ipv6",
    bgp_asn=65351,
    virtual_interface_id=aws_dx_private_virtual_interface["foo"]["id"])

