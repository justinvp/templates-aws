import pulumi
import pulumi_aws as aws
import pulumi_pulumi as pulumi

accepter = pulumi.providers.Aws("accepter")
accepter_caller_identity = aws.get_caller_identity()
# Creator's side of the VIF
creator = aws.directconnect.HostedTransitVirtualInterface("creator",
    address_family="ipv4",
    bgp_asn=65352,
    connection_id="dxcon-zzzzzzzz",
    owner_account_id=accepter_caller_identity.account_id,
    vlan=4094,
    opts=ResourceOptions(depends_on=["aws_dx_gateway.example"]))
# Accepter's side of the VIF.
example = aws.directconnect.Gateway("example", amazon_side_asn=64512,
opts=ResourceOptions(provider="aws.accepter"))
accepter_hosted_transit_virtual_interface_acceptor = aws.directconnect.HostedTransitVirtualInterfaceAcceptor("accepterHostedTransitVirtualInterfaceAcceptor",
    dx_gateway_id=example.id,
    tags={
        "Side": "Accepter",
    },
    virtual_interface_id=creator.id,
    opts=ResourceOptions(provider="aws.accepter"))

