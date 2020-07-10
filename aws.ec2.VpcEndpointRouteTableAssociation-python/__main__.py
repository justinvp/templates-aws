import pulumi
import pulumi_aws as aws

example = aws.ec2.VpcEndpointRouteTableAssociation("example",
    route_table_id=aws_route_table["example"]["id"],
    vpc_endpoint_id=aws_vpc_endpoint["example"]["id"])

