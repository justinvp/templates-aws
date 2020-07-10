import pulumi
import pulumi_aws as aws
import pulumi_pulumi as pulumi

accepter = pulumi.providers.Aws("accepter")
accepter_caller_identity = aws.get_caller_identity()
# Creator's side of the VIF
creator = aws.directconnect.HostedPrivateVirtualInterface("creator",
    address_family="ipv4",
    bgp_asn=65352,
    connection_id="dxcon-zzzzzzzz",
    owner_account_id=accepter_caller_identity.account_id,
    vlan=4094,
    opts=ResourceOptions(depends_on=["aws_vpn_gateway.vpn_gw"]))
# Accepter's side of the VIF.
vpn_gw = aws.ec2.VpnGateway("vpnGw", opts=ResourceOptions(provider="aws.accepter"))
accepter_hosted_private_virtual_interface_accepter = aws.directconnect.HostedPrivateVirtualInterfaceAccepter("accepterHostedPrivateVirtualInterfaceAccepter",
    tags={
        "Side": "Accepter",
    },
    virtual_interface_id=creator.id,
    vpn_gateway_id=vpn_gw.id,
    opts=ResourceOptions(provider="aws.accepter"))

