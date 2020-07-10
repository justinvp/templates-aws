import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.RouteTable("example", transit_gateway_id=aws_ec2_transit_gateway["example"]["id"])

