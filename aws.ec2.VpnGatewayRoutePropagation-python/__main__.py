import pulumi
import pulumi_aws as aws

example = aws.ec2.VpnGatewayRoutePropagation("example",
    route_table_id=aws_route_table["example"]["id"],
    vpn_gateway_id=aws_vpn_gateway["example"]["id"])

