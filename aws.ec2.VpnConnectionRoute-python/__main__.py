import pulumi
import pulumi_aws as aws

vpc = aws.ec2.Vpc("vpc", cidr_block="10.0.0.0/16")
vpn_gateway = aws.ec2.VpnGateway("vpnGateway", vpc_id=vpc.id)
customer_gateway = aws.ec2.CustomerGateway("customerGateway",
    bgp_asn=65000,
    ip_address="172.0.0.1",
    type="ipsec.1")
main = aws.ec2.VpnConnection("main",
    customer_gateway_id=customer_gateway.id,
    static_routes_only=True,
    type="ipsec.1",
    vpn_gateway_id=vpn_gateway.id)
office = aws.ec2.VpnConnectionRoute("office",
    destination_cidr_block="192.168.10.0/24",
    vpn_connection_id=main.id)

