import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_transit_gateway(id="tgw-12345678")

