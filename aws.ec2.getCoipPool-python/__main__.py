import pulumi
import pulumi_aws as aws

config = pulumi.Config()
coip_pool_id = config.require_object("coipPoolId")
selected = aws.ec2.get_coip_pool(id=coip_pool_id)

