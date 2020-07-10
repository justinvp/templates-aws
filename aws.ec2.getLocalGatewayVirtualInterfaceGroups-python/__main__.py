import pulumi
import pulumi_aws as aws

all = aws.ec2.get_local_gateway_virtual_interface_groups()

