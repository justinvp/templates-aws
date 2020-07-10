import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_direct_connect_gateway_attachment(dx_gateway_id=aws_dx_gateway["example"]["id"],
    transit_gateway_id=aws_ec2_transit_gateway["example"]["id"])

