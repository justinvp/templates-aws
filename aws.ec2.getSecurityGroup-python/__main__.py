import pulumi
import pulumi_aws as aws

config = pulumi.Config()
security_group_id = config.require_object("securityGroupId")
selected = aws.ec2.get_security_group(id=security_group_id)
subnet = aws.ec2.Subnet("subnet",
    cidr_block="10.0.1.0/24",
    vpc_id=selected.vpc_id)

