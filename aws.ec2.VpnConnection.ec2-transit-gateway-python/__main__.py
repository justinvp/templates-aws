import pulumi
import pulumi_aws as aws

example_transit_gateway = aws.ec2transitgateway.TransitGateway("exampleTransitGateway")
example_customer_gateway = aws.ec2.CustomerGateway("exampleCustomerGateway",
    bgp_asn=65000,
    ip_address="172.0.0.1",
    type="ipsec.1")
example_vpn_connection = aws.ec2.VpnConnection("exampleVpnConnection",
    customer_gateway_id=example_customer_gateway.id,
    transit_gateway_id=example_transit_gateway.id,
    type=example_customer_gateway.type)

