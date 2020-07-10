import pulumi
import pulumi_aws as aws

selected = aws.ec2.get_vpn_gateway(filters=[{
    "name": "tag:Name",
    "values": ["vpn-gw"],
}])
pulumi.export("vpnGatewayId", selected.id)

