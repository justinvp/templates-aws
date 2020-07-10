import pulumi
import pulumi_aws as aws

foo = aws.directconnect.HostedPublicVirtualInterface("foo",
    address_family="ipv4",
    amazon_address="175.45.176.2/30",
    bgp_asn=65352,
    connection_id="dxcon-zzzzzzzz",
    customer_address="175.45.176.1/30",
    route_filter_prefixes=[
        "210.52.109.0/24",
        "175.45.176.0/22",
    ],
    vlan=4094)

