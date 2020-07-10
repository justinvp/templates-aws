import pulumi
import pulumi_aws as aws

config = pulumi.Config()
vpc_id = config.require_object("vpcId")
default = aws.ec2.get_internet_gateway(filters=[{
    "name": "attachment.vpc-id",
    "values": [vpc_id],
}])

