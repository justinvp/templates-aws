import pulumi
import pulumi_aws as aws

foo = aws.ec2.get_customer_gateway(filters=[{
    "name": "tag:Name",
    "values": ["foo-prod"],
}])
main = aws.ec2.VpnGateway("main",
    amazon_side_asn=7224,
    vpc_id=aws_vpc["main"]["id"])
transit = aws.ec2.VpnConnection("transit",
    customer_gateway_id=foo.id,
    static_routes_only=False,
    type=foo.type,
    vpn_gateway_id=main.id)

