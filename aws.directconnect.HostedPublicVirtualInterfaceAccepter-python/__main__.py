import pulumi
import pulumi_aws as aws
import pulumi_pulumi as pulumi

accepter = pulumi.providers.Aws("accepter")
accepter_caller_identity = aws.get_caller_identity()
# Creator's side of the VIF
creator = aws.directconnect.HostedPublicVirtualInterface("creator",
    address_family="ipv4",
    amazon_address="175.45.176.2/30",
    bgp_asn=65352,
    connection_id="dxcon-zzzzzzzz",
    customer_address="175.45.176.1/30",
    owner_account_id=accepter_caller_identity.account_id,
    route_filter_prefixes=[
        "210.52.109.0/24",
        "175.45.176.0/22",
    ],
    vlan=4094)
# Accepter's side of the VIF.
accepter_hosted_public_virtual_interface_accepter = aws.directconnect.HostedPublicVirtualInterfaceAccepter("accepterHostedPublicVirtualInterfaceAccepter",
    tags={
        "Side": "Accepter",
    },
    virtual_interface_id=creator.id,
    opts=ResourceOptions(provider="aws.accepter"))

