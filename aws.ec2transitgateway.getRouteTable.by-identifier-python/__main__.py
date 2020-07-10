import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_route_table(id="tgw-rtb-12345678")

