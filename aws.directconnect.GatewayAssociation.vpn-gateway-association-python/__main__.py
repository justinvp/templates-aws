import pulumi
import pulumi_aws as aws

example_gateway = aws.directconnect.Gateway("exampleGateway", amazon_side_asn="64512")
example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.255.255.0/28")
example_vpn_gateway = aws.ec2.VpnGateway("exampleVpnGateway", vpc_id=example_vpc.id)
example_gateway_association = aws.directconnect.GatewayAssociation("exampleGatewayAssociation",
    associated_gateway_id=example_vpn_gateway.id,
    dx_gateway_id=example_gateway.id)

