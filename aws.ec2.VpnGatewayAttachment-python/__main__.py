import pulumi
import pulumi_aws as aws

network = aws.ec2.Vpc("network", cidr_block="10.0.0.0/16")
vpn = aws.ec2.VpnGateway("vpn", tags={
    "Name": "example-vpn-gateway",
})
vpn_attachment = aws.ec2.VpnGatewayAttachment("vpnAttachment",
    vpc_id=network.id,
    vpn_gateway_id=vpn.id)

