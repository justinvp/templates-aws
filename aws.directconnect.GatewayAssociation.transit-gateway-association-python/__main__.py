import pulumi
import pulumi_aws as aws

example_gateway = aws.directconnect.Gateway("exampleGateway", amazon_side_asn="64512")
example_transit_gateway = aws.ec2transitgateway.TransitGateway("exampleTransitGateway")
example_gateway_association = aws.directconnect.GatewayAssociation("exampleGatewayAssociation",
    allowed_prefixes=[
        "10.255.255.0/30",
        "10.255.255.8/30",
    ],
    associated_gateway_id=example_transit_gateway.id,
    dx_gateway_id=example_gateway.id)

