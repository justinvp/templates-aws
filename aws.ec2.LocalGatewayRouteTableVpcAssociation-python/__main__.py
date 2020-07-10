import pulumi
import pulumi_aws as aws

example_local_gateway_route_table = aws.ec2.get_local_gateway_route_table(outpost_arn="arn:aws:outposts:us-west-2:123456789012:outpost/op-1234567890abcdef")
example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.0.0.0/16")
example_local_gateway_route_table_vpc_association = aws.ec2.LocalGatewayRouteTableVpcAssociation("exampleLocalGatewayRouteTableVpcAssociation",
    local_gateway_route_table_id=example_local_gateway_route_table.id,
    vpc_id=example_vpc.id)

