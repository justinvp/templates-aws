import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_vpn_attachment(transit_gateway_id=aws_ec2_transit_gateway["example"]["id"],
    vpn_connection_id=aws_vpn_connection["example"]["id"])

