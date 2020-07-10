import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.Route("example",
    blackhole=True,
    destination_cidr_block="0.0.0.0/0",
    transit_gateway_route_table_id=aws_ec2_transit_gateway["example"]["association_default_route_table_id"])

