import pulumi
import pulumi_aws as aws

config = pulumi.Config()
subnet_id = config.require_object("subnetId")
default = aws.ec2.get_nat_gateway(subnet_id=aws_subnet["public"]["id"])

