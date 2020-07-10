import pulumi
import pulumi_aws as aws

config = pulumi.Config()
aws_ec2_local_gateway_route_table = config.require_object("awsEc2LocalGatewayRouteTable")
selected = aws.ec2.get_local_gateway_route_table(local_gateway_route_table_id=aws_ec2_local_gateway_route_table)

