import pulumi
import pulumi_aws as aws

config = pulumi.Config()
local_gateway_id = config.require_object("localGatewayId")
selected = aws.ec2.get_local_gateway(id=local_gateway_id)

